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
	var (
		dataGroupLogin    string
		dataGroupPassword string
	)
	flag.StringVar(&dataGroupLogin, "datagroup_login", "", "Login to personal cabinet https://my.datagroup.ua")
	flag.StringVar(&dataGroupPassword, "datagroup_password", "", "Password to personal cabinet")
	if dataGroupLogin == "" || dataGroupPassword == "" {
		dataGroupLogin = os.Getenv("DATAGROUP_LOGIN")
		dataGroupPassword = os.Getenv("DATAGROUP_PASSWORD")
	}

	provider := datagroup.New(datagroup.Config{
		Login:    dataGroupLogin,
		Password: dataGroupPassword,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	isOnline, err := provider.IsOnline(ctx)
	if err != nil {
		log.Fatalf("Failed to retrive status: %v.", err)
	}

	if isOnline {
		log.Printf("DataGroup is online\n")
	} else {
		log.Printf("DataGroup is offline\n")
	}
}
