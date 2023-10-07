package bindata

import (
	"strings"

	"github.com/davyxu/tabtoy/v3/model"
	"github.com/davyxu/tabtoy/v3/report"
)

func MakeTag(globals *model.Globals, tf *model.TypeDefine, fieldIndex int) uint32 {
	convertedType := model.LanguagePrimitive(tf.FieldType, "go")
	isHadderStruct := model.HadderStructCache[tf.FieldType]
	var t int
	switch {
	case convertedType == "int16":
		t = 1
	case convertedType == "int32":
		t = 2
	case convertedType == "int64":
		t = 3
	case convertedType == "uint16":
		t = 4
	case convertedType == "uint32":
		t = 5
	case convertedType == "uint64":
		t = 6
	case convertedType == "float32":
		t = 7
	case convertedType == "string":
		t = 8
	case convertedType == "bool":
		t = 9
	case globals.Types.IsEnumKind(tf.FieldType):
		t = 10
	case convertedType == "float64":
		t = 12
		// 注意, t = 11是结构体
	case isHadderStruct != nil:
		t = 11
	default:
		panic("unknown type:" + tf.FieldType)
	}

	if tf.IsArray() {
		t += 100
	}

	return uint32(t<<16 | fieldIndex)
}

func MakeTagStructArray() uint32 {

	var t int
	t = 11

	// 结构体默认是数组
	t += 100

	return uint32(t << 16)
}

func writePair(globals *model.Globals, structWriter *BinaryWriter, fieldType *model.TypeDefine, goType, value string, fieldIndex int, lenWrite bool) error {
	tag := MakeTag(globals, fieldType, fieldIndex)
	if err := structWriter.WriteUInt32(tag); err != nil {
		return err
	}
	return structWritePair(globals, structWriter, fieldType, goType, value, fieldIndex, lenWrite)
}

func structWritePair(globals *model.Globals, structWriter *BinaryWriter, fieldType *model.TypeDefine, goType, value string, fieldIndex int, lenWrite bool) error {
	var fields = model.HadderStructCache[fieldType.FieldType]
	if fields != nil {
		// 结构体二进制边界
		newStructWriter := NewBinaryWriter()

		if fieldType.ArraySplitter != "" {
			dataList := strings.Split(value, fieldType.ArraySplitter)
			structWriter.WriteUInt32(uint32(len(dataList)))

			newFieldType := *fieldType
			newFieldType.ArraySplitter = ""
			for structIndex, element := range dataList {
				err := structWritePair(globals, newStructWriter, &newFieldType, newFieldType.FieldType, element, structIndex, false)
				if err != nil {
					return err
				}
			}
		} else {
			for structIndex, tmp := range strings.Split(value, " ") {
				data := strings.Split(tmp, ":")
				if len(data) != 2 {
					report.ReportError("UnknownTypeKind", fieldType.ObjectType, fieldType.FieldName)
				}
				typeDefine := fields.TypeInfo[data[0]]
				err := writePair(globals, newStructWriter, typeDefine, typeDefine.FieldType, data[1], structIndex, false)
				if err != nil {
					return err
				}
			}
		}

		structData := newStructWriter.Bytes()
		if lenWrite {
			// 结构体二进制边界
			structWriter.WriteUInt32(uint32(len(structData)))
		}
		structWriter.Write(structData)
		return nil
	}
	return writeValue(globals, structWriter, fieldType, goType, value)
}
