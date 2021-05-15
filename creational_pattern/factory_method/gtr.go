package main

type gtr struct {
	car
}

func newGtr() iCar {
	return &gtr{
		car : car{
			brandName: "Nissan",
			name: "GTR 50 year",
			speed: 500
		},
	}
}