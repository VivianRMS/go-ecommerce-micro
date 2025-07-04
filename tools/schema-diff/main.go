package main

import (
	"flag"
	"fmt"
	"github.com/VivianRMS/go-ecommerce-micro/tools/schema-diff/diff"
	"log"
)

func main() {
	oldPath := flag.String("old", "", "path to old schema")
	newPath := flag.String("new", "", "path to new schema")
	flag.Parse()

	if *oldPath == "" || *newPath == "" {
		log.Fatal("Please specify --old and --new proto file paths.")
	}

	changes, err := diff.CompareProtoFiles(*oldPath, *newPath)

	if err != nil {
		log.Fatalf("Failed to compare proto files: %v", err)
	}

	if len(changes) == 0 {
		fmt.Println("✅ Compatible: No breaking changes found.")
	} else {
		fmt.Println("❗Potential Issues:")
		for _, c := range changes {
			fmt.Println("  -", c)
		}
	}
}
