package cyoa

import (
	"testing"
)

func TestReadAdventure(t *testing.T) {
	fname := "adventures/gopher.json"
	adv, err := ReadAdventure(fname)
	if len(adv) != 7 {
		t.Errorf("There should have been %d story arcs, but there was only %d", 7, len(adv))
	}
	if err != nil {
		t.Errorf("There was an error %v", err)
	}
}

func TestReadFileDoesNotExist(t *testing.T) {
	fname := "unknown_file.json"
	adv, err := ReadAdventure(fname)
	if len(adv) != 0 {
		t.Errorf("There should be no data in adv")
	}
	if err != nil && err.Error() != "Asset unknown_file.json not found" {
		t.Errorf("Wrong error message returned")
	}
}
