package nexigo

import (
	"fmt"
)

type SqlDB struct {
}

func OpenDB() *SqlDB {
	return &SqlDB{}
}

func (s *SqlDB) Execute(query string) {
	fmt.Println("Execute", query)
}
