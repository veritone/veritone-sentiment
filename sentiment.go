package sentiment

import (
	"math"
	"strings"
)

// SentimentAnalysis takes in a (possibly 'dirty')
// sentence (or any block of text,) cleans the
// text, finds the sentiment of each word in the
// text, finds the sentiment of the sentence as
// a whole, adn returns an Analysis struct
func (m Models) SentimentAnalysis(sentence string, lang Language) *Analysis {
	if _, ok := m[lang]; !ok {
		lang = English
	}

	analysis := &Analysis{
		Language: lang,
		Words:    []Score{},
		Score:    uint8(0),
	}

	sentences := strings.FieldsFunc(sentence, SplitSentences)
	if len(sentences) > 1 {
		analysis.Sentences = []SentenceScore{}

		for _, s := range sentences {
			class, P := m[lang].Probability(s)

			analysis.Sentences = append(analysis.Sentences, SentenceScore{
				Sentence:    s,
				Score:       ScaleScores(class, P, lang),
				Probability: P,
			})
		}
	}

	w := strings.Split(sentence, " ")
	for _, word := range w {
		class, P := m[lang].Probability(word)

		analysis.Words = append(analysis.Words, Score{
			Word:        word,
			Score:       ScaleScores(class, P, lang),
			Probability: P,
		})
	}

	analysis.Score = ScaleScores(m[lang].Predict(sentence), math.NaN(), lang)

	return analysis
}

// ScaleScores turns a class and it's probability
// into a descretized version on {0,1,...,10} so
// it'll be easier to plot
func ScaleScores(class uint8, probability float64, lang Language) uint8 {
	// scale P to reflect the values
	var P float64
	switch lang {
	case English:
		P = probability
		if class == uint8(0) {
			P = 1 - P
		}
		if math.IsNaN(probability) {
			P = float64(class)
		}
		P *= 10
	default:
		return uint8(0)
	}

	// bucket P into descrete values
	return uint8(math.Floor(P))
}
