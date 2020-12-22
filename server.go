package main

import (
	"golang_api/config"
	"golang_api/controller"
	"golang_api/repository"
	"golang_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)

	jwtService  service.JWTService  = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(userRepository)

	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	//pasang middleware
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
