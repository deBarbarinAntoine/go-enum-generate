package main

import (
	"flag"
	"fmt"
	"os"
	
	"github.com/debarbarinantoine/go-enum-generate/internal"
)

func main() {
	
	isOverwrite := flag.Bool("force", false, "Overwrite existing enum files without prompting")
	flag.BoolVar(isOverwrite, "f", false, "Overwrite existing enum files without prompting")
	isHelp := flag.Bool("help", false, "Show help")
	flag.BoolVar(isHelp, "h", false, "Show help")
	
	flag.Parse()
	
	if *isHelp {
		flag.Usage()
		os.Exit(0)
	}
	
	enums, err := internal.GetEnums()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for _, enum := range enums {
		err = enum.Generate()
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = enum.CreateEnumFile(*isOverwrite)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	
	fmt.Println("Enums generated")
}
