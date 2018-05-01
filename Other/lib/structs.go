package lib

//inheritance in golang
type Person struct {
	Name string
	Age  int
}
type Author struct {
	Person         *Person
	PublishedBooks []string
}

func (p *Person) IngrementAge() int {
	p.Age++
	return p.Age
}
func NewPerson(name string, age int) *Person{
	return &Person{
		Name:name,
		Age: age,
	}
}
