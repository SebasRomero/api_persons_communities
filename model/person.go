package model

// Community is a community struct
type Community struct {
	Name string `json:"name"`
}

// Communities is a communities slice
type Communities []Community

// Person is a person struct
type Person struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

type Persons []Person
