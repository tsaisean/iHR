package route

import (
	"github.com/gin-gonic/gin"
	"iHR/config"
	"iHR/db"
	"iHR/db/repositories"
	. "iHR/handler/authenticate"
	. "iHR/handler/employee"
)

func RegisterRoutes(r *gin.Engine, config *config.Config) {
	// Signup/Login
	{
		accountRepo := repositories.NewAccountRepository(db.DB)
		authRepo := repositories.NewAuthRepository(db.DB)
		accountHandler := NewAuthenticateHandler(config.JWTSecret, accountRepo, authRepo)
		r.POST("/signup", accountHandler.Signup)
		r.POST("/login", accountHandler.Login)
		r.POST("/refresh", accountHandler.RefreshToken)
	}

	// Employee
	employeeRoutes := r.Group("/employee")
	{
		employeeRepo := repositories.NewEmployeeRepo(db.DB)
		employeeHandler := NewEmployeeHandler(employeeRepo)
		employeeRoutes.POST("/", employeeHandler.CreateEmployee)
		employeeRoutes.GET("/", employeeHandler.GetAllEmployees)
		employeeRoutes.GET("/:id", employeeHandler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}
}
