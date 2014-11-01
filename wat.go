package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ENDPOINT = "https://api.wolframalpha.com/v2/query?appid=%s&input=%s&format=plaintext"

var appId = os.Getenv("WOLFRAM_API_ID")

type Result struct {
	Name xml.Name `xml:"queryresult"`
	Pods []Pod    `xml:"pod"`
}

type Pod struct {
	Title string   `xml:"title,attr"`
	Entry []string `xml:"subpod>plaintext"`
}

func main() {
	flag.Parse()
	query := strings.Join(flag.Args(), " ")
	url := fmt.Sprintf(ENDPOINT, appId, url.QueryEscape(query))
	fmt.Print(url + "\n")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		log.Panicln(err)
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	result := new(Result)
	decoder.Decode(result)

	if result == nil {
		log.Fatalf("No results parsed!\n")
	}

	displayResults(result)
}

func displayResults(result *Result) {
	for _, pod := range result.Pods {
		fmt.Printf("%s\n", pod.Title)
		for _, subpod := range pod.Entry {
			fmt.Printf("\t%s\n", strings.Replace(subpod, "\n", "\n\t\t", -1))
		}
		fmt.Printf("\n")
	}
}
