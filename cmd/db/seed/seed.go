package main

import (
	"encoding/csv"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"iHR/config"
	"iHR/db"
	"iHR/db/model"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.Database.Host = "localhost"
	db.Connect(&cfg.Database)

	employees := generateRandomEmployees(100)

	if err := db.DB.Create(&employees).Error; err != nil {
		log.Fatalf("Failed to create employees: %v", err)
	} else {
		log.Printf("Seed employees: %v", employees)
	}
}

func generateRandomEmployees(count int) []model.Employee {
	employees := make([]model.Employee, 0, count)
	for i := 0; i < count; i++ {
		supervisorID := uint(rand.Intn(count))

		employees = append(employees, model.Employee{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Email:        gofakeit.Email(),
			Position:     gofakeit.JobTitle(),
			SupervisorID: &supervisorID,
			CreatedAt:    gofakeit.Date(),
		})
	}

	return employees
}

func generateFromFile() {
	file, err := os.Open("db/model/employee_seed.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse CSV: %v", err)
	}

	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		employee := model.Employee{
			FirstName: strings.TrimSpace(record[0]),
			LastName:  strings.TrimSpace(record[1]),
			Email:     strings.TrimSpace(record[2]),
			Position:  strings.TrimSpace(record[3]),
		}

		if err := db.DB.Create(&employee).Error; err != nil {
			log.Fatalf("Failed to create employee: %v", err)
		} else {
			log.Printf("Seed employee: %v", employee)
		}
	}

	fmt.Println("Seeding successfully!")
}
