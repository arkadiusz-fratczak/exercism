package house

import "strings"

const testVersion = 1

var thisVerses = [12]string{
	"This is the house that Jack built.",
	"This is the malt",
	"This is the rat",
	"This is the cat",
	"This is the dog",
	"This is the cow with the crumpled horn",
	"This is the maiden all forlorn",
	"This is the man all tattered and torn",
	"This is the priest all shaven and shorn",
	"This is the rooster that crowed in the morn",
	"This is the farmer sowing his corn",
	"This is the horse and the hound and the horn",
}

var thatVerses = [12]string{
	"",
	"that lay in the house that Jack built.",
	"that ate the malt",
	"that killed the rat",
	"that worried the cat",
	"that tossed the dog",
	"that milked the cow with the crumpled horn",
	"that kissed the maiden all forlorn",
	"that married the man all tattered and torn",
	"that woke the priest all shaven and shorn",
	"that kept the rooster that crowed in the morn",
	"that belonged to the farmer sowing his corn",
}

func Song() string {
	result := Verse(1)
	for i := 2; i <= 12; i++ {
		result += "\n\n" + Verse(i)
	}
	return result
}

func Verse(no int) string {
	return strings.Join(VerseRec([]string{thisVerses[no-1]}, no), "\n")
}

func VerseRec(verses []string, no int) []string {
	if no <= 1 {
		return verses
	}
	return VerseRec(append(verses, thatVerses[no-1]), no-1)
}
