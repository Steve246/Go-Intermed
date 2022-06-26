package utils

const (
	INSERT_CUSTOMER = `insert into customer(id,name,address,phone,email,saldo) values(:id, :name, :address, :phone, :email, :saldo)`

	INSERT_CUSTOMER_PS = `insert into customer(id,name,address,phone,email,saldo) values($1, $2, $3, $4, $5, $6)`
)
