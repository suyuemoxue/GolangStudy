package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}

type Phones struct {
	Goods
	Brand string
}

type Books struct {
	Goods
	Write string
}

func (g *Goods) information() {
	fmt.Println("产品名:", g.Name)
	fmt.Println("价格:", g.Price, "元")
}

func (p *Phones) test() {
	fmt.Println("品牌:", p.Brand)
}

func (b *Books) test() {
	fmt.Println("作者:", b.Write)
}

func main() {
	phone := &Phones{
		Goods: Goods{
			Name:  "小米14",
			Price: 4999,
		},
		Brand: "小米",
	}

	book := &Books{
		Goods{"三国演义", 49},
		"罗贯中",
	}

	phone.information()
	phone.test()
	fmt.Println("----------------------------")
	book.information()
	book.test()
}
