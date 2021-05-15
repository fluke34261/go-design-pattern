package main

import "fmt"

func main() {
	gtr, _ := getCar("gtr")
	evolution, _ := getCar("evolution")

	fmt.Printf("GTR spec : %s %s %d", gtr.getBrandName(), gtr.getName(), gtr.getSpeed())
	fmt.Printf("Evolution spec : %s %s %d", evolution.getBrandName(), evolution.getName(), evolution.getSpeed())
}