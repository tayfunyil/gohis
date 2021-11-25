package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/fatih/color"
	"github.com/likexian/whois"
	io "gohis/lib"
	"log"
	"os"
	"strings"
)

func run() {
	x, _ := io.ReadLines("data.txt")

	for _, d := range x {
		color.Green("Checking domain: ")
		fmt.Print(d + "\n")

		result, err := whois.Whois(d)
		if err == nil {
			io.WriteStringToFile("out/"+d+".txt", result)
		}
	}
}

func main() {

	parser := argparse.NewParser("gohis", "Searches whois data for domains given in a file")

	s := parser.Flag("s", "search", &argparse.Options{Required: false, Help: "Search for domains"})
	d := parser.Flag("d", "delete", &argparse.Options{Required: false, Help: "Delete all domain files"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	isEmpty, _ := io.IsEmpty("./out")

	if *s {

		if !isEmpty {
			color.Red("Folder contains files!")

			overwrite, _ := io.GetUserInput("Overwrite? (y/n) ")

			if strings.ToLower(strings.Trim(overwrite, "\n")) == "y" {
				run()
			} else {
				fmt.Println("Aborting script ...")
				os.Exit(-1)
			}
		} else {
			run()
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
