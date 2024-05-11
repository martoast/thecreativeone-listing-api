package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Property struct {
	gorm.Model
	Address         string   `json:"address"`
	Price           *float64 `json:"price,omitempty"`             // DECIMAL(10, 2)
	DownPayment     *float64 `json:"down_payment,omitempty"`      // DECIMAL(10, 2)
	TotalPrice      *float64 `json:"total_price,omitempty"`       // DECIMAL(10, 2)
	Interest        *float64 `json:"interest,omitempty"`          // DECIMAL(5, 2)
	MonthlyPayment  *float64 `json:"monthly_payment,omitempty"`   // DECIMAL(10, 2)
	Description     *string  `json:"description,omitempty"`       // TEXT
	ARV             *float64 `json:"arv,omitempty"`               // DECIMAL(10, 2) After Repair Value
	Benefits        *string  `json:"benefits,omitempty"`          // TEXT
	Images          *string  `json:"images,omitempty"`            // JSON array of image URLs or a string of delimited URLs
	Sold            *bool    `json:"sold,omitempty"`              // BOOLEAN
	Bedrooms        *int     `json:"bedrooms,omitempty"`          // INT
	Bathrooms       *float64 `json:"bathrooms,omitempty"`         // DECIMAL(3, 1)
	SquareFootage   *int     `json:"square_footage,omitempty"`    // INT
	RentZestimate   *float64 `json:"rent_zestimate,omitempty"`    // DECIMAL(10, 2)
	PropertyType    *string  `json:"property_type,omitempty"`     // VARCHAR(255)
	ParkingDetails  *string  `json:"parking_details,omitempty"`   // VARCHAR(255)
	YearBuilt       *int     `json:"year_built,omitempty"`        // INT
	LotSize         *int     `json:"lot_size,omitempty"`          // INT
	MortgageBalance *float64 `json:"mortgage_balance,omitempty"`  // DECIMAL(10, 2)
	InterestRate    *float64 `json:"interest_rate,omitempty"`     // DECIMAL(5, 2)
	PITI            *float64 `json:"piti,omitempty"`              // Principal, Interest, Taxes, Insurance (DECIMAL(10, 2))
	ExitROIStrategy *string  `json:"exit_roi_strategy,omitempty"` // VARCHAR(255)
	EstimateROI     *float64 `json:"estimate_roi,omitempty"`      // ROI/Cash on Cash (DECIMAL(5, 2))
	MonthlyCashFlow *float64 `json:"monthly_cash_flow,omitempty"` // DECIMAL(10, 2)
	EquityBuildup   *float64 `json:"equity_buildup,omitempty"`    // DECIMAL(10, 2)
}

func init() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Property{})

	// Seed data if needed
	var count int64
	db.Model(&Property{}).Count(&count)
	if count == 0 {
		SeedProperties()
	}
}

func GetAllProperties() []Property {
	var Properties []Property
	db.Find(&Properties)
	return Properties
}

func GetPropertyById(ID int64) (*Property, *gorm.DB) {
	var getProperty Property
	db := db.Where("ID=?", ID).Find(&getProperty)
	return &getProperty, db
}

func (b *Property) CreateProperty() *Property {
	db.Create(&b)
	return b

}

func DeleteProperty(ID int64) Property {
	var property Property
	db.Where("ID = ?", ID).Delete(&property)
	return property
}

func SeedProperties() {
	// Example set of properties to seed
	properties := []Property{
		{
			Address:         "4949 Corrado Ave, Ave Maria, FL 34142",
			Price:           newFloat64(300000),
			DownPayment:     newFloat64(60000),
			TotalPrice:      newFloat64(320000),
			Interest:        newFloat64(5.0),
			MonthlyPayment:  newFloat64(1500),
			Description:     newString("Beautiful family home in a quiet neighborhood."),
			ARV:             newFloat64(350000),
			Benefits:        newString("Close to schools and shopping centers, high appreciation potential."),
			Images:          newString("[\"https://static.tildacdn.com/stor3630-6334-4663-b532-393032356238/65960768.jpg\", \"https://static.tildacdn.com/stor3663-3339-4534-b332-393563363363/61347039.jpg\"]"),
			Sold:            newBool(false),
			Bedrooms:        newInt(3),
			Bathrooms:       newFloat64(2.5),
			SquareFootage:   newInt(1500),
			RentZestimate:   newFloat64(2500),
			PropertyType:    newString("Single Family"),
			ParkingDetails:  newString("2 car garage"),
			YearBuilt:       newInt(1990),
			LotSize:         newInt(5000),
			MortgageBalance: newFloat64(200000),
			InterestRate:    newFloat64(3.5),
			PITI:            newFloat64(2000),
			ExitROIStrategy: newString("Sell after two years"),
			EstimateROI:     newFloat64(10.5),
			MonthlyCashFlow: newFloat64(500),
			EquityBuildup:   newFloat64(25000),
		},
	}

	for _, property := range properties {
		db.Create(&property)
	}
}

func newFloat64(v float64) *float64 { return &v }
func newString(s string) *string    { return &s }
func newInt(i int) *int             { return &i }
func newBool(b bool) *bool          { return &b }
