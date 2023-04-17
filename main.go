package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

// Sem o asterisco (ponteiro), o método não altera o valor do atributo Price do objeto Order que chamou método em questão
func (o *Order) SetPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do SetPrice: ", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func main() {
	order := Order{
		ID:       "123",
		Price:    10.00,
		Quantity: 5,
	}
	order.SetPrice(20.00)
	fmt.Println("Preço total: ", order.getTotal()) //100
}
