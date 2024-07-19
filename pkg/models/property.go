package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Property struct {
	gorm.Model
	Address            string   `json:"address"`
	Price              *float64 `json:"price"`                 // DECIMAL(10, 2)
	Description        *string  `json:"description"`           // TEXT
	Images             *string  `json:"images"`                // JSON array of image URLs or a string of delimited URLs
	Sold               *bool    `json:"sold"`                  // BOOLEAN
	Bedrooms           *int     `json:"bedrooms"`              // INT
	Bathrooms          *float64 `json:"bathrooms"`             // DECIMAL(3, 1)
	RentZestimate      *float64 `json:"rent_zestimate"`        // DECIMAL(10, 2)
	Zestimate          *float64 `json:"zestimate"`             // DECIMAL(10, 2)
	PropertyType       *string  `json:"property_type"`         // VARCHAR(255)
	Zoning             *string  `json:"zoning"`                // VARCHAR(255)
	YearBuilt          *int     `json:"year_built"`            // INT
	LotSize            *int     `json:"lot_size"`              // INT
	PricePerSquareFoot *float64 `json:"price_per_square_foot"` // DECIMAL(10, 2)
	LivingArea         *int     `json:"living_area"`           // INT
	PurchasePrice      *float64 `json:"purchase_price"`        // DECIMAL(10,2)
	BalanceToClose     *float64 `json:"balance_to_close"`      // DECIMAL(10,2)
	MonthlyHoldingCost *float64 `json:"monthly_holding_cost"`  // DECIMAL(10,2)
	InterestRate       *float64 `json:"interest_rate"`         // DECIMAL(10,2)
	NearbyHospitals    *string  `json:"nearby_hospitals"`      // JSON array of nearby_hospitals
	NearbySchools      *string  `json:"nearby_schools"`        // JSON array of nearby_schools
	NearbyHomes        *string  `json:"nearby_homes"`          // JSON array of nearby_homes
	PriceHistory       *string  `json:"price_history"`         // JSON array of price history
	TaxHistory         *string  `json:"tax_history"`           // JSON array of tax history
	MonthlyHoaFee      *int     `json:"monthly_hoa_fee"`       // INT
	AssistedLiving     *bool    `json:"assisted_living"`       // BOOLEAN
}

func init() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Property{})

	// DeleteProperty(8)

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

func GetPaginatedProperties(limit int, offset int, sold *bool, assisted_living *bool) ([]Property, int64) {
	var properties []Property
	var total int64

	query := db.Model(&Property{})

	if sold != nil {
		query = query.Where("sold = ?", *sold)
	}

	if assisted_living != nil {
		query = query.Where("assisted_living = ?", *assisted_living)

	}

	query.Count(&total)
	query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&properties)

	return properties, total
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
			Address:            "4949 Corrado Ave, Ave Maria, FL 34142",
			Price:              newFloat64(300000),
			Description:        newString("Beautiful family home in a quiet neighborhood."),
			Images:             newString("[\"https://static.tildacdn.com/stor3630-6334-4663-b532-393032356238/65960768.jpg\", \"https://static.tildacdn.com/stor3663-3339-4534-b332-393563363363/61347039.jpg\"]"),
			Sold:               newBool(false),
			Bedrooms:           newInt(3),
			Bathrooms:          newFloat64(2.5),
			RentZestimate:      newFloat64(2500),
			Zestimate:          newFloat64(300000),
			PropertyType:       newString("Single Family"),
			Zoning:             newString("R-1:SINGLE FAM-RES"),
			YearBuilt:          newInt(1990),
			LotSize:            newInt(5000),
			LivingArea:         newInt(3000),
			PricePerSquareFoot: newFloat64(300),
			PurchasePrice:      newFloat64(300000),
			BalanceToClose:     newFloat64(10000),
			MonthlyHoldingCost: newFloat64(5000),
			InterestRate:       newFloat64(300),
			NearbyHospitals:    newString("[\"Hospital A\", \"Hospital B\"]"),
			NearbySchools:      newString("[\"School A\", \"School B\"]"),
			NearbyHomes:        newString("[\"Home A\", \"Home B\"]"),
			PriceHistory:       newString("[{\"date\": \"2022-01-01\", \"price\": 295000}, {\"date\": \"2023-01-01\", \"price\": 300000}]"),
			TaxHistory:         newString("[{\"year\": 2022, \"tax\": 3500}, {\"year\": 2023, \"tax\": 3600}]"),
			MonthlyHoaFee:      newInt(1000),
			AssistedLiving:     newBool(false),
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
