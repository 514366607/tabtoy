package compiler

import (
	"github.com/davyxu/tabtoy/v3/model"
	"github.com/davyxu/tabtoy/v3/table"
	"path/filepath"
	"sort"
	"strings"
)

func loadIndexData(tab *model.DataTable, symbols *model.TypeTable) (pragmaList []*table.TablePragma) {

	for row := 1; row < len(tab.Rows); row++ {

		var pragma table.TablePragma
		if !model.ParseRow(&pragma, tab, row, symbols) {
			continue
		}

		if pragma.TableMode == table.TableMode_Type {
			pragma.TableType = "TableField"
		}

		if pragma.TableType == "" {

			_, name := filepath.Split(pragma.TableFileName)

			pragma.TableType = strings.TrimSuffix(name, filepath.Ext(pragma.TableFileName))
		}

		pragmaList = append(pragmaList, &pragma)
	}

	return
}

func LoadIndexTable(globals *model.Globals, fileName string) error {

	if fileName == "" {
		return nil
	}

	tabs, err := LoadDataTable(globals.IndexGetter, fileName, "TablePragma")

	if err != nil {
		return err
	}

	var pragmaList []*table.TablePragma

	for _, tab := range tabs {

		ResolveHeaderFields(tab, "TablePragma", globals.Types)

		pragmaList = append(pragmaList, loadIndexData(tab, globals.Types)...)
	}

	// 按表类型排序，保证类型表先读取
	sort.Slice(pragmaList, func(i, j int) bool {
		a := pragmaList[i]
		b := pragmaList[j]

		if a.TableMode != b.TableMode {
			return a.TableMode < b.TableMode
		}

		if a.TableType != b.TableType {
			return a.TableType < b.TableType
		}

		return a.TableFileName < b.TableFileName

	})

	globals.IndexList = pragmaList

	return nil
}