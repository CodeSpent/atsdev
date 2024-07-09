package parser

import (
	"atsdev/internal/models"
	"strings"
)

func ParseJobDescription(text string) models.JobListing {
	lines := strings.Split(text, "\n")
	var jobDescription models.JobListing

	for _, line := range lines {
		if strings.Contains(line, "Title:") {
			jobDescription.Title = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Company:") {
			jobDescription.Company = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Skills:") {
			skills := strings.TrimSpace(strings.Split(line, ":")[1])
			jobDescription.Skills = strings.Split(skills, ",")
		} else if strings.Contains(line, "Responsibilities:") {
			// Add logic for parsing responsibilities
		} else if strings.Contains(line, "Requirements:") {
			// Add logic for parsing requirements
		}
	}

	return jobDescription
}
