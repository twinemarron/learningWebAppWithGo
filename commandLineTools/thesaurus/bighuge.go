package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verv *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Sunonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge: %q の類語検索に失敗しました: %v", term, err)
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verv.Syn...)
	return syns, nil
}
