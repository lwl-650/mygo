package model

import (
	"fmt"
	"strconv"
)

type Book struct {
	name string
	num  int
}

func (book *Book) SetBook(name string, num int) {
	book.name = name
	book.num = num
}
func (book *Book) Siy() {
	strings := strconv.Itoa(book.num)
	fmt.Println("hellow,i" + book.name + "=====" + strings)
}

func (book Book) Siy2() {
	strings := strconv.Itoa(book.num)
	fmt.Println("hellow,i" + book.name + "=====" + strings)
}
