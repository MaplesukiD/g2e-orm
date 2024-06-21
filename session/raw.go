package session

import (
	"database/sql"
	"g2e-orm/dialect"
	"g2e-orm/log"
	"g2e-orm/schema"
	"strings"
)

type Session struct {
	db *sql.DB
	//sql语句，包含占位符
	sql strings.Builder
	// 占位符对应值
	sqlParams []any
	//方言
	dialect dialect.Dialect
	//映射表
	refTable *schema.Schema
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlParams = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw 给session实例赋值
func (s *Session) Raw(sql string, param ...any) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlParams = append(s.sqlParams, param...)
	return s
}

// Exec 执行不需要返回结果集的sql语句
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlParams...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow 查询单行数据
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	return s.DB().QueryRow(s.sql.String(), s.sqlParams...)
}

// QueryRows 查询多行数据
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlParams...); err != nil {
		log.Error(err)
	}
	return
}
