package core

import (
	"dictionary/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func GetDictionary(word string) (*models.Output, error) {
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

	var response models.Response
	s := string(resBody)
	s = `{"output":` + s + `}`
	b := []byte(s)
	result := new(models.Output)
	if err := json.Unmarshal(b, &response); err != nil {
		log.Printf("Failed to unmarshal response body: %s", resBody)
		var e models.ErrorResponse
		if err := json.Unmarshal(resBody, &e); err != nil {
			log.Printf("Failed to unmarshal error response: %v", err)
		}

		result.Error = &e
		return result, nil
	}

	for i := 0; i < len(response.Output); i++ {
		dictionary := response.Output[i]
		for j := 0; j < len(dictionary.Meanings); j++ {
			if result.Noun == nil && dictionary.Meanings[j].PartOfSpeech == "noun" {
				result.Noun = &dictionary.Meanings[j].Definitions[0].Definition
			}
			if result.Verb == nil && dictionary.Meanings[j].PartOfSpeech == "verb" {
				result.Verb = &dictionary.Meanings[j].Definitions[0].Definition
			}
			if result.Adverb == nil && dictionary.Meanings[j].PartOfSpeech == "adverb" {
				result.Adverb = &dictionary.Meanings[j].Definitions[0].Definition
			}
			if result.Adjective == nil && dictionary.Meanings[j].PartOfSpeech == "adjective" {
				result.Adjective = &dictionary.Meanings[j].Definitions[0].Definition
			}
			if result.Interjection == nil && dictionary.Meanings[j].PartOfSpeech == "interjection" {
				result.Interjection = &dictionary.Meanings[j].Definitions[0].Definition
			}
		}
	}
	return result, nil
}
