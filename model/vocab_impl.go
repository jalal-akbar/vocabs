package model

import (
	"fmt"
)

// for i := 0; i < len(vocabs); i++ {
// 	if vocabs[i] == guess {
// 		count = vocabs[i]
// 	}
// }
// fmt.Println(model.VocabTranslate[count])

func IsMatching(v []string, guess string) bool {
	res := ""
	for i := 0; i < len(v); i++ {
		if v[i] == guess {
			res = v[i]
			fmt.Println("Artinya :", VocabTranslate[res])
			return true
		}
	}
	fmt.Println("Coba Ulangi")
	return false
}

func GetInputVocabs() string {
	var input string
	fmt.Print("Masukkan Kata Yang ingin Di Translate : ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Masukkan 1 Kata")
	} else {
		fmt.Println("Kata yang Dimasukkan :", input)
	}
	return input

}
