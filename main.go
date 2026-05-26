package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed 01-embed-test/version.txt
var version string

//go:embed 01-embed-test/logo.png
var logo []byte

//go:embed file/*.txt
var filesTxt2 embed.FS

func main() {
	fmt.Println(version)

	// Versi baru yang bersih tanpa dicoret:
	err := os.WriteFile("logo_copy.png", logo, fs.ModePerm) // fs.ModePerm bisa diganti dengan permission mentah seperti 0666
	if err != nil {
		fmt.Println(err)
	}

	entries, err := filesTxt2.ReadDir("file")
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			data, err := filesTxt2.ReadFile("file/" + entry.Name())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Isi file %s: %s\n", entry.Name(), string(data))
		}
	}
}
