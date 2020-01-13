package main

import (
	"fmt"
	"github.com/maateen/getipinfo/internal/ipapi"
	"os"
)

func main()  {
	if len(os.Args) > 1 {
		ipAddress := os.Args[1]
		printData(ipapi.GetJson(ipAddress))
	} else {
		ipAddress := ""
		printData(ipapi.GetJson(ipAddress))
	}
}

func printData(jsonData map[string]interface{}) {
	fmt.Println("IP Address:", jsonData["query"])
	fmt.Println("City:", jsonData["city"])
	fmt.Println("Region:", jsonData["regionName"])
	fmt.Println("Country:", jsonData["country"])
	fmt.Println("Postal Code:", jsonData["zip"])
	fmt.Println("Time Zone:", jsonData["timezone"])
	fmt.Println("ISP:", jsonData["isp"])
}