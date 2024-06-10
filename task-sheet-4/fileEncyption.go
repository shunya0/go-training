package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func encryptDecrypt(arr []byte, key byte) []byte {
	var xorKey byte
	// xorKey := 0xAA
	xorKey = key

	for i, v := range arr {
		arr[i] = v ^ xorKey
	}
	return arr

}

func readWriteFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("ERROR: ", err)
		// return nil, err
	}

	defer file.Close()
	var encryptedlines [][]byte

	// buf := make([]byte , 4096)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Bytes()
		line = encryptDecrypt(line, 212)
		encryptedline := make([]byte, len(line))
		copy(encryptedline, line)
		encryptedlines = append(encryptedlines, encryptedline)
		// fmt.Println(encryptedlines)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR: ", err)
		// return nil, err
	}
	if err := file.Truncate(0); err != nil {
		fmt.Println("ERROR: ", err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("ERROR: ", err)
	}

	write := bufio.NewWriter(file)

	for _, v := range encryptedlines {
		encryptString := string(v)

		fmt.Println(encryptString)
		if _, err := write.WriteString(encryptString); err != nil {
			fmt.Println("ERROR: ", err)
		}
		if _, err := write.WriteString("\n"); err != nil {
			fmt.Println("ERROR: ", err)
		}

		if err := write.Flush(); err != nil {
			fmt.Println("ERROR: ", err)
		}
	}

}

func main() {
	// fileName := "file.txt"
	// readWriteFile(fileName)
	dirPath := "C:\\Users\\surya\\OneDrive\\Desktop\\Files\\Codes\\Go\\task-sheet-4\\SampleDir"
	

	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer dir.Close()
	entries, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("ERROR: ", err)

	}
	for _, v := range entries {
		// fmt.Println(v)
		ext := filepath.Ext(v.Name())
		if strings.ToLower(ext) == ".txt" {
			readWriteFile(dirPath + "\\" + v.Name())
		}
	}

}
