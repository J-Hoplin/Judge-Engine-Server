package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"judge-engine/cmd"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(os.Getenv("DB_URLs"))
	cmd.Execute()
}
