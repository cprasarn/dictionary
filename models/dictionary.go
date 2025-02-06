package models

type ErrorResponse struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	Resolution string `json:"resolution"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Phonetic struct {
	Text      *string  `json:"text"`
	Audio     *string  `json:"audio"`
	SourceUrl *string  `json:"sourceUrl"`
	License   *License `json:"license"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
	Example    *string  `json:"example"`
}

type Meanings struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type DictionaryResponse struct {
	Word       string     `json:"word"`
	Phonetic   *string    `json:"phonetic"`
	Phonetics  []Phonetic `json:"phonetics"`
	Meanings   []Meanings `json:"meanings"`
	License    *License   `json:"license"`
	SourceUrls []string   `json:"sourceUrls"`
}

type Response struct {
	Output []DictionaryResponse `json:"output"`
}

type Output struct {
	Noun         *string        `json:"noun"`
	Verb         *string        `json:"verb"`
	Adverb       *string        `json:"adverb"`
	Adjective    *string        `json:"adjective"`
	Interjection *string        `json:"interjection"`
	Error        *ErrorResponse `json:"error"`
}
