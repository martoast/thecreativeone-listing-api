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
	PropertyModel := &models.Property{}
	utils.ParseBody(r, PropertyModel)
	b := PropertyModel.CreateProperty()
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PropertyId := vars["PropertyId"]
	ID, err := strconv.ParseInt(PropertyId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}
	property := models.DeleteProperty(ID)
	res, err := json.Marshal(property)
	if err != nil {
		http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	var updateProperty = &models.Property{}
	utils.ParseBody(r, updateProperty)
	vars := mux.Vars(r)
	PropertyId := vars["PropertyId"]
	ID, err := strconv.ParseInt(PropertyId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	propertyDetails, db := models.GetPropertyById(ID)

	propertyDetails.Price = updateProperty.Price
	propertyDetails.Description = updateProperty.Description
	propertyDetails.Images = updateProperty.Images
	propertyDetails.Sold = updateProperty.Sold
	propertyDetails.Bedrooms = updateProperty.Bedrooms
	propertyDetails.Bathrooms = updateProperty.Bathrooms
	propertyDetails.RentZestimate = updateProperty.RentZestimate
	propertyDetails.Zestimate = updateProperty.Zestimate
	propertyDetails.PropertyType = updateProperty.PropertyType
	propertyDetails.Zoning = updateProperty.Zoning
	propertyDetails.YearBuilt = updateProperty.YearBuilt
	propertyDetails.LotSize = updateProperty.LotSize
	propertyDetails.PricePerSquareFoot = updateProperty.PricePerSquareFoot
	propertyDetails.LivingArea = updateProperty.LivingArea

	db.Save(&propertyDetails)
	res, err := json.Marshal(propertyDetails)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
