package models

import "sync"

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var (
	Products = []Product{
		{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
		{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
		{ID: 3, Name: "kecap", Price: 12000, Stock: 20},
	}
	Mu sync.RWMutex
)

func FindProductByID(id int) (*Product, int) {
	for i := range Products {
		if Products[i].ID == id {
			return &Products[i], i
		}
	}
	return nil, -1
}

func GetNextID() int {
	maxID := 0
	for _, p := range Products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	return maxID + 1
}
