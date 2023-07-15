package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	SecondTsLen = 10
	MilliTsLen  = 13
	MicroTsLen  = 16
	NanoTsLen   = 19
)

type CVTOption struct {
	HelpFlag  bool
	D2B       string
	D2H       string
	H2B       string
	H2D       string
	H2S       string
	TimeStamp string
	MD5       string
}

func parse() (*CVTOption, []string) {
	options := &CVTOption{}
	flag.BoolVar(&options.HelpFlag, "help", false, "Displays usage information and exit.")
	flag.StringVar(&options.D2B, "d2b", "", "Decimal number to Binary number.")
	flag.StringVar(&options.D2H, "d2h", "", "Decimal number to Hex number.")
	flag.StringVar(&options.H2B, "h2b", "", "Hex number to Binary number.")
	flag.StringVar(&options.H2D, "h2d", "", "Hex number to Decimal number.")
	flag.StringVar(&options.H2S, "h2s", "", "decode Hex string to utf8 string.")
	flag.StringVar(&options.TimeStamp, "ts", "", "convert timestamp to date time. (e.g. 1612345678)")
	flag.StringVar(&options.MD5, "md5", "", "md5 hash(16bytes). (e.g. 123456)")

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
	case options.TimeStamp != "":
		tss := options.TimeStamp
		ts, err := strconv.ParseInt(tss, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch len(tss) {
		case SecondTsLen:
		case MilliTsLen:
			ts = ts / 1000
		case MicroTsLen:
			ts = ts / 1000000
		case NanoTsLen:
			ts = ts / 1000000000
		default:
			fmt.Println("invalid timestamp")
			return
		}
		timeNow := time.Now().Second()
		fmt.Printf("current timestamp secondes: %d\n", timeNow)
		fmt.Printf("current timestamp mill-seconeds: %d\n", timeNow*1000)
		t := time.Unix(ts, 0).Format("2006-01-02 15:04:05")
		fmt.Println(t)
	case options.MD5 != "":
		s := options.MD5
		bs := md5.Sum([]byte(s))
		fmt.Println(hex.EncodeToString(bs[:]))
		fmt.Println(strings.ToUpper(hex.EncodeToString(bs[:])))
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
