package main

import (
	"errors"
	"github.com/go-martini/martini"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	CHARS        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	CHARS_LENGTH = len(CHARS)
)

const BASE_URL = "http://rbrickis.me/"

func main() {
	rand.Seed(time.Now().UnixNano())

	m := martini.Classic()

	m.Use(martini.Static("js"))

	// Handles uploading of images
	m.Post("/upload_image", func(response http.ResponseWriter, request *http.Request) {
		imgName := saveImage(request)
		response.Write([]byte(BASE_URL + "i/" + imgName))
	})

	m.Get("/shorten", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "public/shortener.html")
	})

	m.Get("/i/:id", func(response http.ResponseWriter, request *http.Request, params martini.Params) {
		id := params["id"]
		fileName := "image/" + id + ".png"

		if _, err := os.Open(fileName); os.IsNotExist(err) {
			response.WriteHeader(404)
			http.ServeFile(response, request, "public/404.html")
			return
		}

		response.Header().Set("Content-Type", "image/png")
		http.ServeFile(response, request, fileName)
	})

	m.NotFound(func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "public/404.html")
	})

	m.RunOnAddr(":80")
}

func findFile(dir, fileName string) (*os.File, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, errors.New("Provided directory does not exist!")
	}

	var found *os.File

	visit := func(path string, info os.FileInfo, err error) error {
		if info.Name() == fileName {
			found, _ = os.Open(path)
		}
		return nil
	}

	filepath.Walk(dir, visit)
	return found, nil
}

// Generate a pseudo-random string of x length
func generateRandomString(x int) string {
	result := make([]rune, x)
	for i := range result {
		result[i] = CHARS[rand.Intn(CHARS_LENGTH)]
	}
	return string(result)
}
