package main

import (
	"flag"
	"log"

	"github.com/kidoda/unipuzzy/puzzler"
	"github.com/robmerrell/comandante"
	"github.com/spf13/pflag"
)

func init() {
	flags()
}

func flags() {

}

func main() {
	app := comandante.New("unipuzzy", "Multi-approach n-piece puzzle solver.")
	cmd1 := comandante.NewCommand("by", "algorithm or method by which to solve puzzle")
}
