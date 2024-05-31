package main

import "fmt"

func main() {

	myContractorEmployee := contractor{name: "Ash", hoursPerYear: 10, hourlyPay: 5}
	fmt.Println(myContractorEmployee.getSalary())

	myFullTimeEmployee := fullTime{name: "Ash", salary: 20}
	fmt.Println(myFullTimeEmployee.getSalary())
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}
