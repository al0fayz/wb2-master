package databases

import (
	"fmt"
	"log"
	"wb2-master/api/databases/migration"
	"wb2-master/api/databases/seed"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database \n", Dbdriver)
	}
	
}
func Seed() {
	//seed
	seed.Load(DB)
}
func Migrate() {
	migration.Migrate(DB)
}
func DropTables() {
	migration.DropTable(DB)
}