package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"iHR/repositories/db"
	. "iHR/repositories/model"
	"log"
	"time"
)

//go:generate mockery --all --output=./mocks
type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error)
	GetAllEmployeesAfter(ctx context.Context, id int, pageSize int) ([]Employee, error)
	GetAllEmployeesFrom(ctx context.Context, offset int, pageSize int) ([]Employee, error)
	GetEmployeeByID(id uint) (*Employee, error)
	UpdateEmployeeByID(ctx context.Context, id uint, updated *Employee) (*Employee, error)
	DeleteEmployee(ctx context.Context, id uint) error
	GetTotal() (int, error)
}

type EmployeeRepo struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewEmployeeRepo(db *gorm.DB, redis *redis.Client) *EmployeeRepo {
	return &EmployeeRepo{
		db:    db,
		cache: redis,
	}
}

func (r *EmployeeRepo) CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error) {
	if err := r.db.Create(employee).Error; err != nil {
		return nil, err
	}
	r.deleteEmployeeCache(ctx)
	return employee, nil
}

// GetAllEmployeesAfter Support cache
func (r *EmployeeRepo) GetAllEmployeesAfter(ctx context.Context, id int, pageSize int) ([]Employee, error) {
	var employees []Employee

	cacheKey, employees, err := r.getFromCache(ctx, id, 0, pageSize)
	if err != nil {
		return nil, err
	} else if employees != nil {
		return employees, nil
	}

	query := r.db.Limit(pageSize).Where("id > ?", id)
	if err := query.Find(&employees).Error; err != nil {
		return nil, err
	}

	if employees != nil {
		r.setCache(ctx, cacheKey, employees)
	}

	return employees, nil
}

// GetAllEmployeesFrom Support cache
func (r *EmployeeRepo) GetAllEmployeesFrom(ctx context.Context, offset int, pageSize int) ([]Employee, error) {
	var employees []Employee

	cacheKey, employees, err := r.getFromCache(ctx, 0, offset, pageSize)
	if err != nil {
		return nil, err
	} else if employees != nil {
		return employees, nil
	}

	query := r.db.Limit(pageSize).Offset(offset)
	if err := query.Find(&employees).Error; err != nil {
		return nil, err
	}

	if employees != nil {
		r.setCache(ctx, cacheKey, employees)
	}

	return employees, nil
}

func (r *EmployeeRepo) GetEmployeeByID(id uint) (*Employee, error) {
	employee := new(Employee)
	if err := r.db.First(employee, id).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeRepo) UpdateEmployeeByID(ctx context.Context, id uint, updated *Employee) (*Employee, error) {
	employee, err := r.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Model(&employee).Updates(updated).Error; err != nil {
		return employee, err
	}
	r.deleteEmployeeCache(ctx)
	return employee, nil
}

func (r *EmployeeRepo) DeleteEmployee(ctx context.Context, id uint) error {
	_, err := r.GetEmployeeByID(id)
	if err != nil {
		return err
	}
	r.deleteEmployeeCache(ctx)
	return db.DB.Delete(&Employee{}, id).Error
}

func (r *EmployeeRepo) GetTotal() (int, error) {
	total := new(int64)
	if err := db.DB.Model(&Employee{}).Count(total).Error; err != nil {
		return 0, err
	}
	return int(*total), nil
}

func (r *EmployeeRepo) setCache(c context.Context, key string, employees []Employee) {
	if data, err := json.Marshal(employees); err != nil {
		log.Println("Error marshalling employees!")
	} else {
		r.cache.Set(c, key, data, 5*time.Minute)
	}
}

func getGetCacheKey(cursor int, offset int, pageSize int) string {
	return fmt.Sprintf("employees_c:%d_o:%d_ps:%d", cursor, offset, pageSize)
}

func (r *EmployeeRepo) getFromCache(c context.Context, cursor int, offset int, pageSize int) (string, []Employee, error) {
	cacheKey := getGetCacheKey(cursor, offset, pageSize)
	cache, err := r.cache.Get(c, cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return cacheKey, nil, nil
		} else {
			return cacheKey, nil, err
		}
	}

	var employees []Employee
	err = json.Unmarshal([]byte(cache), &employees)
	if err != nil {
		return cacheKey, nil, err
	}

	log.Printf("Hit the cache for key: %s.", cacheKey)
	return cacheKey, employees, nil
}

func (r *EmployeeRepo) deleteEmployeeCache(ctx context.Context) {
	script := `
        local keys = redis.call('KEYS', ARGV[1])
        if #keys > 0 then
            return redis.call('DEL', unpack(keys))
        else
            return 0
        end
    `

	result, err := r.cache.Eval(ctx, script, []string{}, "employees:*").Result()
	if err != nil {
		fmt.Println("Error deleting keys:", err)
	} else {
		fmt.Println("Deleted", result, "keys")
	}
}
