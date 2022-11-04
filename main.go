package main

import (
	"github.com/jalal-akbar/vocabs/model"
)

// var vocabs = []string{"avoid", "love"}

func main() {

	//vocabs := []string{"Avoid", "Love"}
	//getVocabs := model.GetVocab(vocabs)

	for matching := false; !matching; {
		guess := model.GetInputVocabs()

		matching = model.IsMatching(model.ListVocab, guess)
	}

}
