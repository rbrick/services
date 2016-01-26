package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Saves an image from the request
func saveImage(request *http.Request) string {
	imgName := generateRandomString(8)
	saveFileForm("image/", imgName, "img", request)
	return imgName
}

// Saves a FileForm from the given request
func saveFileForm(dir, name, key string, req *http.Request) {
	file, fileHeader, err := req.FormFile(key)
	if err != nil {
		panic(err)
		return
	}

	ext := getExtension(fileHeader.Filename)

	defer file.Close()
	content, err := ioutil.ReadAll(file)

	saveFile(dir, name, ext, content)
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
