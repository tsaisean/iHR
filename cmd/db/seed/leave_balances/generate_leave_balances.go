package main

import (
	"github.com/brianvoe/gofakeit/v7"
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

	leaveTypes := generateLeaveBalances()

	if err := db.DB.Create(&leaveTypes).Error; err != nil {
		log.Fatalf("Failed to create leave balances: %v", err)
	} else {
		log.Printf("Seed leave balances: %v", leaveTypes)
	}
}

func generateLeaveBalances() []model.LeaveBalances {
	return []model.LeaveBalances{
		{
			EmployeeID:  1,
			LeaveTypeID: 1,
			Allocated:   10 * 8 * 60, // days*hours*minutes
			StartDate:   gofakeit.Date(),
			Status:      "active",
		}}
}
