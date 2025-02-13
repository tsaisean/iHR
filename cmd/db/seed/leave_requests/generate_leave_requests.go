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

	leaveTypes := generateLeaveRequests()

	if err := db.DB.Create(&leaveTypes).Error; err != nil {
		log.Fatalf("Failed to create leave requests: %v", err)
	} else {
		log.Printf("Seed leave requests: %v", leaveTypes)
	}
}

func generateLeaveRequests() []model.LeaveRequest {
	return []model.LeaveRequest{
		{
			CreatorID:   1,
			EmployeeID:  1,
			ApproverID:  2,
			LeaveTypeID: 1,
			StartDate:   gofakeit.FutureDate(),
			EndDate:     gofakeit.FutureDate(),
			Duration:    1 * 8 * 60,
			Reason:      gofakeit.Breakfast(),
			CreatedAt:   gofakeit.Date(),
			UpdatedAt:   gofakeit.Date(),
			Status:      "active",
		}}
}
