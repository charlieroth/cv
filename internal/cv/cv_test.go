package cv

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Correctly parses cv.json into CV{}", func(t *testing.T) {
		jsonFilePath := "/Users/charlie/github.com/charlieroth/cv/cv.json"

		cv, err := Parse(jsonFilePath)
		if err != nil {
			t.Error(err)
			return
		}

		if cv.Name != "Charles Roth III" {
			t.Errorf("expected parsed name to be 'Charles Roth III'")
		}

		if cv.Contact.Email != "charlieroth4@icloud.com" {
			t.Errorf("expected parsed email to be 'charlieroth4@icloud.com'")
		}

		if len(cv.Employment) < 1 {
			t.Errorf("expected employment to be at least length 1")
		}
	})

	t.Run("Fails to parse given invalid file path", func(t *testing.T) {
		jsonFilePath := "/invalid/file/path"

		_, err := Parse(jsonFilePath)
		if err == nil {
			t.Error("expected parse to file with invalid file path")
			return
		}
	})
}
