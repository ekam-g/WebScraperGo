package main

import (
	"WebScraperGo/func/web_scrap"
	"log"
)

func main() {
	Demo := web_scrap.Scrape{}.Demo()
	log.Print(Demo)
}
