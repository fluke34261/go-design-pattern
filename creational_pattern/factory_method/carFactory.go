package main

import "fmt"

func getCar(carName string) (iCar, error) {
	if carName == "gtr" {
		return newGtr(), nil
	}
	if carName == "evoltionx" {
		return newEvolution(), nil
	}

	return nil, fmt.Errorf("Wrong car name")
}