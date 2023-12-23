package setupDB

import (
	"database/sql"
	"fmt"
)

func Validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println("\nTransaction Rollback!")
	} else {
		fmt.Println("\nTransaction Successfully " + message + " Data!")
	}
}
