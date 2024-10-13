package product

type Product struct {
	Name  string
	price float64
}

func NewProduct(name string, price float64) Product {
	return Product{
		Name:  name,
		price: price,
	}
}

func (p Product) GetName() string {
	return p.Name
}

func (p Product) getPrice() float64 {
	return p.price
}
