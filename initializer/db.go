package initializer

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_redis "github.com/go-redis/redis/v8"
	

	"github.com/hoon3051/TilltheCop/model"
)

var DB *gorm.DB
var Redis *_redis.Client

func InitDB() {
	ConnectDB()
	SyncDB()
	InitRedis(1)
}


func InitRedis(selectDB ...int) {
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	Redis = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
	})
}

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

}

func SyncDB() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Oauth{})
	DB.AutoMigrate(&model.Profile{})
	DB.AutoMigrate(&model.Report{})
	DB.AutoMigrate(&model.Record{})
}