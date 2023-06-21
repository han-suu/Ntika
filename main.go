package main

// TODO :
// ENCRYPT AND BCRYPT MASIH GA JELAS NARONYA DIMANA
// ERROR UNTUK : REGISTER DOUBLE, AKSES TANPA JWT (KE /users)
import (
	"fmt"
	"ntika/auth"
	"ntika/handler"
	"ntika/item"
	"ntika/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	DB_USER := os.Getenv("USER")
	DB_PASS := os.Getenv("PASS")
	DB := os.Getenv("DB")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("GAGAl")
	}

	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&item.Item{})
	db.AutoMigrate(&item.CartItem{})
	db.AutoMigrate(&item.Orders{})
	db.AutoMigrate(&item.OrderItem{})
	db.AutoMigrate(&item.Images{})
	db.AutoMigrate(&item.Images2{})
	// db.AutoMigrate(&tag.SongTag{})

	userRepository := auth.NewRepo(db)
	userService := auth.NewService(userRepository)
	userHandler := handler.NewHandler(userService)

	itemRepository := item.NewRepo(db)
	itemService := item.NewService(itemRepository)
	itemHandler := handler.NewHandlerItems(itemService, userService)

	r := gin.Default()
	// r.Use(CORSMiddleware())

	v1 := r.Group("/v1").Use(CORSMiddleware())
	// ex := r.Group("/ex").Use(CORSMiddleware2())
	// api := r.Group("/api").Use(CORSMiddleware())
	// v1.Use(CORSMiddleware())

	// dev only
	v1.GET("/users", middleware.RequireAuth, userHandler.GetAllUsers)
	// v1.POST("/image", itemHandler.Pap)

	// UNIVERSAL
	v1.Static("/image", "./static/")

	// AUTH
	v1.POST("/sign-up", userHandler.CreateUser)
	v1.POST("/sign-in", userHandler.SignIn)

	// ADMIN
	v1.POST("/item", itemHandler.Create)

	// USER
	v1.GET("/user", middleware.RequireAuth, userHandler.UserProfile)
	v1.POST("/cange_address", userHandler.UpdateAddress)

	// ITEM
	v1.GET("/catalog", itemHandler.Catalog)
	// v1.GET("/recommended", )

	// ORDER
	v1.GET("/cart", itemHandler.Catalog)
	// v1.DELETE("/cart/", )
	// v1.POST("/order", )
	// v1.POST("/cart", )
	// v1.GET("/user/shipping_address", )
	// v1.GET("/shipping_price", )

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
