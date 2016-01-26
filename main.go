package main

import (
	"math/rand"
	"net/http"
	"time"
    "github.com/go-martini/martini"
)

var (
	CHARS        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	CHARS_LENGTH = len(CHARS)
)

const BASE_URL = "http://rbrickis.me/"

func main() {
	rand.Seed(time.Now().UnixNano())

    m := martini.Classic()

	// Handles uploading of images
	m.Post("/upload_image", func(response http.ResponseWriter, request *http.Request) {
		imgName := saveImage(request)
		response.Write([]byte(BASE_URL + "i/" + imgName))
	})

	m.Get("/i", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte(request.RequestURI))
	})

	m.RunOnAddr(":80")
}

// Generate a pseudo-random string of x length
func generateRandomString(x int) string {
	result := make([]rune, x)
	for i := range result {
		result[i] = CHARS[rand.Intn(CHARS_LENGTH)]
	}
	return string(result)
}
