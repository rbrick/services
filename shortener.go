package main

import (
	"fmt"
	"github.com/gosexy/redis"
	"log"
)

var client *redis.Client

func init() {
	client = redis.New()
	err := client.Connect("localhost", 6379)

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to the Redis Database")
	}
}

func shorten(longUrl string) (string, error) {
	if hasShortUrl(longUrl) {
		return getShortUrl(longUrl), nil
	}

	code := generateRandomString(5)

	_, err := client.HSet("longToShort", longUrl, code)

	if err != nil {
		return "", err
	}
	_, err = client.HSet("shortToLong", code, longUrl)

	if err != nil {
		return "", err
	}

	return code, nil
}

func getLongUrl(id string) string {
	if !hasLongUrl(id) {
		return "URL Not Found!"
	}
	return hget("shortToLong", id)
}

func hasShortUrl(longUrl string) bool {
	return hexist("longToShort", longUrl)
}

func hasLongUrl(id string) bool {
	return hexist("shortToLong", id)
}

func getShortUrl(longUrl string) string {
	if !hasShortUrl(longUrl) {
		return "URL Not Found!"
	}
	return hget("longToShort", longUrl)
}

func hexist(key, field string) bool {
	exists, err := client.HExists(key, field)
	if err != nil {
		return false
	}
	return exists
}

func hget(key, field string) string {
	result, err := client.HGet(key, field)
	if err != nil {
		return "An error has occurred!"
	}
	return result
}
