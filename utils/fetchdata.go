package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func FetchQuestionData(day int) (*string, error) {
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	Assert(day > 0 && day < 32, "Incorrect day for Advent of code provided")

	client := &http.Client{
		Transport: &http.Transport{},
	}

	year, err := strconv.Atoi(k.String("year"))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	currntYear, _, _ := time.Now().Date()
	Assert(year <= int(currntYear) && year >= 2015, "Not a valid Advent of Code year")

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", strconv.Itoa(year), strconv.Itoa(day))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error
		fmt.Print(err.Error())
		os.Exit(1)
	}

	session := k.String("session")
	Assert(session != "", "Please provide a session")

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
		return nil, err
	}

	srtringData := string(responseData)
	return &srtringData, nil
}
