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
	"github.com/sirupsen/logrus"
)

var wg sync.WaitGroup
var threads int
var client http.Client

func main() {
	var target, dict string
	var verbose bool
	var wait, timeout int

	flag.StringVar(&target, "t", "", "target domain name")
	flag.StringVar(&dict, "d", "", "dictionnary path")
	flag.IntVar(&wait, "w", 0, "waiting time between requests")
	flag.IntVar(&threads, "T", 1, "number of threads. default : 1")
	flag.IntVar(&timeout, "to", 4, "client timeout")
	flag.BoolVar(&verbose, "v", false, "verbose mode. default : false")
	flag.Parse()

	if target == "" || dict == "" {
		fmt.Println("You didn't provide enough arguments. Refer to README.md to have the usage detail.")
		return
	}

	client.Timeout = time.Duration(timeout) * time.Second

	startTime := time.Now()

	list := getList(dict)
	fmt.Printf("\nTARGET : %s\n", target)
	fmt.Printf("DICT : %s\n", dict)
	fmt.Println("START TIME : " + time.Now().Format("15:04:05"))
	fmt.Printf("THREADS : %d\n", threads)
	fmt.Printf("-- Threads init\n\n")

	// Create goroutines
	for i := 0; i < threads; i++ {
		min := (len(list) / threads) * i
		max := (len(list) / threads) * (i + 1)
		wg.Add(1)
		go checkURL(list[min:max], target, verbose, wait)

		logrus.Infof("thread %d created", i)
	}

	fmt.Printf("\n-- Scan started\n\n")

	// Wait for all the goroutines to end
	wg.Wait()

	elapsedTime := time.Now().Sub(startTime)
	fmt.Printf("\n-- Scan terminated in %v\n", elapsedTime)
}

// contact sends a request to a specified target
func contact(target string) (int, error) {
	resp, err := client.Get(target)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

// displayResults only processes the statusCode to display the result in a specific color
func displayResult(statusCode int, target, url string, v bool) {
	if statusCode >= 400 && statusCode <= 499 && v == true {
		color.Red("%v : %s is not present\n", statusCode, target+url)
	} else if statusCode >= 200 && statusCode <= 299 {
		color.Green("%v : %s\n", statusCode, target+url)
	} else if statusCode >= 500 && statusCode <= 599 && v == true {
		color.Magenta("%v : %s respond internal server error\n", statusCode, target+url)
	}
}

// getList creates and returns a string array based on the filename given in parameter
func getList(dict string) []string {
	file, err := os.Open(dict)

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

// checkURL is the core function : it calls contact to perform a GET request to the provided target with a specific path,
// then calls displayResults.
func checkURL(givenList []string, target string, verbose bool, wait int) {

	defer wg.Done()

	for _, url := range givenList {
		if url[0] != '/' {
			url = "/" + url
		}
		statusCode, err := contact(target + url)
		if err != nil && verbose == true {
			logrus.Warnf("an error occured : %v\n", err)
		}

		displayResult(statusCode, target, url, verbose)
		if wait != 0 {
			time.Sleep(time.Duration(wait) * time.Millisecond)
		}
	}
}
