package matcher

import (
	"atsdev/internal/models"
	"atsdev/pkg/keywords"
	"fmt"
	"strings"
)

type MatchMetrics struct {
	SingleMatch     int
	MultipleMatch   int
	TotalMatched    int
	UnmatchedSkills []string
}

func Match(resume models.Resume, jobListing models.JobListing) float64 {
	hardSkillMetrics := calculateMatchMetrics(resume.Skills, jobListing.Skills, keywords.HardSkills)
	softSkillMetrics := calculateMatchMetrics(resume.Skills, jobListing.Skills, keywords.SoftSkills)

	logMetrics("Hard Skills", hardSkillMetrics)
	logMetrics("Soft Skills", softSkillMetrics)

	hardSkillScore := calculateScore(hardSkillMetrics)
	softSkillScore := calculateScore(softSkillMetrics)

	totalScore := (hardSkillScore * 0.7) + (softSkillScore * 0.3)
	return totalScore
}

func calculateMatchMetrics(resumeSkills, jobSkills, predefinedSkills []string) MatchMetrics {
	var metrics MatchMetrics
	skillCount := make(map[string]int)

	for _, skill := range predefinedSkills {
		if containsIgnoreCase(resumeSkills, skill) || containsIgnoreCase(jobSkills, skill) {
			skillCount[strings.ToLower(skill)]++
		} else {
			metrics.UnmatchedSkills = append(metrics.UnmatchedSkills, skill)
		}
	}

	for _, count := range skillCount {
		if count == 1 {
			metrics.SingleMatch++
		} else if count > 4 {
			metrics.MultipleMatch++
		}
		metrics.TotalMatched += count
	}

	return metrics
}

func calculateScore(metrics MatchMetrics) float64 {
	if metrics.TotalMatched == 0 {
		return 0
	}
	return float64(metrics.SingleMatch+metrics.MultipleMatch) / float64(metrics.TotalMatched)
}

func containsIgnoreCase(skills []string, skill string) bool {
	for _, s := range skills {
		if strings.EqualFold(s, skill) {
			return true
		}
	}
	return false
}

func logMetrics(skillType string, metrics MatchMetrics) {
	fmt.Printf("Metrics for %s:\n", skillType)
	fmt.Printf("Single Match: %d\n", metrics.SingleMatch)
	fmt.Printf("Multiple Match: %d\n", metrics.MultipleMatch)
	fmt.Printf("Total Matched: %d\n", metrics.TotalMatched)
	fmt.Printf("Unmatched Skills: %v\n", metrics.UnmatchedSkills)
	fmt.Println()
}
