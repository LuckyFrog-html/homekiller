package main

import (
	"fmt"
	"server/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Russia";
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});
	//
	//_ = err
	//db.
}
