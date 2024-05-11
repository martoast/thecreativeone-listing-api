package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"api/pkg/models"
	"api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewProperty models.Property

func GetProperties(w http.ResponseWriter, r *http.Request) {
	newProperties := models.GetAllProperties()

	res, err := json.Marshal(newProperties)
	if err != nil {
		http.Error(w, "Failed to encode properties as JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPropertyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PropertyId := vars["PropertyId"]
	ID, err := strconv.ParseInt(PropertyId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	PropertyDetails, _ := models.GetPropertyById(ID)
	res, _ := json.Marshal(PropertyDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProperty(w http.ResponseWriter, r *http.Request) {
	BookModel := &models.Property{}
	utils.ParseBody(r, CreateProperty)
	b := BookModel.CreateProperty()
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	propertyId := vars["propertyId"]
	ID, err := strconv.ParseInt(propertyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteProperty(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	var updateProperty = &models.Property{}
	utils.ParseBody(r, updateProperty)
	vars := mux.Vars(r)
	propertyId := vars["propertyId"]
	ID, err := strconv.ParseInt(propertyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	propertyDetails, db := models.GetPropertyById(ID)
	if updateProperty.Address != "" {
		propertyDetails.Address = updateProperty.Address
	}

	db.Save(&propertyDetails)
	res, err := json.Marshal(propertyDetails)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
