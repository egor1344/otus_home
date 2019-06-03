package shortlink

import "testing"

func TestShorten(t *testing.T) {
	var urllist URLList
	urllist.NewURLList()
    firstShortenURL := urllist.Shorten("https://otus.ru/learning/24319/")
    if firstShortenURL != "https://otus.ru/1" {
        t.Error("Wrong result ", firstShortenURL)
	}
	secondtShortenURL := urllist.Shorten("https://otus.ru/1234123421/1234/")
    if secondtShortenURL != "https://otus.ru/2" {
        t.Error("Wrong result ", secondtShortenURL)
    }
}

func TestResolve(t *testing.T) {
	var urllist URLList
	urllist.NewURLList()
	shortlink := urllist.Shorten("https://otus.ru/learning/24319/")
	if shortlink != "https://otus.ru/1" {
		t.Error("Wrong result ", shortlink)
	}
	normallink := urllist.Resolve(shortlink)
	if normallink != "https://otus.ru/learning/24319/" {
		t.Error("Wrong result ", normallink)
	}
}