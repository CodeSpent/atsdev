package main

import (
	"atsdev/internal/matcher"
	"atsdev/internal/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: atsdev <resume.pdf> <job_description.pdf>")
		return
	}

	resumeFile := os.Args[1]
	jobDescriptionFile := os.Args[2]

	resumeText, err := parser.ParsePDF(resumeFile)
	if err != nil {
		fmt.Printf("Error reading resume: %v\n", err)
		return
	}

	jobDescriptionText, err := parser.ParsePDF(jobDescriptionFile)
	if err != nil {
		fmt.Printf("Error reading job description: %v\n", err)
		return
	}

	resume := parser.ParseResume(resumeText)
	jobDescription := parser.ParseJobDescription(jobDescriptionText)
	score := matcher.Match(resume, jobDescription)

	fmt.Printf("Match Score: %.2f\n", score)
}
