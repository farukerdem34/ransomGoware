package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const DEBUG bool = true

var file_paths []string

func main() {
	// Key Generation
	key, err := GenerateKey()

	if DEBUG {
		fmt.Println("Key: ", hex.EncodeToString(key))
	}

	if err != nil {
		panic(err)
	}

	// Path Discovery
	root := "/home/kali/Documents/GoDemo2/"
	err2 := filepath.WalkDir(root, visitFile)
	if err2 != nil {
		fmt.Printf("Error walking the path: %v\n", err)
	}

	// Encryption Loop
	for _, path := range file_paths {
		fmt.Println("Path: ", path)
	}

}

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := io.ReadFull(rand.Reader, key)

	if err != nil {
		return nil, err
	}

	return key, nil
}

func visitFile(fp string, fi os.DirEntry, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}

	if DEBUG {
		if !fi.IsDir() {
			file_paths = append(file_paths, fp)
		}
	}
	return nil
}
