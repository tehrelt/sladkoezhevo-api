package main

import (
	"fmt"
	"sladkoezhevo-api/internal/config"
)

func main() {
	config := config.NewConfig("config/.yaml")

	fmt.Print(config)
}
