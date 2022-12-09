package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/alexandear/isonline/pkg/datagroup"
	"github.com/alexandear/isonline/pkg/xflag"
)

func main() {
	var dgLoginFlag xflag.StringFlag
	flag.Var(&dgLoginFlag, "datagroup_login", "Login to personal cabinet https://my.datagroup.ua")
	var dgPasswordFlag xflag.StringFlag
	flag.Var(&dgPasswordFlag, "datagroup_password", "Password to personal cabinet")
	flag.Parse()

	dgLogin := dgLoginFlag.Value()
	dgPassword := dgPasswordFlag.Value()
	if !dgLoginFlag.IsSet() {
		dgLogin = os.Getenv("DATAGROUP_LOGIN")
		dgPassword = os.Getenv("DATAGROUP_PASSWORD")
	}

	log.Printf("Using provider: DataGroup, login: %s", dgLogin)

	provider, err := datagroup.New(datagroup.Config{
		Login:    dgLogin,
		Password: dgPassword,
	})
	if err != nil {
		log.Fatalf("Failed to create DataGroup provider: %v\n", err)
	}

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
