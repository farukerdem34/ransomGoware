package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const DEBUG bool = false

var file_paths []string

const ROOT string = "/path/to/root"

const isFile bool = false

func main() {
	key := whatTheKey(isFile)
	if DEBUG {
		fmt.Println("Key: ", []byte(key))
	}
	err := filepath.WalkDir(ROOT, visitFile)
	if err != nil {
		panic(err)
	}

	for _, path := range file_paths {
		if DEBUG {
			fmt.Println(path)
		}
		var cypherText []byte = readFile(path)
		if DEBUG {
			fmt.Println("Encrpyted Content: ", cypherText)
		}
		decrypted_content, _ := decrypter(cypherText, []byte(key))
		if DEBUG {
			fmt.Println("Decrypted Content: ", string(decrypted_content))
		}
	}

}

func whatTheKey(isFile bool) []byte {
	var key []byte
	if isFile {
		key = readFile("key.txt")
	} else {
		key = []byte("YOUR-32-BIT-KEY-STRING")
	}
	return key
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

	return content
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

func decrypter(cypherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	plainText, err := gcm.Open(nil, cypherText[:gcm.NonceSize()], cypherText[gcm.NonceSize():], nil)
	if err != nil {
		panic(err)
	}
	return plainText, nil
}
