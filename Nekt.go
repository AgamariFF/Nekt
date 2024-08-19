package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

func main() {
	optionsAi := append(
		chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.ProxyServer("45.8.211.64:80"),
		chromedp.Flag("headless", false),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"),
		chromedp.Flag("enable-automation", false),
		// chromedp.Flag("disable-web-security", false),
		chromedp.Flag("disable-web-security", false),
		chromedp.Flag("allow-running-insecure-content", true),
	)
	allocCtxAi, cancel := chromedp.NewExecAllocator(context.Background(), optionsAi...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtxAi)
	defer cancel()
	url := "https://nekto.me/audiochat#/"
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		panic(err)
	}
	for {
		err = chromedp.Run(ctx,
			chromedp.WaitVisible(".go-scan-button"),
			chromedp.Click(".go-scan-button"),
		)
		if err != nil {
			panic(err)
		}
	}
}
