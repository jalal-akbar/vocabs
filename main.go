package main

import (
	"github.com/jalal-akbar/vocabs/model"
)

func main() {
	for matching := false; !matching; {
		for i := 0; i < 100; i++ {
			guess := model.GetInputVocabs()
			matching = model.IsMatching(model.ListVocab, guess)
		}
	}
}
