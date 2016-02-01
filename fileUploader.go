package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Saves an image from the request
func saveImage(request *http.Request) (string, error) {
	imgName := generateRandomString(8)
	err := saveFileForm("image/", imgName, "img", request, []string{".jpeg", ".jpg", ".png", ".gif"})

	if err != nil {
		return "Unable to upload image!", err
	}
	return imgName, nil
}

// Saves a FileForm from the given request
func saveFileForm(dir, name, key string, req *http.Request, allowedExtensions []string) error {
	file, fileHeader, err := req.FormFile(key)
	if err != nil {
		panic(err)
		return err
	}

	ext := getExtension(fileHeader.Filename)

	defer file.Close()

	if contains(allowedExtensions, ext) || len(allowedExtensions) < 1 {
		content, _ := ioutil.ReadAll(file)
		saveFile(dir, name, ext, content)
		return nil
	}

	return errors.New("Unable to save file!")
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
