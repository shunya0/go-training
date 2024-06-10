// Reading file with ReadAll() can work for small size but will lag for file with biggerr size

package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	cr "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	mr "math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	mr.Seed(time.Now().UnixNano())
}

func generateUniqueString() string {
	timestamp := time.Now().UnixNano()
	randomString := generateRandomString(8) // Adjust the length of the random string as needed
	return fmt.Sprintf("%d-%s", timestamp, randomString)
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[mr.Intn(len(letterBytes))]
	}
	return string(b)
}

func xorEncryptDecrypt(arr []byte, key byte) []byte {

	for i, v := range arr {
		arr[i] = v ^ key
	}
	return arr

}

func aesEncrypt(arr []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	// fmt.Println(aes.BlockSize)

	if err != nil {
		fmt.Println("Error while creating block(encrypting): ", err)
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(arr))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cr.Reader, iv); err != nil {
		fmt.Println("Error while reading: ", err)
		return nil, err
	}
	// fmt.Println(iv)
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], arr)

	return ciphertext, nil

}

func aesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 32 && len(key) != 24 {
		fmt.Println("Key length invalid")
		return nil, fmt.Errorf("invalid key length")
	}
	// fmt.Println(aes.BlockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error while creating block (decrypting): ", err)
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:] // uncommented
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("Cipher text too short")
	}
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext) //arr[aes.blocksize:] instead of encrypteddata

	return ciphertext, nil
}

func aesMain(filePath string, wg *sync.WaitGroup, flag bool, key []byte) bool {
	defer wg.Done()

	dirPathSlcie := strings.Split(filePath, "\\")
	dirPath := strings.Join(dirPathSlcie[:len(dirPathSlcie)-1], "\\")
	fileName := dirPathSlcie[len(dirPathSlcie)-1]
	ext := fileName[len(fileName)-4:]

	inputFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error while opening file (aesMain): ", err)
		return false
	}

	defer inputFile.Close()

	newUUID := generateUniqueString()
	outputFileName := newUUID + ext
	outputFile, err := os.Create(dirPath + "\\" + outputFileName)
	if err != nil {
		fmt.Println("Error while creating file(aesMain): ", err)
		return false
	}
	defer outputFile.Close()
	// content, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Println("ERROR while reading file :", err)
	}

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	data, err := io.ReadAll(reader)

	if flag {
		result, err := aesEncrypt(data, key)
		if err != nil {
			fmt.Println("error in encryption: ", err)
			return false
		}
		if _, err = writer.Write(result); err != nil {
			fmt.Println("Error while writing: ", err)
			return false
		}
		writer.Flush()
	} else {
		decryptedResult, err := aesDecrypt(data, key)
		if err != nil {
			fmt.Println("Error while decrypting: ", err)
		}
		if _, err := writer.Write(decryptedResult); err != nil {
			fmt.Println("Error while writing decrypted result: ", err)
			return false
		}
		writer.Flush()
	}
	// reader := bufio.NewReader(inputFile)

	// for {
	// 	buf := make([]byte, 4096)

	// 	data, err := reader.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		fmt.Println("Error while reading data(aesMain): ", err)
	// 		return false
	// 	}
	// 	if data == 0 {
	// 		break
	// 	}
	// 	var arr []byte
	// 	if flag {
	// 		arr, err = aesEncrypt(buf[:data], key)
	// 		if err != nil {
	// 			fmt.Println("Error while encrypting(aesMain): ", err)
	// 			return false
	// 		}
	// 		if _, err := writer.Write(arr); err != nil {
	// 			fmt.Println("Error while writing(aesMain): ", err)
	// 			return false

	// 		}
	// 	} else {
	// 		// fmt.Println("decrypting")
	// 		arr, err = aesDecrypt(buf[:data], key)
	// 		if err != nil {
	// 			fmt.Println("Error while decrypting(aesMain): ", err)
	// 			return false
	// 		}
	// 		if _, err := writer.Write(arr); err != nil {
	// 			fmt.Println("Error while writing(aesMain): ", err)
	// 			return false

	// 		}

	// 	}
	// }

	// if err := writer.Flush(); err != nil {
	// 	fmt.Errorf("failed to flush file: ", err)
	// 	return false
	// }

	inputFile.Close()
	outputFile.Close()
	if err := os.Remove(filePath); err != nil {
		fmt.Println("Error while removing file(aesMain): ", err)
		return false
	}
	if err := os.Rename(dirPath+"\\"+outputFileName, filePath); err != nil {
		fmt.Println("Error while renaming the file: ", err)
		return false
	}

	fmt.Println("Special Program finished executing for: ", filePath)
	return true
}
func Encrypt(filePath string, wg *sync.WaitGroup) bool {
	defer wg.Done()

	dirPathSlice := strings.Split(filePath, "\\")
	dirPath := strings.Join(dirPathSlice[:len(dirPathSlice)-1], "\\")
	fileName := dirPathSlice[len(dirPathSlice)-1]
	ext := fileName[len(fileName)-4:]

	// fmt.Println(dirPath)

	// fmt.Println("\n , went through , ", fileName)

	inputFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return false
	}

	defer inputFile.Close()

	newUUID := generateUniqueString()

	outputFileName := newUUID + ext
	outputFile, err := os.Create(dirPath + "\\" + outputFileName)
	if err != nil {
		fmt.Println("Error while creating output file: ", err)
		return false
	}

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	buf := make([]byte, 4096)
	for {
		data, err := reader.Read(buf)

		if err != nil && err != io.EOF {
			fmt.Println("Error while reading file: ", err)
			return false
		}
		if data == 0 {
			break
		}
		arr := xorEncryptDecrypt(buf[:data], 234)
		if _, err := writer.Write(arr); err != nil {
			fmt.Println("Error while writing file: ", err)
			return false
		}

	}
	if err := writer.Flush(); err != nil {
		fmt.Errorf("failed to flush file", err)
		return false
	}
	inputFile.Close()
	outputFile.Close()
	if err := os.Remove(filePath); err != nil {
		fmt.Println("Error while removing original file: ", err)
	}
	if err := os.Rename(dirPath+"\\"+outputFileName, filePath); err != nil {
		fmt.Println("Error while renaming output file: ", err)
	}
	fmt.Println("Special Program finished execution for: ", filePath)
	return true
}
func encryptDir(dirPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(dirPath, func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking the ", dirPath+" ", err)
			return err
		}
		if fileinfo.IsDir() {
			return nil
		}
		wg.Add(1)
		fmt.Println("Special program executing for: ", path)
		// go Encrypt(path, wg)
		key := sha256.Sum256([]byte("My Secret Key 12My Secret Key 12"))
		go aesMain(path, wg, true, key[:])
		return nil
	})
	if err != nil {
		fmt.Println("Error walking ", dirPath+" 2")
	}
}

func main() {

	dirPath := "C:\\Users\\surya\\OneDrive\\Desktop\\Files\\Codes\\Go\\task-sheet-4\\SampleDir"
	var wg sync.WaitGroup

	wg.Add(1)
	go encryptDir(dirPath, &wg)

	wg.Wait()
	fmt.Println("Successfully encrypted/Decrypted")

}
