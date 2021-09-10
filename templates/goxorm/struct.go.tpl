package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
    "easygoadmin/utils"
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

// 根据条件查询单条数据
func (r *{{Mapper .Name}}) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *{{Mapper .Name}}) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *{{Mapper .Name}}) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *{{Mapper .Name}}) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&{{Mapper .Name}}{})
}

//批量删除
func (r *{{Mapper .Name}}) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&{{Mapper .Name}}{})
}
{{end}}

