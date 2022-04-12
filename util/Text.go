package util

type Book struct {
	name string
	age  int
}

func (this *Book) SetName(name string) {
	this.name = name
}
func (this *Book) GetName() string {
	return this.name
}

func (this *Book) SetAge(age int) {
	this.age = age
}
func (this *Book) GetAge() int {
	return this.age
}
