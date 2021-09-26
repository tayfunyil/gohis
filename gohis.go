package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/likexian/whois"
	io "gohis/lib"
	"log"
	"os"
)


func main() {

	parser := argparse.NewParser("gohis", "Searches whois data for domains given in a file")

	s := parser.Flag("s", "search", &argparse.Options{Required: false, Help: "Search for domains"})
	d := parser.Flag("d", "delete", &argparse.Options{Required: false, Help: "Delete all domain files"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *s {
		x, _ := io.ReadLines("data.txt")

		for _, d := range x {
			fmt.Println("Checking domain: " + d)

			result, err := whois.Whois(d)
			if err == nil {
				io.WriteStringToFile("out/" + d + ".txt", result)
			}
		}
	} else if *d {
			x, _ := io.ListAllFilesInDir("./out")

			for _, d := range x {
				fmt.Println("Deleting file: " + d)

				e := os.Remove("./out/" + d)
				if e != nil {
					log.Fatal(e)
				}

			}
	}







}
