package repo

import (
	"belajar-database/enrollment"
	"belajar-database/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateUserAdmin() {
	var roleSwitch int
	var role string
a:
	fmt.Print("\nMembuat Akses Users Baru Sebagai \n>> 1. Admin\n>> 2. Kasir\n")
	fmt.Print(strings.Repeat("=", 11), " :")
	fmt.Scanln(&roleSwitch)
	switch roleSwitch {
	case 1:
		role = "admin"
	case 2:
		role = "kasir"
	default:
		fmt.Println("\nRole tidak tersedia!!")
		goto a
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println("|Biodata Users|")
	fmt.Println(strings.Repeat("-", 15))

	fmt.Print("\nMasukkan Nama Users     : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Masukkan Email Users    : ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Masukkan Password Users : ")
	scanner.Scan()
	password := scanner.Text()
	newUsers := model.Users{Name: nama, Email: email, Password: password, Role: role}
	enrollment.EnrollCreateUsers(newUsers)
}
