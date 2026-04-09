package main

import (
	"fmt"
	"game-app/entity"
	"game-app/repository/mysql"
)

func main() {

}

func testUserMySqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "0913",
		Name:        "Morteza Mollaiee",
	})

	if err != nil {
		fmt.Println(err)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println(err)
	}
	if isUnique {
		fmt.Println("user is unique")
	} else {
		fmt.Println("user is not unique")
	}
}
