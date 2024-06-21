package schema

import (
	"g2e-orm/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	//字段名
	Name string
	//类型
	Type string
	//约束条件
	Tag string
}

type Schema struct {
	//被映射对象的实例
	Model any
	//表名
	Name string
	//字段
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func Parse(model any, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	schema := &Schema{
		Model:    model,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("g2e-orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
