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
	accountRepo := repositories.NewAccountRepository(db.DB)
	authRepo := repositories.NewAuthRepository(db.DB)
	authenticationHandler := NewAuthenticateHandler(config.JWTSecret, accountRepo, authRepo)

	// Signup/Login
	{
		r.POST("/signup", authenticationHandler.Signup)
		r.POST("/login", authenticationHandler.Login)
		r.POST("/refresh", authenticationHandler.RefreshToken)
	}

	// Employee
	employeeRoutes := r.Group("/employee")
	{
		employeeRepo := repositories.NewEmployeeRepo(db.DB)
		employeeHandler := NewEmployeeHandler(employeeRepo)
		employeeRoutes.Use(authenticationHandler.AuthMiddleware)
		employeeRoutes.POST("/", employeeHandler.CreateEmployee)
		employeeRoutes.GET("/", employeeHandler.GetAllEmployees)
		employeeRoutes.GET("/:id", employeeHandler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}
}
