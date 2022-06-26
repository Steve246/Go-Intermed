package utils

const (
	INSERT_CUSTOMER = `insert into customer(id,name,address,phone,email,saldo) values(:id, :name, :address, :phone, :email, :saldo)`

	INSERT_CUSTOMER_PS = `insert into customer(id,name,address,phone,email,saldo) values($1, $2, $3, $4, $5, $6)`

	UPDATE_CUSTOMER_PS = `update customer set name=$1, address=$2, phone=$3, email=$4 where id=$5`
)
