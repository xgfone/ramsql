package main

import (
	"database/sql"
	"fmt"

	"github.com/xgfone/ramsql/cli"
	_ "github.com/xgfone/ramsql/driver"
)

func main() {
	db, err := sql.Open("ramsql", "")
	if err != nil {
		fmt.Printf("Error : cannot open connection : %s\n", err)
		return
	}
	cli.Run(db)
}
