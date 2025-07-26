package main

import (
	"flag"
	"fmt"
	"os"
	
	"github.com/debarbarinantoine/go-enum-generate/internal"
)

var Version = "dev"

func main() {
	
	isOverwrite := flag.Bool("force", false, "Overwrite existing enum files without prompting")
	flag.BoolVar(isOverwrite, "f", false, "Overwrite existing enum files without prompting")
	
	isHelp := flag.Bool("help", false, "Show help")
	flag.BoolVar(isHelp, "h", false, "Show help")
	
	showVersion := flag.Bool("version", false, "Show version information and exit")
	flag.BoolVar(showVersion, "v", false, "Show version information and exit")
	
	flag.Parse()
	
	if *isHelp {
		flag.Usage()
		os.Exit(0)
	}
	
	if *showVersion { // Handle the new --version flag
		fmt.Printf("%s v%s\n", os.Args[0], Version) // Print the binary name and the embedded version
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
			fmt.Printf(":: go-enum-generate: [ERROR] failed to generate enum %s\n", enum.Name)
			continue
		}
		err = enum.CreateEnumFile(*isOverwrite)
		if err != nil {
			fmt.Printf(":: go-enum-generate: [ERROR] failed to writing file for enum %s\n", enum.Name)
			continue
		}
	}
	
	fmt.Println(":: go-enum-generate: [INFO] enum files generation ended successfully")
}
