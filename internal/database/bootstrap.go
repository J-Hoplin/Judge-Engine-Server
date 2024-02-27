package database

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"judge-engine/ent"
	"log"
	"os"
)

var client *ent.Client

// Bootstrap Function of ent
func MySQLBootstrap() error {
	var err error
	client, err = ent.Open("mysql", os.Getenv("ENT_DB_URL"))
	if err != nil {
		return err
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}

// Get client
func GetClient() *ent.Client {
	return client
}

func CloseClient() {
	log.Println("Closing database connection...")
	client.Close()
	log.Println("Done!")
}
