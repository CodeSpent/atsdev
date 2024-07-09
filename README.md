# ATSDev - Resume to Job Description Matching Tool

ATSDev is a Go-based application designed to match a resume against a job description and calculate a match score based on hard and soft skills. This tool is tailored towards developers, considering both technical skills and soft skills critical for the role.

## Features
- Parse resumes and job descriptions from PDF files.
- Extract text using OCR (Optical Character Recognition).
- Match skills from the resume with the job description.
- Calculate a match score based on predefined hard and soft skills.
- Output detailed metrics for analysis.

## Installation

### Prerequisites

- Go (1.16 or later)
- ImageMagick
- Tesseract OCR

### Setup

1. **Clone the Repository**

   ```shell
   git clone https://github.com/yourusername/atsdev.git
   cd atsdev
   ```
2. Install Dependencies 
Make sure you have ImageMagick and Tesseract OCR installed on your system.

    For macOS:
    ```shell
    brew install imagemagick
    brew install tesseract 
    ```

    For Ubuntu:

    ```shell
    sudo apt-get install imagemagick
    sudo apt-get install tesseract-ocr
    ```
   
## Usage

### Running the Application
```shell
./atsdev <resume.pdf> <job_description>
```



- <resume.pdf>: Path to the resume PDF file.
- <job_description>: String of text from a Job Listing.

**Example:**
```shell
./atsdev resume.pdf "Looking for a 10x rockstar Go engineer under 30 years old with 55 years of experience in Rust"
```

## Contributing

Always welcoming contributions from the community. Please bare with me as I work further on this then open up for contribution!

In the meantime, of course, feel free to open an issue if you see anything wrong or have a feature request.