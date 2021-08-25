package main

import (
	"fmt"

	"github.com/seungjelee/goCoinProject/person"
)

func main() {
	sj := person.Person{}
	sj.SetDetails("seungje", 22)
	fmt.Println(sj)
}

/*
type person struct {
	name string
	age  int
}

func (p person) sayKoreanAge() {
	fmt.Printf("내 한국 나이는 %d입니다", p.age+1)
}

func (p person) sayHello(loop int) {
	for i := 0; i < loop; i++ {
		fmt.Printf("Hello! My name is %s and I'm %d\n", p.name, p.age)
	}
}

func main() {
	sj := person{"seungje", 33}
	sj.sayHello(3)
	sj.sayKoreanAge()
}
*/
/*
const call_name string = "cannot change"

func main() {
	a := 2
	b := &a
	a = 12
	fmt.Println(a, *b)
}
*/
/*
func plus(a ...int) int {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum
}
func main() {
	foods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods)
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)

	for _, dish := range foods {
		fmt.Println(dish)
	}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
	// name := "it works!"
	// age := 12
	// result, name := plus(2, 2, "seungje")
	// sum := plus(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// biSum := fmt.Sprintf("%b\n", sum)
	// fmt.Println(sum, biSum)
}
*/
