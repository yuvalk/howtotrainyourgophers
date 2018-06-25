type Vehicle struct {
	wheels int
}

type Bike struct {
	Vehicle      // embedding
	pedals int
}

var b Bike
b.wheels // b.Vehicle.wheels
