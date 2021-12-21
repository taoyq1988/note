package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CVTOption struct {
	HelpFlag bool
	D2B      string
	D2H      string
	H2B      string
	H2D      string
	H2S      string
}

func parse() (*CVTOption, []string) {
	options := &CVTOption{}
	flag.BoolVar(&options.HelpFlag, "help", false, "Displays usage information and exit.")
	flag.StringVar(&options.D2B, "d2b", "", "Decimal number to Binary number.")
	flag.StringVar(&options.D2H, "d2h", "", "Decimal number to Hex number.")
	flag.StringVar(&options.H2B, "h2b", "", "Hex number to Binary number.")
	flag.StringVar(&options.H2D, "h2d", "", "Hex number to Decimal number.")
	flag.StringVar(&options.H2S, "h2s", "", "decode Hex string to utf8 string.")

	flag.Parse()
	return options, flag.Args()
}

func main() {
	options, _ := parse()
	switch {
	case options.D2B != "":
		convertHexNum(options.D2B, 10, 2)
	case options.D2H != "":
		convertHexNum(options.D2H, 10, 16)
	case options.H2B != "":
		convertHexNum(options.H2B, 16, 2)
	case options.H2D != "":
		convertHexNum(options.H2D, 16, 10)
	case options.H2S != "":
		s := chPrefix(options.H2S)
		r, err := hex.DecodeString(s)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(r))
	case options.HelpFlag:
		fallthrough
	default:
		flag.PrintDefaults()
	}
}

func convertHexNum(origin string, base, targetBase int) {
	if strings.HasPrefix(origin, "0x") {
		origin = origin[2:]
	}
	i, err := strconv.ParseInt(origin, base, 64)
	if err != nil {
		fmt.Printf("parse number '%s' error for %s\n", origin, err.Error())
	} else {
		r := strconv.FormatInt(i, targetBase)
		fmt.Println(r)
	}
}

func chPrefix(s string) string {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	return s
}
