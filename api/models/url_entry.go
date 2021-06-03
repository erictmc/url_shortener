package models

import (
	"crypto/rand"
)

type UrlEntry struct {
	ShortUrl string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}

func (adb *AppDB) CreateUrlEntry(origUrl string) (*UrlEntry, error) {
	u := UrlEntry{}
	u.OriginalUrl = origUrl

	// TODO: For creating a random tiny url, this should be left to an outside service, to prevent name collisions.
	newUrl, _ := GenerateRandomString(10)
	u.ShortUrl = newUrl
	if err := adb.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (adb *AppDB) FetchUrlEntry(shortUrl string) (*UrlEntry, error) {
	urlEntry := UrlEntry{}
	if err := adb.DB.Where("short_url = ?", shortUrl).First(&urlEntry).Error; err != nil {
		return nil, err
	}

	return &urlEntry, nil
}

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}


func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := Bytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}