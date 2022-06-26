package main

import (
	"fmt"
	"go_db_fundamental/config"
	"go_db_fundamental/model"
	"go_db_fundamental/utils"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// fmt.Println("Test")

	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "postgres"
	// dbPassword := "12345678"
	// dbName := "db_enigma_shop"

	//urutan url koneksi ke db postgres
	// postgres://dbUser:dbPassword@dbHost:dbPort/dbName

	//tambain ?sslmode=disable abis dbName
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// db, err := sqlx.Connect("postgres", dsn)

	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	log.Println("connected")
	// }

	// defer func(db *sqlx.DB) {
	// 	err := db.Close()

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// }(db)

	// db.Close()

	//semua diatas diganti sama method & interface

	config := config.NewConfigDB()
	db := config.DbConn()

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			utils.IsError(err)
		}
	}(db)

	fmt.Println("Jalan")

	//INSERT
	//bikin insert jadi banyak, tinggal nambain [] dibelakang map

	// customer := []map[string]interface{}{
	// 	{
	// 		"id":      "C001",
	// 		"name":    "Selestin Sutisna",
	// 		"address": "Bandung",
	// 		"phone":   "028292992",
	// 		"email":   "bulan.sutina@gmail.com",
	// 		"saldo":   100000,
	// 	},
	// 	{
	// 		"id":      "C002",
	// 		"name":    "Bulan Sutisna",
	// 		"address": "Bandung",
	// 		"phone":   "028292992",
	// 		"email":   "bulan.sutina@gmail.com",
	// 		"saldo":   100000,
	// 	},
	// }

	//pake model

	//tambain id
	// id := uuid.New()
	//id = "C007" --> id.String()

	customer_model := model.Customer{
		Name:    "Lalala",
		Address: "Bandung",
		Phone:   "078956767",
		Email:   "diubah.unil@gmail.com",
		Id:      "C006",
	}

	// id2 := uuid.New()

	// customer_model2 := model.Customer{
	// 	Id:      id2.String(),
	// 	Name:    "Unil Kunya",
	// 	Address: "Bandung",
	// 	Phone:   "08159674566",
	// 	Email:   "kunya.unil@gmail.com",
	// 	Balance: 2300000,
	// }

	// _, err = db.NamedExec(utils.INSERT_CUSTOMER, customer_model)

	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println("insert successa")
	// }

	// With Prepare stetemtn
	//-> mengindari adanya sql injection
	// -> code lebnih effisien karena bisa digunakan berulang kali

	//-> database.Preparex(query)
	//-> stmt.MustExec(value)

	// stmt, err := db.Preparex(utils.UPDATE_CUSTOMER_PS)
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Println("Update is Success")
	// }

	// stmt.MustExec( customer_model.Name, customer_model.Address, customer_model.Phone, customer_model.Email, customer_model.Balance,customer_model.Id)

	// stmt.MustExec(customer_model2.Id, customer_model2.Name, customer_model2.Address, customer_model2.Phone, customer_model2.Email, customer_model2.Balance)

	//nambain update

	// stmt.MustExec(customer.Name, customer.Address, customer.Phone, customer.Email, customer.Id)

	//nambain delete

	customer01 := model.Customer{Id: "C006"}
	// customer02 := model.Customer{Id: "C002"}

	fmt.Println("Jalan2")

	stmt, err := db.Preparex(utils.DELETE_CUSTOMER_PS_SD)

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Delete is Success")
	}

	fmt.Println("Jalan3")

	stmt.MustExec(customer01.Id)

}
