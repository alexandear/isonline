package datagroup

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

type Config struct {
	Login    string
	Password string
}

func (c Config) IsValid() error {
	if c.Login == "" {
		return fmt.Errorf("login must be non-empty")
	}
	return nil
}

type DataGroup struct {
	config Config
}

func New(config Config) (*DataGroup, error) {
	if err := config.IsValid(); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}
	return &DataGroup{config: config}, nil
}

func (dg *DataGroup) IsOnline(ctx context.Context) (bool, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var ip string
	if err := chromedp.Run(ctx, userIP(dg.config.Login, dg.config.Password, &ip)); err != nil {
		return false, fmt.Errorf("run: %w", err)
	}

	return ip != "", nil
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
