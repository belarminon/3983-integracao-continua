package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (	
	DB	*gorm.DB
	err	error
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" { return fallback }
	return value
}

func ConectaComBancoDeDados() {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "root")
	dbname := getEnv("DB_NAME", "root")
	port := getEnv("DB_PORT", "5432")
	conectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conectionString))
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}
