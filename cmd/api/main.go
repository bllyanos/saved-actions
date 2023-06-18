package main

import (
	"log"
	"time"

	"github.com/bllyanos/saved-actions/internals/models"
	"github.com/bllyanos/saved-actions/internals/routes"
	"github.com/bllyanos/saved-actions/pkg/dbutils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	app := gin.Default()

	dbDsn := dbutils.LoadPostgresDSN()

	var db *gorm.DB
	if dbOpen, err := gorm.Open(
		postgres.Open(dbDsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
		},
	); err != nil {
		log.Fatalf("cannot connect to database: %v\n", err.Error())
	} else {
		db = dbOpen
	}

	migrationStart := time.Now()
	if err := db.AutoMigrate(
		&models.Action{},
		&models.HttpActionDetail{},
		&models.ActionParameter{},
		&models.ActionRun{},
		&models.ActionRunParameter{},
		&models.ActionRunEvent{},
	); err != nil {
		log.Fatalf("migrations failed: %v\n", err.Error())
	} else {
		deltaTime := time.Now().Sub(migrationStart)
		log.Printf("migrations finished in %s\n", deltaTime)
	}

	routes.SetupIndexRoute(app, db)
	routes.SetupActionsRoute(app, db)

	if err := app.Run(":8080"); err != nil {
		log.Fatal("run error", err.Error())
	}
}
