package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	CHARS        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	CHARS_LENGTH = len(CHARS)
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Handles uploading of images
	http.HandleFunc("/upload_image", func(response http.ResponseWriter, request *http.Request) {
		saveImage(request)
	})

	http.ListenAndServe(":80", nil)
}

// Saves an image from the request
func saveImage(request *http.Request) {
	saveFileForm("image/", "img", request)
}

// Saves a FileForm from the given request
func saveFileForm(dir, key string, req *http.Request) {
	file, fileHeader, err := req.FormFile(key)
	if err != nil {
		panic(err)
		return
	}

	ext := getExtension(fileHeader.Filename)

	defer file.Close()
	content, err := ioutil.ReadAll(file)

	saveFile(dir, generateRandomString(8), ext, content)
	return
}

// Saves a file within the given dir, with the name and extension, and will contain the given contents
func saveFile(dir, name, extension string, content []byte) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModeDir)
	}

	newFile, err := os.Create(dir + "/" + name + extension)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	newFile.Write(content)
}

// Get the file extension based off the file name
func getExtension(filename string) string {
	return filename[strings.LastIndex(filename, "."):]
}

// Generate a pseudo-random string of x length
func generateRandomString(x int) string {
	result := make([]rune, x)
	for i := range result {
		result[i] = CHARS[rand.Intn(CHARS_LENGTH)]
	}
	return string(result)
}
