package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	// Simpan file logo_copy.png dengan isi yang sama dengan logo.png
	// ini versi lama go 1.16
	// err := ioutil.WriteFile("logo_copy.png", logo, fs.ModePerm)

	// Versi baru yang bersih tanpa dicoret:
	err := os.WriteFile("logo_copy.png", logo, fs.ModePerm) // fs.ModePerm bisa diganti dengan permission mentah seperti 0666
	if err != nil {
		t.Fatal(err)
	}
}

//go:embed file/a.txt
//go:embed file/b.txt
//go:embed file/c.txt
var filesTxt embed.FS

func TestMultipleFile(t *testing.T) {
	// Baca file a.txt
	dataA, err := filesTxt.ReadFile("file/a.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dataA))

	// Baca file b.txt
	dataB, err := filesTxt.ReadFile("file/b.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dataB))

	// Baca file c.txt
	dataC, err := filesTxt.ReadFile("file/c.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dataC))
}

//go:embed file/*.txt
var filesTxt2 embed.FS

func TestMultipleFileWithPattern(t *testing.T) {
	entries, err := filesTxt2.ReadDir("file")
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			data, err := filesTxt2.ReadFile("file/" + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("Isi file %s: %s\n", entry.Name(), string(data))
		}
	}
}
