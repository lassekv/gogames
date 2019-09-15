package cyoa

import (
	"encoding/json"
)

type Adventure map[string]StoryArc

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// ReadAdventure reads and parses the JSON formated aventure in the file adventure and returns an unmarshalled version
func ReadAdventure(adventure string) (adv Adventure, err error) {

	byteValue, err := Asset(adventure)
	if err != nil {
		return Adventure{}, err
	}
	err = json.Unmarshal(byteValue, &adv)
	if err != nil {
		return Adventure{}, err
	}
	return adv, nil
}
