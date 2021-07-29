package entity

type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int	`json:"stock"`
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock <= 3{
		status = "(Stok Almost 0)"
	} else if p.Stock < 10{
		status = "(Stok Limited)"
	} 
	return status
}
