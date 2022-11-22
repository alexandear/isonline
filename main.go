package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	login := os.Getenv("DATAGROUP_LOGIN")
	password := os.Getenv("DATAGROUP_PASSWORD")

	var ip string
	if err := chromedp.Run(ctx, userIP(login, password, &ip)); err != nil {
		log.Fatal(err)
	}

	log.Printf("Is online: %t, ip: %s", ip != "", ip)
}

func userIP(login, password string, ip *string) chromedp.Tasks {
	loginSel := `//input[@name="login"]`
	passwordSel := `//input[@name="password"]`
	return chromedp.Tasks{
		chromedp.Navigate("https://my.datagroup.ua/"),
		chromedp.WaitVisible(loginSel),
		chromedp.WaitVisible(passwordSel),
		chromedp.SendKeys(loginSel, login),
		chromedp.SendKeys(passwordSel, password),
		chromedp.Submit(loginSel),
		chromedp.WaitVisible(`#user_ip`),
		chromedp.Text(`#user_ip`, ip),
	}
}
