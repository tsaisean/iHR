package route

import (
	"github.com/gin-gonic/gin"
	"iHR/config"
	. "iHR/handler/authenticate"
	"iHR/handler/authenticate/oauth/google"
	. "iHR/handler/employee"
	"iHR/repositories"
	"iHR/repositories/db"
	"iHR/repositories/redis"
)

func RegisterRoutes(r *gin.Engine, config *config.Config) {
	accountRepo := repositories.NewAccountRepository(db.DB)
	authRepo := repositories.NewAuthRepository(db.DB)
	authenticationHandler := NewAuthenticateHandler(config.JWTSecret, accountRepo, authRepo)
	googleOAuthHandler := google.NewGoogleOAuthHandler(config.JWTSecret, config.Oauth.Google, authRepo, accountRepo)

	// Signup/Login
	{
		r.POST("/signup", authenticationHandler.Signup)
		r.POST("/login", authenticationHandler.Login)
		r.POST("/refresh", authenticationHandler.RefreshToken)
	}

	// Oauth
	{
		r.GET("/auth/google/signup", googleOAuthHandler.Signup)
		r.GET("/auth/google/login", googleOAuthHandler.Login)
		r.GET("/auth/google/callback", googleOAuthHandler.Callback)
	}

	// Employee
	employeeRoutes := r.Group("/employees")
	{
		employeeRepo := repositories.NewEmployeeRepo(db.DB, redis.RedisClient)
		employeeHandler := NewEmployeeHandler(employeeRepo)
		employeeRoutes.Use(authenticationHandler.AuthMiddleware)
		employeeRoutes.POST("/", employeeHandler.CreateEmployee)
		employeeRoutes.GET("/", employeeHandler.GetAllEmployees)
		employeeRoutes.GET("/:id", employeeHandler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}
}
