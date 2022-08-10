package main

import (
	"WebScraperGo/func/web_scrap"
	"log"
	"time"
)

func main() {
	x := 1
	go func() {
		for true {
			DemoData, err := web_scrap.Scrape{}.Demo()
			if err != nil {
				log.Fatal(err)
			}
			for _, slice := range DemoData {
				log.Print(slice)

			}
			x++
			print(x)
		}
	}()
	time.Sleep(time.Second * 30)
}
