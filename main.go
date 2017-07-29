/*
	ping1 test program
	Author: Kjell Almgren, Tetracon
	Date: 2017-07-29
	Version: 0.1
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {

	color.Set(color.FgHiGreen)
	fmt.Printf("ping1 test is started on server: ")
	color.Set(color.FgHiWhite)
	fmt.Printf("%s\r\n", GetHostname())
	color.Unset()
	resp, err := http.Get("http://www.svt.se")

	if err != nil {
		fmt.Printf("Error: %s\r\n", err.Error())
	} else {
		fmt.Printf("Output (www.svt.se): %s\r\n", resp.Header)
	}
}

//
//	Get hostname of running server
//
func GetHostname() string {

	hostname, err1 := os.Hostname()
	if err1 == nil {
		//log.Printf("Hostname: %s", hostname)
		//fmt.Println("Hostname: ", hostname)
	}
	return hostname
}
