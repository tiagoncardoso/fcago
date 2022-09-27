package main

import (
	"aceleracao/golang/internal/order/infra/database"
	"aceleracao/golang/internal/order/usecase"
	"database/sql"
)

// https://www.youtube.com/watch?v=YqqzKsJEFRM

func main() {
	db, err := sql.Open("mysql", "root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculatePriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    "12",
		Price: 100,
		Tax:   10,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	println(output.FinalPrice)
}
