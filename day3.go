package main

import (
	"fmt"
	"math"
)

type employee struct {
	name     string
	age      int
	position string
}

// lab1
func findOldestEmployee(employees []employee) employee {
	if len(employees) == 0 {
		return employee{}
	}
	oldest := employees[0]
	for _, employee := range employees {
		if employee.age > oldest.age {
			oldest = employee
		}
	}
	return oldest
}

// lab2
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Prameter() float64 {
	return 2 * math.Pi * c.Radius
}

// lab3
type Shape interface {
	Area() float64
	Prameter() float64
}
type Cricle struct {
	Radius float64
}

func (c Cricle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Cricle) Prameter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Prameter() float64 {
	return 2 * (r.Width + r.Height)
}

// lab4
type Vehicle interface {
	Start()
	Stop()
}
type Car struct {
	Name string
}

func (c Car) Start() {
	fmt.Printf("xe hoi %s dang chay\n", c.Name)
}
func (c Car) Stop() {
	fmt.Printf("xe hoi %s dang dung\n", c.Name)
}

type Motorbike struct {
	Name string
}

func (m Motorbike) Start() {
	fmt.Printf("xe may %s dang chay\n", m.Name)
}
func (m Motorbike) Stop() {
	fmt.Printf("xe may %s dang dung\n", m.Name)
}

// lab5
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// lab6
func increment(num *int) {
	*num++
}

// lab7
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// lab8
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func printtFinbonacci(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}

// lab9
func divide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Loi:%s\n", r)
		}
	}()
	if b == 0 {
		panic("K the chia cho 0")
	}
	return a / b
}

// lab10
func exFunc() {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello")
	panic("Panic xay ra")
}

// lab11
func sum(numbers ...int) int {
	tatal := 0
	for _, num := range numbers {
		tatal += num
	}
	return tatal
}

// lab12
func max(numbers ...int) int {
	maxValue := numbers[0]
	for _, num := range numbers {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}
func main() {
	employess := []employee{
		{"John", 25, "Software Engineer"},
		{"Jane", 30, "Senior Software Engineer"},
		{"Doe", 35, "Tech Lead"},
	}
	oldestEmployee := findOldestEmployee(employess)
	fmt.Printf("Nhan vien lon tuoi nhat la:%s\n", oldestEmployee.name)

	circle := Circle{Radius: 10}
	fmt.Printf("Dien tich hinh tron la:%.2f\n", circle.Area())
	fmt.Printf("Chu vi hinh tron la:%.2f\n", circle.Prameter())

	rectangle := Rectangle{Width: 10, Height: 20}
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Dien tich:%.2f\n Chu vi:%.2f\n", shape.Area(), shape.Prameter())
	}

	car := Car{Name: "Lamborghini"}
	motorbike := Motorbike{Name: "Honda"}
	vehicles := []Vehicle{car, motorbike}
	for _, vehicle := range vehicles {
		vehicle.Start()
		vehicle.Stop()
	}

	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%d, b=%d\n", a, b)

	num := 10
	increment(&num)
	fmt.Printf("num=%d\n", num)

	n := 10
	fmt.Printf("Giai thua cua %d la:%d\n", n, factorial(n))

	fmt.Printf("Day so fibonacci cua %d so dau tien la: ", n)
	printtFinbonacci(n)

	fmt.Printf("10/2 la:%d\n", divide(10, 2))
	fmt.Printf("10/0 la:%d\n", divide(10, 0))

	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Tong cua cac so la: %d\n", total)

	maxValue := max(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("So lon nhat la: %d\n", maxValue)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	exFunc()
}
