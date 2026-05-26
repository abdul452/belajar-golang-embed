# Belajar Golang Embed

Repositori ini berisi dokumentasi, implementasi kode, dan skenario pengujian untuk memahami fitur **Golang Embed** (dirilis sejak Go 1.16). Fitur ini memungkinkan kita untuk menanamkan (*compile-in*) file aset statis seperti teks, gambar, atau seluruh direktori folder secara langsung ke dalam satu file binary executable aplikasi Go.

---

## 💡 Mengapa Fitur Embed Sangat Penting?

* **Single Binary Deployment**: Aplikasi backend tidak lagi bergantung pada file eksternal yang rawan tertinggal saat proses distribusi (*deployment*). Cukup membawa satu file hasil *build*, aplikasi sudah bisa berjalan mandiri.
* **Mencegah Error Runtime**: Menghilangkan risiko terjadinya kendala sistem seperti `file not found` akibat kesalahan penulisan jalur (*path relativity*) pada server produksi.
* **Keamanan Konfigurasi**: Aset bawaan aplikasi terintegrasi langsung di level byte memori, menyulitkan manipulasi data statis dari luar secara tidak sengaja.

---

## 🏗️ Struktur Direktori Proyek

Proyek ini disusun dengan memisahkan layer pengujian lokal (*package-level test*) dan pengumpulan data global pada berkas utama (*root executable*):

```text
│── belajar-golang-embed/
│    ├── 01-embed-test/
│    │   ├── file/              # Folder lokal berisi file .txt eksperimen
│    │   ├── embed_test.go      # Unit testing fitur directive embed
│    │   ├── logo.png           # Aset gambar mentah (source)
│    │   └── version.txt        # Berkas teks informasi versi aplikasi
│    ├── file/                  # Folder aset statis tingkat root
│    │   ├── a.txt
│    │   ├── b.txt
│    │   └── c.txt
│    ├── go.mod                 # Modul utama aplikasi Go
│    ├── logo_copy.png          # Hasil eksperimen os.WriteFile
│    └── main.go                # Root file entrypoint aplikasi & global embed
```

## 🛠️ Aturan & Implementasi Kode
1. Standar Kode Baru (Go Modern)
Sejak fungsi di package `io/ioutil` dinyatakan usang (deprecated), proyek ini sepenuhnya bermigrasi menggunakan standar library bawaan terbaru yaitu `os.WriteFile` untuk memproses konversi byte hasil embed ke dalam bentuk file fisik.

2. Aturan Isolasi Folder Golang Embed
Directive `//go:embed` terikat secara ketat dengan aturan keamanan modular Go. Kompilator melarang keras penggunaan path relatif naik ke folder atasnya (`../`).

Batasan tersebut disiasati dengan meletakkan variabel penampung embed di berkas `main.go` (Root) agar mendapatkan hak akses legal membaca seluruh sub-direktori di bawahnya secara bebas (`file/*.txt` atau `01-embed-test/version.txt`).

## 🧪 Dokumentasi Kode Utama (main.go)
Berkas utama mendemonstrasikan 3 skenario tipe data penampung koleksi berkas statis:

```go
package main

import (
	"embed"
	"fmt"
	"os"
)

// 1. Embed Berkas Teks Tunggal ke Tipe Data String
//go:embed 01-embed-test/version.txt
var Version string

// 2. Embed Berkas Gambar/Binary ke Tipe Data Slice Byte []byte
//go:embed 01-embed-test/logo.png
var Logo []byte

// 3. Embed Multi-File atau Folder Menggunakan embed.FS (File System virtual)
//go:embed file/*.txt
var FilesTxt2 embed.FS

func main() {
	fmt.Println("Isi Version:", Version)

	// Menyimpan data byte logo menjadi berkas baru
	err := os.WriteFile("logo_copy.png", Logo, 0666)
	if err != nil {
		fmt.Println("Error write file:", err)
	}

	// Membaca isi direktori virtual dari filesystem embed.FS
	// Catatan: Gunakan prefix folder "file" saat mengakses isinya
	entries, err := FilesTxt2.ReadDir("file")
	if err != nil {
		fmt.Println("Error read dir:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println("File Terdeteksi:", entry.Name())
		isiByte, _ := FilesTxt2.ReadFile("file/" + entry.Name())
		fmt.Println("Konten:", string(isiByte))
	}
}
```

## 🚀 Kompilasi & Cara Menjalankan
Menjalankan Unit Pengujian
Untuk mengeksekusi fungsi uji coba lokal di dalam sub-folder package:

```Bash
go test -v ./01-embed-test/...
```

Setelah proses eksekusi file `./main` berjalan, berkas asli di dalam folder `file/` maupun `logo.png` bisa dihapus atau dipindahkan dari komputer, namun aplikasi akan tetap mampu mencetak data secara sempurna karena data tersebut sudah terkunci mati di dalam memory block file aplikasi hasil build.
