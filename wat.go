package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ENDPOINT = "https://api.wolframalpha.com/v2/query?appid=%s&input=%s&format=plaintext"
const USAGE = `wat - Wolfram Alpha command-line Tool

wat is a command-line utility to more easily query Wolfram Alpha's API

A Wolfram Alpha API key is required, and can be obtained from https://developer.wolframalpha.com
The Wolfram Alpha API key is loaded from the 'WOLFRAM_API_ID' environment variable.

Usage:
    WOLFRAM_API_ID="<appid>" wat <query>
  or
    export WOLFRAM_API_ID="<appid>"
    wat <query>
`

var appId = os.Getenv("WOLFRAM_API_ID")

type Result struct {
	Name    xml.Name `xml:"queryresult"`
	Success bool     `xml:"success,attr"`
	Pods    []Pod    `xml:"pod"`
	Error   Err      `xml:"error"`
}

type Pod struct {
	Title string   `xml:"title,attr"`
	Entry []string `xml:"subpod>plaintext"`
}

type Err struct {
	Code    int32  `xml:"code"`
	Message string `xml:"msg"`
}

func main() {
	flag.Parse()
	query := strings.Join(flag.Args(), " ")

	if 0 >= len(appId) {
		fmt.Printf("Must set WOLFRAM_API_ID env variable with API Key.\n\n")
		fmt.Println(USAGE)
		os.Exit(1)
	}

	if 0 >= len(query) {
		fmt.Printf("query string must not be blank\n\n")
		fmt.Println(USAGE)
		os.Exit(2)
	}

	result, err := executeQuery(query)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(3)
	}

	if result.Success {
		displaySuccess(result)
	} else {
		displayError(result)
	}
}

func executeQuery(query string) (result *Result, error error) {
	url := fmt.Sprintf(ENDPOINT, appId, url.QueryEscape(query))

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unexpected error fetching from Wolfram Alpha: %s", err))
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	result = new(Result)
	decoder.Decode(result)

	if resp.StatusCode != 200 || result == nil {
		return nil, errors.New(fmt.Sprintf("Unexpected response from Wolfram Alpha. Status Code:%d", resp.StatusCode))
	}

	return result, nil
}

func displaySuccess(result *Result) {
	for _, pod := range result.Pods {
		fmt.Printf("%s\n", pod.Title)
		for _, subpod := range pod.Entry {
			fmt.Printf("\t%s\n", strings.Replace(subpod, "\n", "\n\t\t", -1))
		}
		fmt.Printf("\n")
	}
}

func displayError(result *Result) {
	fmt.Printf("Error response from Wolfram Alpha: %s - Code %d\n", result.Error.Message, result.Error.Code)
}
