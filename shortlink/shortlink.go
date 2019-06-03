package shortlink

import (
	"log"
	"net/url"
	"strconv"
)

// Shortener - интерфейс
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

// URLList - структура, содержит в себе базу(пока мапа), и текущий ключ.
// (написал реализацию что первое в голову пришло)
type URLList struct {
	db  map[string]string
	Key int64
}

// NewURLList - инициализация базы(на данный момент мапы)
func (u *URLList) NewURLList() {
	u.db = make(map[string]string)
}

// shoterAlgoritm - Алгоритм для сжатия url, пока выбрал самый простой в реализации
func shoterAlgoritm(u *URLList, urllong string) string {
	urlparse, err := url.Parse(urllong)
	if err != nil {
		log.Fatal(err)
	}
	u.Key++
	urlparse.Path = strconv.FormatInt(u.Key, 10)
	urlshort := urlparse.String()
	return urlshort
}

// Shorten - Сокращение url
func (u *URLList) Shorten(urllong string) string {
	if urlshort, exists := u.db[urllong]; exists {
		return urlshort
	}
	urlshort := shoterAlgoritm(u, urllong)
	u.db[urlshort] = urllong
	return urlshort
}

// Resolve - Расшифровка url
func (u *URLList) Resolve(urlshort string) string {
	if urllong, exists := u.db[urlshort]; exists {
		return urllong
	}
	return ""
}
