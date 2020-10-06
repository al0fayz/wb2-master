package api 

import (
	"wb2-master/api/databases"
	"wb2-master/api/router"
	"wb2-master/api/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var server = databases.Server{}

func Run() {
	app := fiber.New()
	app.Use(cors.New())

	//koneksi to databases
	server.ConnectDB(config.Config("DB_DRIVER"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_PORT"), config.Config("DB_HOST"), config.Config("DB_NAME"))

	//migrate
	server.Migrate()
	
	//seed
	server.Seed()
	
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

	defer server.DB.Close()
}