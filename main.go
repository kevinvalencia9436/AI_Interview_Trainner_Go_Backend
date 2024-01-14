package main

import (
	"fmt"
	"os"
	"time"
	user_model "up-it-aps-api/app/models/user"
	service "up-it-aps-api/app/services"
	_ "up-it-aps-api/docs"
	"up-it-aps-api/pkg/middleware"
	"up-it-aps-api/pkg/routes"
	"up-it-aps-api/platform/database"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Start of the sample swagger config
// @title			Fiber Example API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {
	app := fiber.New(fiber.Config{DisablePreParseMultipartForm: true, StreamRequestBody: true, PassLocalsToViews: true})
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	config := session.Config{
		Expiration:        24 * time.Hour,
		KeyLookup:         "cookie:session",
		CookieSessionOnly: true,
		CookieSecure:      true,
	}
	store := session.New(config)
	dotEnvError := godotenv.Load()
	if dotEnvError != nil {
		println(dotEnvError)
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     os.Getenv("ALLOWED_ORIGINS"),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-API-KEY",
	}))

	initDatabase()
	assignRoutes(app, store)

	appListeningError := app.Listen("0.0.0.0:8080")
	if appListeningError != nil {
		return
	}
}

func assignRoutes(app *fiber.App, store *session.Store) {
	app.Get("/", healthCheck)

	api := app.Group("/api")
	// This or another middleware function needs to be checking the validity of the jwt
	// Very sus atm
	api.Use("/", middleware.WithAuthenticatedUserApi)

	auth := api.Group("/auth")
	auth.Get("/callback", handleLoginCallback)
	auth.Get("/logout", func(c *fiber.Ctx) error {
		return logout(c, store)
	})

	routes.AiRoutes(api, store)
	routes.UserRoutes(api, store)
	routes.DebuggingRoutes(api, store)
}

func initDatabase() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	database.DBConn = db
	if err != nil {
		panic("failed to connect database")
	}
	migrationError := db.AutoMigrate(&user_model.User{}, &user_model.UserSettings{})
	if migrationError != nil {
		return
	}
	fmt.Println("Connection Opened to Database")
}

// HealthCheck godoc
//
//	@Summary		Show the status of server, pog.
//	@Description	get the status of server.
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ [get]
func healthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

func logout(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	err := sess.Destroy()
	if err != nil {
		return err
	}
	return c.Next()
}

var jwtSecret = []byte("secret")

func handleLoginCallback(c *fiber.Ctx) error {
	fmt.Println("handleLoginCallback")
	userService := service.NewUserService()
	tokenString := c.Get("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	user := fiber.Map{"email": token.Claims.(jwt.MapClaims)["email"], "name": token.Claims.(jwt.MapClaims)["name"]}
	email := token.Claims.(jwt.MapClaims)["email"].(string)
	retrievedUser := userService.GetUserByEmail(email)

	if retrievedUser.Email == "" {
		newUser := user_model.InputUser{
			Email: email,
		}
		userService.CreateUser(&newUser)
	}
	fmt.Println("Set user to session")
	return c.JSON(user)
}
