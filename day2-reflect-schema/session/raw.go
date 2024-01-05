package session

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
	"strings"
)

type Session struct { //三个成员变量
	db       *sql.DB
	sql      strings.Builder
	sqlVars  []interface{}
	dialect  dialect.Dialect
	refTable *schema.Schema
}

func New(db *sql.DB, dialect dialect.Dialect) *Session { //用于创建一个新的 Session 实例
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() { //于清空 Session 中的 SQL 查询语句和变量
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB { //返回当前 Session 中的数据库连接对象
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session { //用于构建 SQL 查询字符串
	s.sql.WriteString(sql)                   // 将传入的 SQL 字符串追加到 Session 结构体中的 sql 字段中
	s.sql.WriteString(" ")                   //在 SQL 字符串之后添加一个空格，用于分隔 SQL 语句的不同部分
	s.sqlVars = append(s.sqlVars, values...) //将传入的可变参数值追加到 Session 结构体中的 sqlVars 切片中。这个切片用于存储 SQL 语句中的参数值。
	return s
}

// Exec raw sql with sqlVars   （使用 sqlVars 执行原始 sql）
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil { //执行 SQL 语句。使用 s.DB().Exec 方法执行 s.sql.String() 中的 SQL 语句，同时传递参数 s.sqlVars...。
		// 如果执行过程中出现错误，将错误记录到日志中。
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db   （QueryRow 从 db 获取一条记录）
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...) //...表示把所有元素打散
}

// QueryRow gets a list of records from db  （QueryRow 从数据库获取记录列表）
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
