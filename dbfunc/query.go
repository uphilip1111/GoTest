package dbfunc

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func QueryBySQL() {
	db, err := sql.Open("mysql", "root:admin@/project")
	//db.SetMaxIdleConns(4000)
	//db.SetMaxOpenConns(3500)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err1 := db.Query(`select * from user_account`)

	if err1 != nil {
		panic(err1.Error())
	}
	defer rows.Close()

	colNames, _ := rows.Columns()

	var cols = make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		cols[i] = new(interface{})
	}
}
