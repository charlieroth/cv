package cv

import (
	"encoding/json"
	"os"
)

type CV struct {
	Name       string            `json:"name"`
	Location   string            `json:"location"`
	Mission    string            `json:"mission"`
	Summary    string            `json:"summary"`
	Contact    Contact           `json:"contact"`
	Education  Education         `json:"education"`
	Employment []EmploymentEntry `json:"employment"`
}

type Contact struct {
	Email   string `json:"email"`
	Twitter string `json:"twitter"`
	GitHub  string `json:"github"`
}

type Education struct {
	Entity     string   `json:"entity"`
	Location   string   `json:"location"`
	Credential string   `json:"credential"`
	Start      int      `json:"start"`
	Stop       int      `json:"stop"`
	Details    []string `json:"details"`
}

type EmploymentEntry struct {
	Title    string   `json:"title"`
	Entity   string   `json:"entity"`
	Start    string   `json:"start"`
	Stop     string   `json:"stop"`
	Location string   `json:"location"`
	Details  []string `json:"details"`
}

func Parse(jsonFilePath string) (CV, error) {
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return CV{}, err
	}

	cv := CV{}
	err = json.Unmarshal(data, &cv)
	if err != nil {
		return CV{}, err
	}

	return cv, nil
}
