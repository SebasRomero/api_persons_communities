package storage

import (
	"fmt"

	"github.com/sebasromero/api/model"
)

// Memory
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory returns a new Memory instance
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

//Create a new Person in memory
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	fmt.Println("before", len(m.Persons))
	m.currentID++
	m.Persons[m.currentID] = *person
	fmt.Println("after", len(m.Persons))
	return nil
}

// Update updates a person in the memory slice
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrPersonDoesNotExist)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete removes a person
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrPersonDoesNotExist)
	}

	delete(m.Persons, ID)

	return nil
}

// GetByID returns a person with the given ID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]

	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrPersonDoesNotExist)
	}

	return person, nil
}

// GetAll returns all the persons in memory
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}

func (m *Memory) GetID() int {
	return len(m.Persons) + 1
}
