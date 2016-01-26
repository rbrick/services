package main

import (
    "math/rand"
    "net/http"
    "time"
    "github.com/go-martini/martini"
    "os"
    "path/filepath"
    "errors"
)

var (
    CHARS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
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

    m.Get("/i/:id", func(response http.ResponseWriter, request *http.Request, params martini.Params) string {
        id := params["id"]
        file, err := findFile("image/", id + ".png")

        if err != nil {
            panic(err)
        }

        return file.Name()
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
