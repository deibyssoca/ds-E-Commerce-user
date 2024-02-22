package bd

import (
	"fmt"

	"github.com/deibyssoca/models"
	"github.com/deibyssoca/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Start SignUp")
	err := DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DateTimeMySQL() + "')"
	fmt.Println(query)

	_, err = Db.Exec(query)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("SignUp => OK!!")
	return nil
}
