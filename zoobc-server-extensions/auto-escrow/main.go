package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zoobc-server-extensions/auto-escrow/cmd"
	conn "github.com/zoobc-server-extensions/auto-escrow/connection"

	/** . "github.com/zoobc-server-extensions/auto-escrow/connection" several ways to declare an import */
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("SERVER AUTO ESCROW RUNNING ON PORT :8080")
	c := conn.Connect()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	cmd.Execute()
}
