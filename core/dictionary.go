package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

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

func GetDictionary(word string) (*Output, error) {
	api := os.Getenv("DICTIONARY_API") + word
	res, err := http.Get(api)
	if err != nil {
		log.Printf("Failed to send request to dictionry: %v", err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}

	var response Response
	s := string(resBody)
	s = `{"output":` + s + `}`
	b := []byte(s)
	result := new(Output)
	if err := json.Unmarshal(b, &response); err != nil {
		log.Printf("Failed to unmarshal response body: %s", resBody)
		var e ErrorResponse
		if err := json.Unmarshal(resBody, &e); err != nil {
			log.Printf("Failed to unmarshal error response: %v", err)
		}

		result.Error = &e
		return result, nil
	}

	dictionary := response.Output[0]
	for i := 0; i < len(dictionary.Meanings); i++ {
		if dictionary.Meanings[i].PartOfSpeech == "noun" {
			result.Noun = &dictionary.Meanings[i].Definitions[0].Definition
		}
		if dictionary.Meanings[i].PartOfSpeech == "verb" {
			result.Verb = &dictionary.Meanings[i].Definitions[0].Definition
		}
		if dictionary.Meanings[i].PartOfSpeech == "adverb" {
			result.Adverb = &dictionary.Meanings[i].Definitions[0].Definition
		}
		if dictionary.Meanings[i].PartOfSpeech == "adjective" {
			result.Adjective = &dictionary.Meanings[i].Definitions[0].Definition
		}
		if dictionary.Meanings[i].PartOfSpeech == "interjection" {
			result.Interjection = &dictionary.Meanings[i].Definitions[0].Definition
		}
	}
	return result, nil
}
