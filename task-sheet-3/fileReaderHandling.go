package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileRead(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return err
	}
	return nil
}

func writeFile(fileName string, text string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {

		fmt.Println("Error writing File", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(text)
	if err != nil {
		fmt.Println("Error ", err)
		return err
	}
	if err := writer.Flush(); err != nil {
		fmt.Println("Error ", err)
		return err
	}
	return nil

}

func main() {
	fileName := "FILE.txt"
	if err := writeFile(fileName, "This is new File\n"); err != nil {
		fmt.Println("ERROR : ", err)
	}
	if err := fileRead(fileName); err != nil {
		fmt.Println("Error", err)

	}

}
