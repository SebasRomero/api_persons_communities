package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/sebasromero/api/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allowed.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	data.ID = p.storage.GetID()
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "Problem decoding data", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "There's a problem creating the person.", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person created succesfully.", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method not allowed.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "Id person must be a number.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	data.ID = ID
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Json decode error", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "There's a problem updating the person.", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person updated succesfully", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method not allowed.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "Id person must be a number.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrPersonDoesNotExist) {
		response := newResponse(Error, "Id does not exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Error deleting the person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person deleted succesfully", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method not allowed.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "Id person must be a number.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data, err := p.storage.GetByID(ID)
	if err != nil {
		response := newResponse(Error, "Error getting the person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method not allowed.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Error getting all the persons.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
