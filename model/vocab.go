package model

type Vocab struct {
	ID    string
	Title string
}

var Vocabs = []Vocab{
	{
		ID:    "A",
		Title: "Avoid",
	},
	{
		ID:    "L",
		Title: "Love",
	},
}
