package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para jsom
	p1 := produto{1, "Notebook", 1899.90, []string{"Promocao", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))
	fmt.Println(p1)

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"caneta","preco":9.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
