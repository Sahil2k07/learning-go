package basics1

import "fmt"

type variant struct {
	id   string
	name string
	data struct {
		inDemand bool
		emi      int
	}
}

// Nested Structs
type car struct {
	name    string
	model   string
	price   float64
	variant variant
}

// Embedded Structs
type Vehicle struct {
	car
	isAvailable bool
}

func structs() {
	swift := car{
		name:  "Swift",
		model: "Hatchback",
		price: 600000.0,
		variant: variant{
			id:   "df4df",
			name: "Base",
			data: struct {
				inDemand bool
				emi      int
			}{
				inDemand: true,
				emi:      20000,
			},
		}}

	fmt.Println(swift)

	//Access values just like javascript objects
	fmt.Println(swift.name)

	vehicle := Vehicle{
		car: car{
			name:  "Scorpio",
			model: "Top",
		},
		isAvailable: true,
	}

	// Anonymours Structs
	info := struct {
		email    string
		password string
	}{
		email:    "random@gmail.com",
		password: "password",
	}

	// Possible in case of Embedded Structs
	fmt.Println(vehicle, vehicle.name)

	fmt.Println(info)
}
