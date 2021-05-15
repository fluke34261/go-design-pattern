package main

// concrete product

type car struct {
	name string
	brandName string
	speed int
}

func (c *car) setBrandName(brandName string) {
	c.brandName = brandName
}

func (c *car) getBrandName() string {
	return c.brandName
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) setSpeed(speed int) {
	c.speed = speed
}

func (c *car) getSpeed() int {
	return c.speed
}