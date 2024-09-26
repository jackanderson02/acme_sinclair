package model

type User struct {
    ID   int    `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
}

type Product struct{
    ID   int    `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
    Price int `json:"price" db:"price"`
    Stock_count int `json:"stock_count" db:"stock_count"` 
}