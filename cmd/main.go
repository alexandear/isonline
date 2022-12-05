package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/alexandear/isonline/pkg/datagroup"
)

func main() {
	dataGroupLogin := flag.String("datagroup_login", "", "Login to personal cabinet https://my.datagroup.ua")
	dataGroupPassword := flag.String("datagroup_password", "", "Password to personal cabinet")
	if dataGroupLogin == nil || dataGroupPassword == nil {
		*dataGroupLogin = os.Getenv("DATAGROUP_LOGIN")
		*dataGroupPassword = os.Getenv("DATAGROUP_PASSWORD")
	}

	provider := datagroup.New(datagroup.Config{
		Login:    *dataGroupLogin,
		Password: *dataGroupPassword,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	isOnline, err := provider.IsOnline(ctx)
	if err != nil {
		log.Fatalf("Failed to retrive status: %v.", err)
	}

	log.Printf("Is online DataGroup: %t", isOnline)
}
