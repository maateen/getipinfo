package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"github.com/maateen/getipinfo/internal/ipapi"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		ipAddress := os.Args[1]
		printData(ipapi.GetJson(ipAddress))
	} else {
		ipAddress := ""
		printData(ipapi.GetJson(ipAddress))
	}
}

func printData(jsonData map[string]interface{}) {
	fmt.Println(Bold("\nIP Address:"), Green(jsonData["query"]))
	fmt.Println(Bold("City:"), Green(jsonData["city"]))
	fmt.Println(Bold("Region:"), Green(jsonData["regionName"]))
	fmt.Println(Bold("Country:"), Green(jsonData["country"]))
	fmt.Println(Bold("Postal Code:"), Green(jsonData["zip"]))
	fmt.Println(Bold("Time Zone:"), Green(jsonData["timezone"]))
	fmt.Println(Bold("ISP:"), Green(jsonData["isp"]))
}
