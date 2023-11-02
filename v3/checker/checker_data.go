package checker

import (
	"sort"
	"strconv"
	"strings"

	"github.com/davyxu/tabtoy/v3/model"
	"github.com/davyxu/tabtoy/v3/report"
)

// 检查数据与定义类型是否匹配
func checkDataType(globals *model.Globals) {

	var currHeader *model.HeaderField
	var crrCell *model.Cell

	for _, tab := range globals.Datas.AllTables() {

		tab.CommonCols = make([]*model.CommonCol, len(tab.Headers))
		// 遍历输入数据的每一列
		for _, header := range tab.Headers {

			// 输入的列头，为空表示改列被注释
			if header.TypeInfo == nil {
				continue
			}

			colCommonCount := make(map[string]*model.CommonCol) // 用于存放最常出现列的数据
			for row := 1; row < len(tab.Rows); row++ {

				inputCell := tab.GetCell(row, header.Cell.Col)

				// 这行被注释，无效行
				if inputCell == nil {
					continue
				}

				crrCell = inputCell
				currHeader = header

				if header.TypeInfo.IsArray() {
					for _, value := range inputCell.ValueList {

						err := checkSingleValue(header, value)
						if err != nil {
							report.ReportError("DataMissMatchTypeDefine", currHeader.TypeInfo.FieldType, crrCell.String())
						}
					}
				} else if inputCell.Value != "" {
					err := checkSingleValue(header, inputCell.Value)
					if err != nil {
						report.ReportError("DataMissMatchTypeDefine", currHeader.TypeInfo.FieldType, crrCell.String())
					}
				}

				var key = inputCell.Value + "..." + strings.Join(inputCell.ValueList, "|")
				if _, ok := colCommonCount[key]; !ok {
					colCommonCount[key] = &model.CommonCol{
						FieldName: header.TypeInfo.FieldName,
						Value:     inputCell.Value,
						ValueList: inputCell.ValueList,
						UnKey:     key,
					}
				}
				colCommonCount[key].Count += 1
			}
			sortData := make([]*model.CommonCol, 0, len(colCommonCount))
			for _, v := range colCommonCount {
				sortData = append(sortData, v)
			}
			sort.Slice(sortData, func(i, j int) bool {
				if sortData[i].Count == sortData[j].Count {
					// 同样次数的处理情况
					return sortData[i].UnKey < sortData[j].UnKey
				}
				return sortData[i].Count > sortData[j].Count
			})
			tab.CommonCols[header.Cell.Col] = sortData[0]
		}
	}
}

func checkSingleValue(header *model.HeaderField, value string) error {
	switch model.LanguagePrimitive(header.TypeInfo.FieldType, "go") {
	case "int16":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
	case "int32":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
	case "int64":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
	case "uint16":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
	case "uint32":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
	case "uint64":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
	case "float32":
		_, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
	case "float64":
		if value == "" {
			return nil
		}
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
	case "bool":
		_, err := model.ParseBool(value)
		if err != nil {
			return err
		}
	}

	return nil
}
