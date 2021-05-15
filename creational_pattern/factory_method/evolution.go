package main

type evolution struct {
	car
}

func newEvolution() iCar {
	return &evolution{
		car : car{
			brandName: "Mitsubishi",
			name: "Evolution X",
			speed: 400
		},
	}
}