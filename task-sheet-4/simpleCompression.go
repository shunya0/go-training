package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func deCompress(inputFile string, outPutfileName string) error {
	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return err
	}
	defer readFile.Close()

	outputeFile, err := os.Create(outPutfileName)
	if err != nil {
		fmt.Println("Error while creting output file: ", err)
		return err
	}
	defer outputeFile.Close()

	gzipReader, err := gzip.NewReader(readFile)
	if err != nil {
		fmt.Println("Error while reading the output file: ", err)
		return err
	}
	defer gzipReader.Close()

	if _, err := io.Copy(outputeFile, gzipReader); err != nil {
		fmt.Println("Error , while copying file: ", err)
		return err
	}
	return nil
}

func compress(inputFile string, outPutfileName string) error {

	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return err
	}
	defer readFile.Close()

	outputFile, err := os.Create(outPutfileName)
	if err != nil {
		fmt.Println("Error while creating output file: ", err)
		return err
	}
	defer outputFile.Close()

	buf := make([]byte, 4096)

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	for {
		data, err := readFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error while rading buffer file: ", err)
			return err
		}
		if data == 0 {
			break
		}
		if _, err := gzipWriter.Write(buf[:data]); err != nil {
			fmt.Println("Error while copying file: ", err)
			return err
		}
		if err := gzipWriter.Flush(); err != nil {
			fmt.Println("Error while flushing: ", err)
			return err
		}
	}

	return nil

}

func main() {
	fileName := "Betty_Boop.mp4"

	outputFile := "output.mp4.gzip"
	if err := compress(fileName, outputFile); err != nil {
		fmt.Println("Error , while compressing file: ", err)
		return
	}
	fmt.Println("========================\nCompression Done!\n========================")

	// inputOne := "decompressed.png"
	// if err := deCompress(outputFile, inputOne); err != nil {
	// 	fmt.Println("Error while decompressing file: ", err)
	// 	return
	// }
	// fmt.Println("========================\nDecompression Done!\n========================")

}
`