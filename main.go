package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

var wg sync.WaitGroup
var multiThreads bool

func main() {
	var target, siteType string
	var verbose bool

	flag.StringVar(&target, "target", "", "target domain name")
	flag.StringVar(&siteType, "type", "", "site type : wp = wordpress")
	flag.BoolVar(&verbose, "v", false, "verbose mode. default : false")
	flag.BoolVar(&multiThreads, "mt", false, "multi threads mode. default : false")
	flag.Parse()

	if target == "" || siteType == "" {
		fmt.Println("You didn't provide enough arguments. Refer to README.md to have the usage detail.")
		return
	}
	startTime := time.Now()

	list := getList(siteType)
	fmt.Printf("\n"+time.Now().Format("15:04:05")+" -- Scan started on %s with %s list\n\n", target, siteType)

	if multiThreads {
		firstPart := list[:(len(list) / 2)]
		secondPart := list[(len(list) / 2):]

		wg.Add(1)
		go checkURL(firstPart, target, verbose)
		wg.Add(1)
		go checkURL(secondPart, target, verbose)

		wg.Wait()

	} else {
		checkURL(list, target, verbose)
	}

	elapsedTime := time.Now().Sub(startTime)
	fmt.Printf("\nScan terminated in %v\n", elapsedTime)
}

func contact(target string) (int, error) {
	resp, err := http.Get(target)
	if err != nil {
		// if webpage doesn't exits, I simplify it as a simple 404 code
		return 0, err
	}

	return resp.StatusCode, nil
}

func displayResult(statusCode int, target, url string, v bool) {
	if statusCode >= 400 && statusCode <= 499 && v == true {
		color.Red("%v : %s is not present\n", statusCode, target+url)
	} else if statusCode >= 200 && statusCode <= 299 {
		color.Green("%v : %s is present\n", statusCode, target+url)
	} else if statusCode >= 500 && statusCode <= 599 {
		color.Magenta("%v : %s respond internal server error\n", statusCode, target+url)
	}
}

func getList(siteType string) []string {
	file, err := os.Open("lists/" + siteType + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var listlines []string

	for scanner.Scan() {
		listlines = append(listlines, scanner.Text())
	}

	file.Close()

	return listlines
}

func checkURL(givenList []string, target string, verbose bool) {
	if multiThreads {
		defer wg.Done()
	}

	for _, url := range givenList {
		if url[0] != '/' {
			url = "/" + url
		}
		statusCode, err := contact(target + url)
		if err != nil {
			log.Fatalf("and error occured : %v\n", err)
		}

		displayResult(statusCode, target, url, verbose)
	}
}
