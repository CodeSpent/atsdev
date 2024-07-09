package parser

import (
	"atsdev/internal/models"
	"strings"
)

func ParseResume(text string) models.Resume {
	lines := strings.Split(text, "\n")
	var resume models.Resume

	for _, line := range lines {
		if strings.Contains(line, "Name:") {
			resume.Name = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Email:") {
			resume.Email = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Phone:") {
			resume.Phone = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Skills:") {
			skills := strings.TrimSpace(strings.Split(line, ":")[1])
			resume.Skills = strings.Split(skills, ",")
		} else if strings.Contains(line, "Experience:") {
			// Add logic for parsing experience
		} else if strings.Contains(line, "Education:") {
			// Add logic for parsing education
		}
	}

	return resume
}
