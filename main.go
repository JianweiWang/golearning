package main

import (
	"datafile/golearning/magazine"
	"fmt"
)

func main() {
	sub := magazine.Subscriber{Name: "james", Rate: float64(0.8), Active: true}
	emp := magazine.Employee{Name: "jhon", Salary: float64(300)}
	addr := magazine.Address{Street: "liangzhu", City: "hangzhou", State: "zhejiang", PostCode: "0571"}
	fmt.Printf("%#v\n", sub)
	fmt.Printf("%#v\n", emp)
	fmt.Printf("%#v\n", addr)
}
