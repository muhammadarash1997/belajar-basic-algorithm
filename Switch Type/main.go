package main

import "fmt"

type Timestamp struct {
	Seconds int64
	Nanos int32
}

//////////////////////////////

type Vehicles []Vehicle
type Vehicle struct {
	No int32
	Name string
}

//////////////////////////////

type VehiclesNearby []VehicleNearby
type VehicleNearby struct {
	No int32
	Name string
	Eta Timestamp
}

func main() {
	yes := false

	mapvehicles := make(map[int32]interface{})

	if yes {
		v := Vehicles{
			{
				No: 1,
				Name: "Pajero",
			},
			{
				No: 2,
				Name: "Alphard",
			},
		}
		mapvehicles[1] = v
	} else {
		vn := VehiclesNearby{
			{
				No: 1,
				Name: "Carry",
			},
			{
				No: 2,
				Name: "Xenia",
			},
		}
		mapvehicles[2] = vn
	}

	if mapvehicles[1] != nil {
		fmt.Println("Not Nil")
	}

	switch v := mapvehicles[1].(type) {
	case Vehicles:
		fmt.Println("Vehicles Type")
		for _, v := range v {
			fmt.Println(v)
		}
	case VehiclesNearby:
		fmt.Println("Vehicles Type")
		for _, v := range v {
			fmt.Println(v)
		}
	}
}