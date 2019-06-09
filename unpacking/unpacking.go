package unpacking

import (
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Unpacking функция производит распаковку строки
func Unpacking(s string) string {
	var result strings.Builder
	if utf8.ValidString(s) && utf8.RuneCountInString(s) >= 2 {
		var lastrune rune   // Запоминаем последнюю руну
		var skipslash bool  // Если идет двойной слэш
		var countrepeat int // Количество повтрорений
		for _, r := range s {
			if unicode.IsDigit(r) {
				if lastrune == '\\' && !skipslash { // Если цифра идет после сплэша
					_, err := result.WriteRune(r)
					if err != nil {
						log.Fatal("Error: ", err)
					}
					lastrune = r
				} else {
					countrepeat = int(r-48) - 1
					if countrepeat > 0 && lastrune > 0 { // запись руны n-ное количество раз
						_, err := result.WriteString(strings.Repeat(string(lastrune), countrepeat))
						if err != nil {
							log.Fatal("Error: ", err)
						}
					}
					skipslash = false
				}
			} else if r == '\\' {
				if lastrune == '\\' { // второй нашли второй сплэш
					_, err := result.WriteRune(r) // то записываем его в строку
					if err != nil {
						log.Fatal("Error: ", err)
					}
					skipslash = true //  и ставим переключатель в True
				}
				lastrune = '\\'
			} else { // Обычная запись руны
				_, err := result.WriteRune(r)
				if err != nil {
					log.Fatal("Error: ", err)
				}
				lastrune = r
			}
		}
	}
	return result.String()
}
