package transaction

import (
	"belajar-database/model"
	"belajar-database/setupDB"
	"belajar-database/setupQuery"
	"database/sql"
)

func AddUsers(users model.Users, tx *sql.Tx) {
	_, err := tx.Exec(setupQuery.InsertUser, users.Name, users.Email, users.Password, users.Role)
	setupDB.Validate(err, "Add New User", tx)
}
