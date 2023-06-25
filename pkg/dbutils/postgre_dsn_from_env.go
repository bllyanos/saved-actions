package dbutils

import (
	"fmt"
	"os"
)

// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func LoadPostgresDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	tz := os.Getenv("DB_TZ")

	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		host,
		user,
		pass,
		name,
		port,
		tz,
	)
}
