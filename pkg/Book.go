package pkg

import "fmt"

type Book struct {
	Title string
	Price int
}

func (b Book) Print() {
	fmt.Printf("title:%s, price:%d", b.Title, b.Price)
}
