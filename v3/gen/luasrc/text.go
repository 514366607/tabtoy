package luasrc

const templateText_luasrc = `-- Generated by github.com/davyxu/tabtoy
-- Version: {{.Version}}

return {
	init = function( g )
		{{range $di, $tab := .Datas.AllTables}}
		g.{{$tab.HeaderType}} = { {{range $unusedrow,$row := $tab.DataRowIndex}}{{$headers := $.Types.AllFieldByName $tab.OriginalHeaderType}}
			{ {{range $col, $header := $headers}}{{if IsWrapFieldName $ $tab $headers $row $col}}{{$header.FieldName}} = {{WrapTabValue $ $tab $headers $row $col}}, {{end}}{{end}}}, {{end}}
		}

		for _, value in pairs(g.{{$tab.HeaderType}}) do
			setmetatable(value, {
				__index = { {{range $fieldName, $value := DefaultField $ $tab}} {{$fieldName}} = {{$value}},{{end}} },
			})
		end
		{{end}}
		{{range $ii, $idx := GetIndices $}}
		-- {{$idx.Table.HeaderType}}
		g.{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}} = {}
		for _, rec in pairs(g.{{$idx.Table.HeaderType}}) do
			g.{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}}[rec.{{$idx.FieldInfo.FieldName}}] = rec
		end
		{{end}}
		{{range $sn, $objName := $.Types.EnumNames}}
		g.{{$objName}} = { {{range $fi,$field := $.Types.AllFieldByName $objName}}
			{{$field.FieldName}} = {{$field.Value}}, -- {{if not $field.Note}}{{$field.Name}}{{else}}{{ $field.Note}}{{end}} {{end}} {{range $fi,$field := $.Types.AllFieldByName $objName}}
			[{{$field.Value}}] = "{{$field.FieldName}}",{{end}}
		}{{end}}
		return g
	end
}
`

const templateText_luadir = `-- Generated by github.com/davyxu/tabtoy
-- Version: {{$.G.Version}}

return {
	init = function( g )
		g.{{$.Tab.HeaderType}} = { {{range $unusedrow,$row := $.Tab.DataRowIndex}}{{$headers := $.G.Types.AllFieldByName $.Tab.OriginalHeaderType}}
			{ {{range $col, $header := $headers}}{{if IsWrapFieldName $.G $.Tab $headers $row $col}}{{$header.FieldName}} = {{WrapTabValue $.G $.Tab $headers $row $col}}, {{end}}{{end}}},{{end}}
		}
		for _, value in pairs(g.{{$.Tab.HeaderType}}) do
			setmetatable(value, {
				__index = { {{range $fieldName, $value := DefaultField $.G $.Tab}}{{$fieldName}} = {{$value}},{{end}} },
			})
		end
		{{range $ii, $idx := GetIndicesByTable $.Tab}}
		-- {{$idx.Table.HeaderType}}
		g.{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}} = {}
		for _, rec in pairs(g.{{$idx.Table.HeaderType}}) do
			g.{{$idx.Table.HeaderType}}By{{$idx.FieldInfo.FieldName}}[rec.{{$idx.FieldInfo.FieldName}}] = rec
		end
		{{end}}
		return g
	end
}
`
const templateText_luatype = `-- Generated by github.com/davyxu/tabtoy
-- Version: {{.Version}}

return {
	init = function( g )
		{{range $sn, $objName := $.Types.EnumNames}}
		g.{{$objName}} = { {{range $fi,$field := $.Types.AllFieldByName $objName}}
			{{$field.FieldName}} = {{$field.Value}}, -- {{if not $field.Note}}{{$field.Name}}{{else}}{{ $field.Note}}{{end}} {{end}} {{range $fi,$field := $.Types.AllFieldByName $objName}}
			[{{$field.Value}}] = "{{$field.FieldName}}",{{end}}
		}{{end}}
		return g
	end
}
`
const templateText_LuaDoc = `-- Generated by github.com/davyxu/tabtoy
-- Version: {{.Version}}
-- 这个文件只是用来生成出注释代码。

return { {{range $sn, $objName := $.Types.StructNames}}
	{{$objName}} = { {{range $fi,$field := $.Types.AllFieldByName $objName}}
		{{$field.FieldName}} = "{{if not $field.Note}}{{$field.Name}}{{else}}{{ $field.Note}}{{end}}", {{end}}
	},
{{end}}
}
`
