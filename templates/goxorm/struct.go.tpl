package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
    "easygoadmin/library/db"
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

func ({{Mapper .Name}}) TableName() string {
    return "sys_{{$table.Name}}"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *{{Mapper .Name}}) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *{{Mapper .Name}}) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *{{Mapper .Name}}) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *{{Mapper .Name}}) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
{{end}}

