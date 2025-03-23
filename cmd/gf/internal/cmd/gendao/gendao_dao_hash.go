package gendao



type (
	HashIdTag struct {
		Pk    bool                `json:"pk" name:"pk" orphan:"true"  brief:"{HashIdTagPk}"`
		Table map[string][]string `json:"table" name:"table" orphan:"true"  brief:"{HashIdTagTable}"`
	}

	// 忽略生成Omitempty标签的字段
	Omitempty struct {
		All   []string            `json:"all" name:"all" orphan:"true"  brief:"{OmitemptyAll}"`
		Table map[string][]string `json:"table" name:"table" orphan:"true"  brief:"{OmitemptyTable}"`
	}

)