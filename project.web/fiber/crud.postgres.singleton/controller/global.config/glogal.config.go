package config
import db "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.singleton/psql"

type dbStruct struct {
	db db.Connect()
}
