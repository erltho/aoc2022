package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	"net/http/cookiejar"
	io "io/ioutil"
	"bufio"
	"strings"
)
var client http.Client
func init() {
    jar, err := cookiejar.New(nil)
    if err != nil {
        log.Fatalf("Got error while creating cookie jar %s", err.Error())
    }
    client = http.Client{
        Jar: jar,
    }
}
func main() {
    cookie := &http.Cookie{
        Name:   "session",
        Value:  os.Getenv("API_TOKEN"),
        MaxAge: 300,
    }
    req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/1/input", nil)
    if err != nil {
        log.Fatalf("Got error %s", err.Error())
    }
    req.AddCookie(cookie)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error occured. Error is: %s", err.Error())
    }
    defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Print(countElves(bodyString))
	}
	
}

func countElves(data string) int {
	scanner := bufio.NewScanner(strings.NewReader(data))
	elves := 0
	for scanner.Scan() {
    	if scanner.Text() == "" {
			elves += 1
		}
	}
	return elves
}