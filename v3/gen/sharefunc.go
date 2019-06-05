package gen

import (
	"github.com/ahmetb/go-linq"
	"github.com/davyxu/tabtoy/util"
	"github.com/davyxu/tabtoy/v3/model"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

type TableIndices struct {
	Table     *model.DataTable
	FieldInfo *model.TypeDefine
}

func KeyValueTypeNames(globals *model.Globals) (ret []string) {
	linq.From(globals.IndexList).WhereT(func(pragma *model.IndexDefine) bool {
		return pragma.Kind == model.TableKind_KeyValue
	}).SelectT(func(pragma *model.IndexDefine) string {

		return pragma.TableType
	}).Distinct().ToSlice(&ret)

	return
}

func WrapSingleValue(globals *model.Globals, valueType *model.TypeDefine, value string) string {
	switch {
	case valueType.FieldType == "string": // 字符串
		return util.StringEscape(value)
	case valueType.FieldType == "float":
		return value
	case globals.Types.IsEnumKind(valueType.FieldType): // 枚举
		return globals.Types.ResolveEnumValue(valueType.FieldType, value)
	case valueType.FieldType == "bool":

		switch value {
		case "是", "yes", "YES", "1", "true", "TRUE", "True":
			return "true"
		case "否", "no", "NO", "0", "false", "FALSE", "False":
			return "false"
		}

		return "false"
	}

	if value == "" {
		return model.FetchDefaultValue(valueType.FieldType)
	}

	return value
}

func init() {
	UsefulFunc["HasKeyValueTypes"] = func(globals *model.Globals) bool {
		return len(KeyValueTypeNames(globals)) > 0
	}

	UsefulFunc["GetKeyValueTypeNames"] = KeyValueTypeNames

	UsefulFunc["GetIndices"] = func(globals *model.Globals) (ret []TableIndices) {

		for _, tab := range globals.Datas.AllTables() {

			// 遍历输入数据的每一列
			for _, header := range tab.Headers {

				// 输入的列头
				if header.TypeInfo == nil {
					continue
				}

				if header.TypeInfo.MakeIndex {

					ret = append(ret, TableIndices{
						Table:     tab,
						FieldInfo: header.TypeInfo,
					})
				}
			}
		}

		return

	}
}
