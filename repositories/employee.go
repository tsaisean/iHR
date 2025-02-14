package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	. "iHR/repositories/model"
	"log"
	"time"
)

//go:generate mockery --all --output=./mocks
type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error)
	GetAllEmployeesAfter(ctx context.Context, id int, pageSize int) ([]Employee, error)
	GetAllEmployeesFrom(ctx context.Context, offset int, pageSize int) ([]Employee, error)
	GetEmployeeByID(ctx context.Context, id uint) (*Employee, error)
	GetEmployeeByAccID(id uint) (*Employee, error)
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

	cacheKey, employees, err := r.getAllEmployeesFromCache(ctx, id, 0, pageSize)
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
		r.cacheAllEmployees(ctx, cacheKey, employees)
	}

	return employees, nil
}

// GetAllEmployeesFrom Support cache
func (r *EmployeeRepo) GetAllEmployeesFrom(ctx context.Context, offset int, pageSize int) ([]Employee, error) {
	var employees []Employee

	cacheKey, employees, err := r.getAllEmployeesFromCache(ctx, 0, offset, pageSize)
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
		r.cacheAllEmployees(ctx, cacheKey, employees)
	}

	return employees, nil
}

func (r *EmployeeRepo) GetEmployeeByID(ctx context.Context, id uint) (*Employee, error) {
	cacheKey, employee, err := r.getEmployeeFromCache(ctx, id)
	if err != nil {
		return nil, err
	} else if employee != nil {
		return employee, nil
	}

	employee = new(Employee)
	if err := r.db.First(employee, id).Error; err != nil {
		return nil, err
	}

	r.cache.Set(ctx, cacheKey, employee, time.Minute*15)

	return employee, nil
}

func (r *EmployeeRepo) GetEmployeeByAccID(id uint) (*Employee, error) {
	employee := new(Employee)
	if err := r.db.Where("acc_id = ?", id).First(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeRepo) UpdateEmployeeByID(ctx context.Context, id uint, updated *Employee) (*Employee, error) {
	employee, err := r.GetEmployeeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&employee).Updates(updated).Error; err != nil {
		return employee, err
	}
	r.deleteEmployeeCache(ctx)
	return employee, nil
}

func (r *EmployeeRepo) DeleteEmployee(ctx context.Context, id uint) error {
	_, err := r.GetEmployeeByID(ctx, id)
	if err != nil {
		return err
	}
	r.deleteEmployeeCache(ctx)
	return r.db.Delete(&Employee{}, id).Error
}

func (r *EmployeeRepo) GetTotal() (int, error) {
	total := new(int64)
	if err := r.db.Model(&Employee{}).Count(total).Error; err != nil {
		return 0, err
	}
	return int(*total), nil
}

func (r *EmployeeRepo) cacheEmployee(c context.Context, key string, employee Employee) {
	if data, err := json.Marshal(employee); err != nil {
		log.Println("Error marshalling employees!")
	} else {
		r.cache.Set(c, key, data, 5*time.Minute)
	}
}

func (r *EmployeeRepo) cacheAllEmployees(c context.Context, key string, employees []Employee) {
	if data, err := json.Marshal(employees); err != nil {
		log.Println("Error marshalling employees!")
	} else {
		r.cache.Set(c, key, data, 5*time.Minute)
	}
}

func getEmployeeCacheKey(id uint) string {
	return fmt.Sprintf("employee:%d", id)
}

func getAllEmployeesCacheKey(cursor int, offset int, pageSize int) string {
	return fmt.Sprintf("employees_c:%d_o:%d_ps:%d", cursor, offset, pageSize)
}

func (r *EmployeeRepo) getEmployeeFromCache(c context.Context, id uint) (string, *Employee, error) {
	cacheKey := getEmployeeCacheKey(id)
	cache, err := r.cache.Get(c, cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return cacheKey, nil, nil
		} else {
			return cacheKey, nil, err
		}
	}

	var employee *Employee
	err = json.Unmarshal([]byte(cache), employee)
	if err != nil {
		return cacheKey, nil, err
	}

	log.Printf("Hit the cache for key: %s.", cacheKey)
	return cacheKey, employee, nil
}

func (r *EmployeeRepo) getAllEmployeesFromCache(c context.Context, cursor int, offset int, pageSize int) (string, []Employee, error) {
	cacheKey := getAllEmployeesCacheKey(cursor, offset, pageSize)
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

const LuaScript = `
        local keys = redis.call('KEYS', ARGV[1])
        if #keys > 0 then
            return redis.call('DEL', unpack(keys))
        else
            return 0
        end
`

func (r *EmployeeRepo) deleteEmployeeCache(ctx context.Context) {
	result, err := r.cache.Eval(ctx, LuaScript, []string{}, "employees:*").Result()
	if err != nil {
		fmt.Println("Error deleting keys:", err)
	} else {
		fmt.Println("Deleted", result, "keys")
	}
}
