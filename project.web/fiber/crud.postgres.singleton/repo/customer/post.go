package rcustomer

import (
	db "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.singleton/psql"
	mcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/models/customer"
)
// Post ..
func Post(db customer mcustomer.Customer) error {
	sqlexec := `insert 
	into ad_customer(
		imp_ip,
		imp_user_create, 
		imp_user_update,
		imp_nome 
	)values($1,$2,$3,$4)`
	_, err := db.Exec(sqlexec, //TIPOS DE CADA CAMPO
		customer.ImpIp,         //string
		customer.ImpUserCreate, //uuid
		customer.ImpUserUpdate, //uuid
		customer.ImpNome,       //string
	)
	if err != nil {
		return err
	}
	return nil

}
