package main

import (
	"flag"
	"fmt"
	"strconv"
)

type CVTOption struct {
	HelpFlag    bool
	D2B string
}

func parse() (*CVTOption, []string) {
	options := &CVTOption{}
	flag.BoolVar(&options.HelpFlag, "help", false, "Displays usage information and exit.")
	flag.StringVar(&options.D2B, "d2b", "", "Decimal number to Binary number.")

	flag.Parse()
	return options, flag.Args()
}

func main() {
	options, _ := parse()
	switch {
	case options.D2B != "":
		i, err := strconv.ParseInt(options.D2B, 10, 64)
		if err != nil {
			fmt.Printf("parse number '%s' error for %s\n", options.D2B, err.Error())
		} else {
			r := strconv.FormatInt(i, 2)
			fmt.Println(r)
		}
	case options.HelpFlag:
		flag.PrintDefaults()
	default:
		flag.PrintDefaults()
	}
}
