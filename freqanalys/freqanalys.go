package freqanalys

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// FreqAnalys функция выполняет частотный анализ
func FreqAnalys(s string) []string {

	var mapWords = make(map[string]int)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(s, f)
	for _, w := range words {
		mapWords[strings.ToLower(w)]++
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range mapWords {
		if v > 0 {
			ss = append(ss, kv{k, v})
		}
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var result []string
	for _, s := range ss[:10] {
		fmt.Println(s)
		result = append(result, s.Key)
	}
	fmt.Println(result)
	return result
}
