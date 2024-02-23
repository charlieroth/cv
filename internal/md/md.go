package md

import (
	"fmt"
	"strings"

	"github.com/charlieroth/cv/internal/cv"
)

func Generate(cv *cv.CV) (string, error) {
	md := strings.Builder{}

	// Title
	name := fmt.Sprintf("# %s\n\n", cv.Name)
	_, err := md.WriteString(name)
	if err != nil {
		return "", err
	}

	// Location
	location := fmt.Sprintf("**Location**: %s\n\n", cv.Location)
	_, err = md.WriteString(location)
	if err != nil {
		return "", err
	}

	// Mission
	mission := fmt.Sprintf("**Mission**: %s\n\n", cv.Mission)
	_, err = md.WriteString(mission)
	if err != nil {
		return "", err
	}

	// Summary
	summary := fmt.Sprintf("**Summary**: %s\n\n", cv.Summary)
	_, err = md.WriteString(summary)
	if err != nil {
		return "", err
	}

	contact := fmt.Sprintf(`### Contact Information

Email: %s

Twitter: %s

GitHub: %s

`, cv.Contact.Email, cv.Contact.Twitter, cv.Contact.GitHub)
	_, err = md.WriteString(contact)
	if err != nil {
		return "", err
	}

	// Education
	education := fmt.Sprintf(`## Education

### %s

Location: %s

Degree: %s

Start: %d

Stop: %d

**Details**

`, cv.Education.Entity, cv.Education.Location, cv.Education.Credential, cv.Education.Start, cv.Education.Stop)
	_, err = md.WriteString(education)
	if err != nil {
		return "", err
	}

	// Education Details
	for _, educationDetail := range cv.Education.Details {
		detail := fmt.Sprintf("- %s\n", educationDetail)
		_, err = md.WriteString(detail)
		if err != nil {
			return "", err
		}
	}

	_, err = md.WriteString("\n")
	if err != nil {
		return "", err
	}

	// Employment
	_, err = md.WriteString("## Employment\n\n")
	if err != nil {
		return "", err
	}

	// Employment Entries
	for _, employmentEntry := range cv.Employment {
		err = writeEmploymentEntry(&md, employmentEntry)
		if err != nil {
			return "", err
		}
	}

	return md.String(), nil
}

func writeEmploymentEntry(s *strings.Builder, entry cv.EmploymentEntry) error {
	entryString := fmt.Sprintf(`### %s - %s

Start: %s

Stop: %s

Location: %s

#### Details

`, entry.Entity, entry.Title, entry.Start, entry.Stop, entry.Location)

	_, err := s.WriteString(entryString)
	if err != nil {
		return err
	}

	for _, entryDetail := range entry.Details {
		detail := fmt.Sprintf("- %s\n", entryDetail)
		_, err = s.WriteString(detail)
		if err != nil {
			return err
		}
	}

	_, err = s.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}
