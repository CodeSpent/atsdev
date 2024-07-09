package matcher

import "atsdev/internal/models"

func Match(resume models.Resume, jobDescription models.JobListing) float64 {
	var score float64
	skillMatchCount := 0

	for _, skill := range jobDescription.Skills {
		for _, resumeSkill := range resume.Skills {
			if skill == resumeSkill {
				skillMatchCount++
			}
		}
	}

	if len(jobDescription.Skills) > 0 {
		score = float64(skillMatchCount) / float64(len(jobDescription.Skills))
	}

	return score
}
