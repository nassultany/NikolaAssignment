package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Person struct {
	Id        int    `json:"id"`
	Address   string `json:"address"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
}

func main() {
	inputFile := flag.String("input", "", "input json file to sort")
	outputFile := flag.String("output", "out.json", "file to output results to")
	flag.Parse()

	// Open and read json file into slice
	jsonFile, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer jsonFile.Close()

	var people []Person

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	json.Unmarshal(bytes, &people)

	// fmt.Println("Before:")
	// printInfo(people)

	// sort
	sort.Slice(people, func(i, j int) bool {
		var li string = strings.ToLower(people[i].LastName)
		var lj string = strings.ToLower(people[j].LastName)
		if li == lj {
			return people[i].Address < people[j].Address
		}
		return people[i].LastName < people[j].LastName
	})

	// fmt.Println("After:")
	// printInfo(people)

	// write to output file
	jsonString, err := json.MarshalIndent(people, "", " ")
	ioutil.WriteFile(*outputFile, jsonString, 0644)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func printInfo(people []Person) {
	for _, person := range people {
		fmt.Println("ID:", person.Id)
		fmt.Printf("Name: %s, %s\n", person.LastName, person.FirstName)
		fmt.Println("Address:", person.Address)
		fmt.Println()
	}
}
