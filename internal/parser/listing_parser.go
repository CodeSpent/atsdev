package parser

import (
	"atsdev/internal/models"
	"strings"
)

func ParseJobDescription(jobDescriptionText string) models.JobListing {
	var jobListing models.JobListing

	lines := strings.Split(jobDescriptionText, "\n")
	if len(lines) > 0 {
		jobListing.Title = lines[0]
	}
	if len(lines) > 1 {
		jobListing.Company = lines[1]
	}

	jobListing.Skills = extractJobSkills(jobDescriptionText)
	jobListing.Duties = extractJobDuties(jobDescriptionText)
	jobListing.Requirements = extractJobRequirements(jobDescriptionText)

	return jobListing
}

func extractJobSkills(text string) []string {
	return []string{"Go", "Python"}
}

func extractJobDuties(text string) []string {
	return []string{}
}

func extractJobRequirements(text string) []string {
	return []string{}
}
