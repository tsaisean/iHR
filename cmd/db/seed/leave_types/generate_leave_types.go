package main

import (
	"iHR/config"
	"iHR/repositories/db"
	"iHR/repositories/model"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.Database.Host = "localhost"
	db.Connect(&cfg.Database)

	leaveTypes := generateLeaveTypes()

	if err := db.DB.Create(&leaveTypes).Error; err != nil {
		log.Fatalf("Failed to create leave types: %v", err)
	} else {
		log.Printf("Seed leave types: %v", leaveTypes)
	}
}

func generateLeaveTypes() []model.LeaveType {
	return []model.LeaveType{
		{
			Name:        "Annual Leave",
			Description: "General time off usually accrued throughout the year.",
		},
		{
			Name:        "Sick Leave",
			Description: "Time off specifically for illness or medical appointments.",
		},
		{
			Name:        "Maternity / Paternity / Parental Leave",
			Description: "Granted to new parents.",
		},
		{
			Name:        "Compassionate Leave",
			Description: "Granted for events like the death or serious illness of a family member.",
		},
		{
			Name:        "Military Leave",
			Description: "Temporarily returning to active military service.",
		}}
}
