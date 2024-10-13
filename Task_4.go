package main

import "fmt"

type Vehicle struct {
	brand string
	model string
	year  int
}

func (v Vehicle) GetInfo() string {
	return fmt.Sprintf("Brand: %s, Model: %s, Year: %d", v.brand, v.model, v.year)
}

type Car struct {
	Vehicle
	numberOfDoors int
}

func (v *Vehicle) SetModel(model string) {
	v.model = model
}

func (c Car) GetCarInfo() string {
	return fmt.Sprintf("%s, Number of doors: %d", c.GetInfo(), c.numberOfDoors)
}

func main() {
	myVehicle := Vehicle{
		brand: "Toyota",
		model: "Corolla",
		year:  2020,
	}

	myCar := Car{
		Vehicle:       myVehicle,
		numberOfDoors: 4,
	}

	fmt.Println("Vehicle Info:", myVehicle.GetInfo())
	fmt.Println("Car Info:", myCar.GetCarInfo())

	myVehicle.SetModel("Camry")
	fmt.Println("Updated Vehicle Info:", myVehicle.GetInfo())
}
