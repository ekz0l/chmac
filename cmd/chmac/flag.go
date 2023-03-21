package main

import "flag"

var (
	verbose = flag.Bool("v", false, "verbose output")
)

func init() {
	flag.Parse()
}
