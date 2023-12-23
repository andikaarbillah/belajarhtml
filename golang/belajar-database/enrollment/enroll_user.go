package enrollment

import (
	"belajar-database/model"
	"belajar-database/setupDB"
	"belajar-database/transaction"
	"fmt"
	"strings"
)

func EnrollCreateUsers(Users model.Users) {
 db := setupDB.ConnectDb()
 defer db.Close()

 tx, err := db.Begin()
 if err != nil{
	panic(err)
 }
transaction.AddUsers(Users,tx)

err = tx.Commit()
if err != nil{
	panic(err)
}else{
	fmt.Println(strings.Repeat("-",23))
	fmt.Println("|Transaction  Commited|")
	fmt.Println(strings.Repeat("-",23))

}
}

