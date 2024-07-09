package parser

import (
	"atsdev/internal/models"
	"regexp"
)

func ParseResume(resumeText string) models.Resume {
	var resume models.Resume

	resume.Name = extractName(resumeText)
	resume.Email = extractEmail(resumeText)
	resume.Phone = extractPhone(resumeText)
	resume.Skills = extractSkills(resumeText)
	resume.Experience = extractExperience(resumeText)
	resume.Education = extractEducation(resumeText)

	resume.Content = resumeText

	return resume
}

func extractName(text string) string {
	return "Isaac Newton"
}

func extractEmail(text string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	return re.FindString(text)
}

func extractPhone(text string) string {
	re := regexp.MustCompile(`$begin:math:text$?\\d{3}$end:math:text$?[-.\s]?\d{3}[-.\s]?\d{4}`)
	return re.FindString(text)
}

func extractSkills(text string) []string {
	return []string{"Go", "Python"}
}

func extractExperience(text string) []models.Experience {
	return []models.Experience{}
}

func extractEducation(text string) []models.Education {
	return []models.Education{}
}
