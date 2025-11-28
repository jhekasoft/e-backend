package models

type Pet struct {
	Name string
	Age  float64 // in years
	Type string
}

type PetMainValues struct {
	Food   int
	Water  int
	Rest   int
	Joy    int
	Health int
}

type PetSecondaryValues struct {
	IsResting bool
}

type HouseValues struct {
	IsHeaped bool
}

type TamagluchiState struct {
	Pet       Pet
	Main      PetMainValues
	Secondary PetSecondaryValues
	House     HouseValues
}
