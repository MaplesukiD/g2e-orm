package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

// Dialect 方言
type Dialect interface {
	// DataTypeOf 将go语言类型转成数据库数据类型
	DataTypeOf(typ reflect.Value) string
	// TableExistSQL 来查看某个表是否存在的sql语句
	TableExistSQL(tableName string) (string, []any)
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
