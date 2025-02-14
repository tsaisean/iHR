package route

import (
	"github.com/gin-gonic/gin"
	"iHR/config"
	. "iHR/handler/authenticate"
	"iHR/handler/authenticate/oauth/google"
	. "iHR/handler/employee"
	"iHR/handler/leave"
	"iHR/repositories"
	"iHR/repositories/db"
	"iHR/repositories/redis"
)

func RegisterRoutes(r *gin.Engine, config *config.Config) {
	accountRepo := repositories.NewAccountRepository(db.DB)
	authRepo := repositories.NewAuthRepository(db.DB)
	employeeRepo := repositories.NewEmployeeRepo(db.DB, redis.RedisClient)
	authenticationHandler := NewAuthenticateHandler(config.JWTSecret, accountRepo, authRepo, employeeRepo)
	googleOAuthHandler := google.NewGoogleOAuthHandler(config.JWTSecret, config.Oauth.Google, authRepo, accountRepo, employeeRepo)

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
		employeeHandler := NewEmployeeHandler(employeeRepo)
		employeeRoutes.Use(authenticationHandler.AuthMiddleware)
		employeeRoutes.POST("/", employeeHandler.CreateEmployee)
		employeeRoutes.GET("/", employeeHandler.GetAllEmployees)
		employeeRoutes.GET("/:id", employeeHandler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}

	// Leave
	leavesRoutes := r.Group("/leave")
	{
		leaveRepo := repositories.NewLeaveRepo(db.DB)
		leaveHandler := leave.NewLeaveHandler(leaveRepo)
		leavesRoutes.Use(authenticationHandler.AuthMiddleware)
		leavesRoutes.POST("/requests", leaveHandler.CreateLeaveRequest)
		leavesRoutes.PUT("/requests/:id", leaveHandler.UpdateLeaveRequest)
		leavesRoutes.GET("/requests", leaveHandler.GetLeaveRequests)
		leavesRoutes.POST("/balances/", leaveHandler.CreateLeaveBalance)
		leavesRoutes.PUT("/balances/:id", leaveHandler.CreateLeaveBalance)
		leavesRoutes.GET("/summaries", leaveHandler.GetLeaveSummary)
	}
}
