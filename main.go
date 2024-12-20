package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problemms.csv", "a csv file in thee format of 'a csv file in the format of 'quetion, answer'")
	flag.Parse()
	_ = csvFilename
}
