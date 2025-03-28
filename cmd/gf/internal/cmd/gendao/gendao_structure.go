// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gendao

import (
	"bytes"
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/olekukonko/tablewriter"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

type generateStructDefinitionInput struct {
	CGenDaoInternalInput
	TableName  string                     // Table name.
	StructName string                     // Struct name.
	FieldMap   map[string]*gdb.TableField // Table field map.
	IsDo       bool                       // Is generating DTO struct.
}

func generateStructDefinition(ctx context.Context, in generateStructDefinitionInput) (string, []string) {
	var appendImports []string
	buffer := bytes.NewBuffer(nil)
	array := make([][]string, len(in.FieldMap))
	names := sortFieldKeyForDao(in.FieldMap)
	for index, name := range names {
		var imports string
		field := in.FieldMap[name]
		array[index], imports = generateStructFieldDefinition(ctx, field, in)
		if imports != "" {
			appendImports = append(appendImports, imports)
		}
	}
	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(array)
	tw.Render()
	stContent := buffer.String()
	// Let's do this hack of table writer for indent!
	stContent = gstr.Replace(stContent, "  #", "")
	stContent = gstr.Replace(stContent, "` ", "`")
	stContent = gstr.Replace(stContent, "``", "")
	buffer.Reset()
	buffer.WriteString(fmt.Sprintf("type %s struct {\n", in.StructName))
	if in.IsDo {
		buffer.WriteString(fmt.Sprintf("g.Meta `orm:\"table:%s, do:true\"`\n", in.TableName))
	}
	buffer.WriteString(stContent)
	buffer.WriteString("}")
	return buffer.String(), appendImports
}

// generateStructFieldDefinition generates and returns the attribute definition for specified field.
func generateStructFieldDefinition(
	ctx context.Context, field *gdb.TableField, in generateStructDefinitionInput,
) (attrLines []string, appendImport string) {
	var (
		err              error
		localTypeName    gdb.LocalType
		localTypeNameStr string
		jsonTag          = gstr.CaseConvert(field.Name, gstr.CaseTypeMatch(in.JsonCase))
	)

	if in.TypeMapping != nil && len(in.TypeMapping) > 0 {
		var (
			tryTypeName string
		)
		tryTypeMatch, _ := gregex.MatchString(`(.+?)\((.+)\)`, field.Type)
		if len(tryTypeMatch) == 3 {
			tryTypeName = gstr.Trim(tryTypeMatch[1])
		} else {
			tryTypeName = gstr.Split(field.Type, " ")[0]
		}
		if tryTypeName != "" {
			if typeMapping, ok := in.TypeMapping[strings.ToLower(tryTypeName)]; ok {
				localTypeNameStr = typeMapping.Type
				appendImport = typeMapping.Import
			}
		}

		if typeMapping, ok := in.TypeMapping[strings.ToLower(field.Type)]; ok {
			localTypeNameStr = typeMapping.Type
			appendImport = typeMapping.Import
		}
	}

	if localTypeNameStr == "" {
		localTypeName, err = in.DB.CheckLocalTypeForField(ctx, field.Type, nil)
		if err != nil {
			panic(err)
		}
		localTypeNameStr = string(localTypeName)
		switch localTypeName {
		case gdb.LocalTypeDate, gdb.LocalTypeTime, gdb.LocalTypeDatetime:
			if in.StdTime {
				localTypeNameStr = "time.Time"
			} else {
				localTypeNameStr = "*gtime.Time"
			}

		case gdb.LocalTypeInt64Bytes:
			localTypeNameStr = "int64"

		case gdb.LocalTypeUint64Bytes:
			localTypeNameStr = "uint64"

		// Special type handle.
		case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
			if in.GJsonSupport {
				localTypeNameStr = "*gjson.Json"
			} else {
				localTypeNameStr = "string"
			}
		}
	}

	var (
		//tagKey         = "`"
		descriptionTag = gstr.Replace(formatComment(field.Comment), `"`, `\"`)
	)
	removeFieldPrefixArray := gstr.SplitAndTrim(in.RemoveFieldPrefix, ",")
	newFiledName := field.Name
	for _, v := range removeFieldPrefixArray {
		newFiledName = gstr.TrimLeftStr(newFiledName, v, 1)
	}

	if in.FieldMapping != nil && len(in.FieldMapping) > 0 {
		if typeMapping, ok := in.FieldMapping[fmt.Sprintf("%s.%s", in.TableName, newFiledName)]; ok {
			localTypeNameStr = typeMapping.Type
			appendImport = typeMapping.Import
		}
	}

	attrLines = []string{
		"    #" + formatFieldName(newFiledName, FieldNameCaseCamel),
		" #" + localTypeNameStr,
	}

	attrs := make([]string, 0, 10)

	if !in.NoJsonTag {
		if slices.Index(in.JsonIgnoreField, field.Name) > -1 {
			//attrs = append(attrs, " #json:\"-\"")
			attrs = append(attrs, fmt.Sprintf(` #json:"-" gconv:"%s"`, jsonTag))
		} else {

			var omitempty bool
			if len(in.Omitempty.All) > 0 {
				for _, s := range in.Omitempty.All {
					if s == jsonTag {
						omitempty = true
						break
					}
				}
			}
			if !omitempty && len(in.Omitempty.Table) > 0 {
				if v001, ok001 := in.Omitempty.Table[in.TableName]; ok001 {
					for _, s := range v001 {
						if s == jsonTag {
							omitempty = true
							break
						}
					}
				}
			}
			if omitempty {
				attrs = append(attrs, fmt.Sprintf(` #json:"%s" gconv:"%s"`, jsonTag, jsonTag))
			} else {
				attrs = append(attrs, fmt.Sprintf(` #json:"%s,omitempty" gconv:"%s"`, jsonTag, jsonTag))
			}
		}
	}

	// 主键
	if in.HashTag.Pk && field.Key == "PRI" {
		attrs = append(attrs, " #hashids:\"true\"")
	}

	if v, ok := in.HashTag.Table[in.TableName]; ok && slices.Index(v, field.Name) > -1 && (field.Key != "PRI" || !in.HashTag.Pk) {
		attrs = append(attrs, " #hashids:\"true\"")
	}

	// orm tag
	if !in.IsDo {
		// entity
		//attrLines = append(attrLines, fmt.Sprintf(` #orm:"%s"`, field.Name))
		attrs = append(attrs, fmt.Sprintf(` #orm:"%s"`, field.Name))
	}
	//attrLines = append(attrLines, fmt.Sprintf(` #description:"%s"%s`, descriptionTag, tagKey))
	//attrLines = append(attrLines, fmt.Sprintf(` #// %s`, formatComment(field.Comment)))

	if in.DescriptionTag { // 生成描述标签
		attrs = append(attrs, fmt.Sprintf(` #description:"%s"`, descriptionTag))
	}
	for k, v := range attrs {
		if in.NoJsonTag {
			v, _ = gregex.ReplaceString(`json:".+"`, ``, v)
		}
		if !in.DescriptionTag {
			v, _ = gregex.ReplaceString(`description:".*"`, ``, v)
		}
		if in.NoModelComment {
			v, _ = gregex.ReplaceString(`//.+`, ``, v)
		}
		attrs[k] = v
	}

	if len(attrs) > 0 {
		attrLines = append(attrLines, "`")
		attrLines = append(attrLines, attrs...)
		attrLines = append(attrLines, "`")
	}

	if !in.NoModelComment { // 生成注释
		attrLines = append(attrLines, fmt.Sprintf(` #// %s`, formatComment(field.Comment)))
	}

	return attrLines, appendImport
}

type FieldNameCase string

const (
	FieldNameCaseCamel      FieldNameCase = "CaseCamel"
	FieldNameCaseCamelLower FieldNameCase = "CaseCamelLower"
)

// formatFieldName formats and returns a new field name that is used for golang codes generating.
func formatFieldName(fieldName string, nameCase FieldNameCase) string {
	// For normal databases like mysql, pgsql, sqlite,
	// field/table names of that are in normal case.
	var newFieldName = fieldName
	if isAllUpper(fieldName) {
		// For special databases like dm, oracle,
		// field/table names of that are in upper case.
		newFieldName = strings.ToLower(fieldName)
	}
	switch nameCase {
	case FieldNameCaseCamel:
		return gstr.CaseCamel(newFieldName)
	case FieldNameCaseCamelLower:
		return gstr.CaseCamelLower(newFieldName)
	default:
		return ""
	}
}

// isAllUpper checks and returns whether given `fieldName` all letters are upper case.
func isAllUpper(fieldName string) bool {
	for _, b := range fieldName {
		if b >= 'a' && b <= 'z' {
			return false
		}
	}
	return true
}

// formatComment formats the comment string to fit the golang code without any lines.
func formatComment(comment string) string {
	comment = gstr.ReplaceByArray(comment, g.SliceStr{
		"\n", " ",
		"\r", " ",
	})
	comment = gstr.Replace(comment, `\n`, " ")
	comment = gstr.Trim(comment)
	return comment
}
