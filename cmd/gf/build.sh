#!/bin/bash

Dir="/media/hanwei/dev/go/repo/gf/cmd/gf/bin/$1"
echo "${Dir}"
if [ -d "${Dir}" ]; then
    echo "目录已存在,清空目录内容"
    rm -rf "${Dir}"
    mkdir -p "${Dir}"
else
      echo "创建目录"
      mkdir -p "${Dir}"
fi


gf build -version="$1" -e "-trimpath -ldflags '%s -w -s' -buildmode=pie"  main.go


function mvgGfToDir() {
    for file in `ls $1`; do
       if [ -d "$1/$file" ]; then
            mvgGfToDir "$1/$file"
          else
            if [[ "$file" == *"gf"* ]]; then
                last_level=$(basename "$1")
                echo "${last_level}"
                if [[  -e "$1/$file" ]]; then
                    upx -9 "$1/$file"
                fi
                # shellcheck disable=SC1068
                filename="$Dir/gf_${last_level}"
                # shellcheck disable=SC2081
                if [[ "${filename}" == *"windows_"* ]]; then
                    # shellcheck disable=SC1068
                    filename="${filename}.exe"
                fi
                mv "$1/$file" "$filename"
                # 删除生成的目录
                rm -rf "$1"
            fi
       fi
    done
}

echo "移动目录"
mvgGfToDir "$Dir"