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
	// "strconv"
	// "sort"
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
    req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/2/input", nil)
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
		totalScore(bodyString)
	}
}

func totalScore(data string) int {
	totalScore := 0
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Fields(line)
		totalScore += calculatePoints(round)
	}
	fmt.Print(totalScore)
	return totalScore
}

// A Rock
// B Paper
// C Scissors

// X Rock
// Y paper
//Z scissors

func calculatePoints(round []string) int {
	m := make(map[string]int)
	m["X"] = 1
    m["Y"] = 2
	m["Z"] = 3
	points := 0
	if round[0] == "A" {
		if round[1] == "Y" {
			points += 6
		}
		if round[1] == "X" {
			points += 3
		}
	} 
	if round[0] == "B" {
		if round[1] == "Z" {
			points += 6
		}
		if round[1] == "Y" {
			points += 3
		}
	} 
	if round[0] == "C" {
		if round[1] == "X" {
			points += 6
		}
		if round[1] == "Z" {
			points += 3
		}
	}
	points += m[round[1]]
	return points
}