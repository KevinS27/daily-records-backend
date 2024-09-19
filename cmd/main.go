package main

import (
	"daily-records-backend/util/config"
	"fmt"
)

func main() {
	cfg := config.ReadConfig()
	fmt.Println(cfg)
}
