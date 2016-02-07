package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
		imgName, err := saveImage(request)
		if err != nil {
			response.Write([]byte(imgName))
		} else {
			response.Write([]byte(BASE_URL + "i/" + imgName))
		}
	})

	m.Get("/shorten", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "public/shortener.html")
	})

	/*
	  What really needs to happen:
	    Make an post request to this path ON THE JAVASCRIPT (CLIENT SIDE!)
	    Return JSON in the /shorten path
	    Parse the JSON in the javascript
	*/

	m.Post("/api/shorten", func(response http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()

		if err != nil {
			panic(err)
		}

		url := request.FormValue("longUrl")
		code, err := shorten(url)

		if err != nil {
			panic(err)
		}

		type ShortenUrl struct {
			Result string `json:"result"`
			AsUrl  string `json:"url"`
		}

		surl := ShortenUrl{code, "rbrickis.me/s/" + code}
		json, err := json.Marshal(surl)

		if err != nil {
			panic(err)
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(json)
	})

	m.Get("/s/:id", func(response http.ResponseWriter, request *http.Request, params martini.Params) {
		id := params["id"]

		if hasLongUrl(id) {
			redirectLink := getLongUrl(id)

			if !strings.HasPrefix(redirectLink, "http://") && !strings.HasPrefix(redirectLink, "https://") {
				redirectLink = "http://" + redirectLink
			}

			http.Redirect(response, request, redirectLink, 307)
		}
	})

	m.Get("/i/:id", func(response http.ResponseWriter, request *http.Request, params martini.Params) {
		id := params["id"]
		file, err := findFile("image/", id)

		if os.IsNotExist(err) {
			response.WriteHeader(404)
			http.ServeFile(response, request, "public/404.html")
			return
		}

		http.ServeFile(response, request, file.Name())
	})

	m.NotFound(func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "public/404.html")
	})

	m.RunOnAddr(":80")
}

func findFile(dir, fileName string) (*os.File, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	var found *os.File

	visit := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := getExtension(info.Name())

		if strings.Replace(info.Name(), ext, "", 1) == fileName {
			found, _ = os.Open(path)
		}
		return nil
	}

	filepath.Walk(dir, visit)

	if found == nil {
		return nil, os.ErrNotExist
	}

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

// Tests if a element is within a slice/array
func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
