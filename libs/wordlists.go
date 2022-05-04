package libs

import (
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

var wordListsMap = map[string][]string{
	"chinese_simplified":  wordlists.ChineseSimplified,
	"chinese_traditional": wordlists.ChineseTraditional,
	"czech":               wordlists.Czech,
	"english":             wordlists.English,
	"french":              wordlists.French,
	"italian":             wordlists.Italian,
	"japanese":            wordlists.Japanese,
	"korean":              wordlists.Korean,
	"spanish":             wordlists.Spanish,
}

func SetWordList(lang string) {
	wordList, ok := wordListsMap[lang]
	if !ok {
		wordList = wordlists.English
	}

	bip39.SetWordList(wordList)
}
