package nexigo

import (
	"database/sql"
	// "fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type SqlDB struct {
	DB *sql.DB
}

func OpenDB() *SqlDB {
	db, _ := sql.Open(driverDB, connInfo)
	// db, _ := sql.Open("mssql", "Server=localhost;Database=test;User Id=sa;Password=123")
	// db, _ := sql.Open("postgres", "server=localhost user=wsugiri password=123 dbname=test")

	return &SqlDB{DB: db}
}

func (s *SqlDB) Execute(query string, args ...interface{}) {
	s.DB.Exec(query, args...)
}

func (s *SqlDB) ExecuteRows(query string, args ...interface{}) *sql.Rows {
	rows, _ := s.DB.Query(query, args...)
	return rows
}

func (s *SqlDB) QueryToList(query string, args ...interface{}) []interface{} {
	rows := s.ExecuteRows(query, args...)
	cols, _ := rows.Columns()

	val1 := make([]interface{}, len(cols))
	val2 := make([]interface{}, len(cols))
	list := make([]interface{}, 0)

	for rows.Next() {
		value := make(map[string]interface{})
		for i, _ := range cols {
			val2[i] = &val1[i]
		}
		rows.Scan(val2...)

		for i, _ := range cols {
			value[cols[i]] = val1[i]
		}

		list = append(list, value)
	}

	return list
}

func (s *SqlDB) CloseDB() {
	s.DB.Close()
}
