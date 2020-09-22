package user

import (
	"log"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/model/usermodel"
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/psql"
)

func InsertUser(user usermodel.User) error {
	sqlStatement := `INSERT INTO adphone (name, location, age) VALUES ($1, $2, $3)`

	Db := pg.Connect()
	err := Db.QueryRow(sqlStatement, user.Name, user.Location, user.Age)
	if err != nil {
		log.Println(err)
	}
	return nil
}