package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	rod "github.com/go-rod/rod"
	launcher "github.com/go-rod/rod/lib/launcher"
	robotgo "github.com/go-vgo/robotgo"
)

var lensUrl = "https://lens.google.com"

func main() {
	url := strings.TrimSpace(strings.Join(os.Args[1:], ""))
	if url == "" {
		log.Fatalf("missing url parameter")
	}

	fullUrl := fmt.Sprintf("%s/uploadbyurl?url=%s", lensUrl, url)

	scrape(fullUrl)
}

func scrape(url string) {
	u := launcher.New().
		Headless(true).
		Set("single-process").
		MustLaunch()

	page := rod.New().
		ControlURL(u).
		MustConnect().
		MustPage(url)

	page.MustElement("#ucj-2").MustClick()

	page.
		MustWaitStable().
		MustElement(".VfPpkd-LgbsSe.VfPpkd-LgbsSe-OWXEXe-k8QpJ.VfPpkd-LgbsSe-OWXEXe-dgl2Hf.nCP5yc.AjY5Oe.DuMIQc.LQeN7.kCfKMb").
		MustClick()

	text := page.MustWaitStable().MustElement("[jsname=\"r4nke\"]").MustText()

	robotgo.WriteAll(text)
}
