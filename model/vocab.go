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

var VocabTranslate = map[string]string{
	"avoid":       "Menghindari, Mencegah",
	"love":        "Cinta",
	"approach":    "Mendekati",
	"capability":  "kemampuan",
	"constraints": "kendala",
	"define":      "menetapkan",
	"enroll":      "mendaftar",
}

var ListVocab = []string{"love", "avoid"}
