package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	cr "crypto/rand"
	"fmt"
	"io"
	"os"
)

func aesEncrypt(plainText []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error , while creating block", err)
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plainText))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(cr.Reader, iv); err != nil {
		fmt.Println("Error while reading iv: ", err)
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)
	return ciphertext, nil

}
func aesDecrypt(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error while creating block (decrypt): ", err)
		return nil, err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	// plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func main() {
	// var key []byte
	key := []byte("My Secret Key 12")

	// test := []byte("This is test string")
	// pathDir := "C:\\Users\\surya\\OneDrive\\Desktop\\Files\\Codes\\Go\\task-sheet-4\\SampleDir\\BhagDkBose.png"
	// file, err := os.Open(pathDir)
	// if err != nil {
	// 	fmt.Println("Error while opening file", err)
	// 	return
	// }
	// defer file.Close()
	// reader := bufio.NewReader(file)
	// data, err := io.ReadAll(reader)

	// if err != nil {
	// 	fmt.Println("Error , while reading data ", err)
	// 	return
	// }
	// result, err := aesEncrypt(data, key)
	// if err != nil {
	// 	fmt.Println("Error in encrypting: ", err)
	// 	return
	// }
	// encFile, err := os.Create("encrypted.dat")
	// if err != nil {
	// 	fmt.Println("Error while creating file: ", err)
	// 	return
	// }
	// defer encFile.Close()

	// writer := bufio.NewWriter(encFile)
	// _, ERR := writer.Write(result)
	// if ERR != nil {
	// 	fmt.Println("Error , while writing: ", ERR)
	// 	return
	// }
	// writer.Flush()

	// fmt.Println(string(result))

	// decryptedResult, errR := aesDecrypt(result, key)
	// if errR != nil {
	// 	fmt.Println("Error in Decrypting: ", errR)
	// 	return
	// }

	decFile, err := os.Open("encrypted.dat")
	if err != nil {
		fmt.Println("Error while opening :", err)
		return
	}
	defer decFile.Close()

	encReader := bufio.NewReader(decFile)

	encData, err := io.ReadAll(encReader)
	if err != nil {
		fmt.Println("Error while reading", err)
		return
	}

	decryptedData, err := aesDecrypt(encData, key)
	if err != nil {
		fmt.Println("error while decrypting", err)
		return
	}

	decOutputFile, err := os.Create("decrypted.png")
	if err != nil {
		fmt.Println("Error while creating output ", err)
	}
	defer decOutputFile.Close()

	decWriter := bufio.NewWriter(decOutputFile)

	_, ERRORS := decWriter.Write(decryptedData)
	if ERRORS != nil {
		fmt.Println("error while writing data ", ERRORS)
		return
	}

	decWriter.Flush()

	// fmt.Println(string(decryptedResult))

}
