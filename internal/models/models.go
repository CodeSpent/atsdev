package models

// Resume Structs
type Experience struct {
	JobTitle    string
	CompanyName string
	StartDate   string
	EndDate     string
	Description string
}

type Education struct {
	Degree      string
	Institution string
	StartDate   string
	EndDate     string
	Description string
}

type Resume struct {
	Name       string
	Email      string
	Phone      string
	Skills     []string
	Experience []Experience
	Education  []Education
	Content    string
}

// Listing Structs
type JobListing struct {
	Title        string
	Company      string
	Skills       []string
	Duties       []string
	Requirements []string
}
