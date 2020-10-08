package api 

import (
	"wb2-master/api/databases"
	"wb2-master/api/routes"
	"wb2-master/api/utils"
	"wb2-master/api/config"
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(cors.New())
	app.Use(logger.New())

	//koneksi to databases
	databases.ConnectDB(config.Config("DB_DRIVER"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_PORT"), config.Config("DB_HOST"), config.Config("DB_NAME"))

	//migrate
	databases.Migrate()
	
	//seed
	databases.Seed()
	
	//global routes
	routes.SetupRoutes(app)
	//admin routes
	routes.AdminRoutes(app)
	//mitra routes
	routes.MitraRoutes(app)
	//agent routes 
	routes.AgentRoutes(app)
	//sub agent routes
	routes.SubAgentRoutes(app)
	
	//server serve
	app.Listen(fmt.Sprintf(":%v", config.Config("PORT")))

	defer databases.DB.Close()
}