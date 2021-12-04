package core

import (
	"log"
	"os"
	"path/filepath"
)

func AbsPath(path string) (absPath string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func OpenFile(path string) (file *os.File) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
