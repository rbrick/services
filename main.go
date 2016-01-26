package main

import (
	"math/rand"
	"net/http"
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

// Generate a pseudo-random string of x length
func generateRandomString(x int) string {
	result := make([]rune, x)
	for i := range result {
		result[i] = CHARS[rand.Intn(CHARS_LENGTH)]
	}
	return string(result)
}
