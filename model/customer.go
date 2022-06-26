package model

type Customer struct {
	Id      string
	Name    string
	Address string
	Phone   string
	Email   string
	Balance int `db:"saldo"`
	// isStatus int `db:"is_status"`
	IsStatus int `db:"is_status"`
}
