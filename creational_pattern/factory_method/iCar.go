package main

// product interface

type iCar interface {
	setBrand(brandName string)
	getBrand()	string

	setName(name string)
	getName() string

	setSpeed(speed int)
	getSpeed() int
}