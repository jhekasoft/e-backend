package models

type (
	DevTimelineItem struct {
		Year uint     `json:"year" yaml:"year"`
		Desc []string `json:"desc" yaml:"desc"`
	}

	CVCommon struct {
		Name        string `json:"name" yaml:"name"`
		YearOfBirth uint   `json:"yearOfBirth" yaml:"yearOfBirth"`
		GithubURL   string `json:"githubUrl" yaml:"githubUrl"`
		LinkedInURL string `json:"linkedInUrl" yaml:"linkedInUrl"`
		Telegram    string `json:"telegram" yaml:"telegram"`
		Email       string `json:"email" yaml:"email"`
	}

	CVEducationItem struct {
		StartYear    uint    `json:"startYear" yaml:"startYear"`
		EndYear      uint    `json:"endYear" yaml:"endYear"`
		School       string  `json:"school" yaml:"school"`
		Degree       string  `json:"degree" yaml:"degree"`
		FieldOfStudy string  `json:"fieldOfStudy" yaml:"fieldOfStudy"`
		Location     string  `json:"location" yaml:"location"`
		Desc         *string `json:"desc,omitempty" yaml:"desc,omitempty"`
	}

	CVExperienceItem struct {
		StartYear  uint    `json:"startYear" yaml:"startYear"`
		EndYear    uint    `json:"endYear" yaml:"endYear"`
		Title      string  `json:"title" yaml:"title"`
		Company    string  `json:"company" yaml:"company"`
		CompanyURL *string `json:"companyUrl,omitempty" yaml:"companyUrl,omitempty"`
		Location   string  `json:"location" yaml:"location"`
		Desc       *string `json:"desc,omitempty" yaml:"desc,omitempty"`
	}

	CVPublication struct {
		Year     uint   `json:"year" yaml:"year"`
		Title    string `json:"title" yaml:"title"`
		URL      string `json:"url" yaml:"url"`
		Desc     string `json:"desc" yaml:"desc"`
		ImageURL string `json:"imageUrl" yaml:"imageUrl"`
	}

	CVSoftwareProject struct {
		Year     uint   `json:"year" yaml:"year"`
		Title    string `json:"title" yaml:"title"`
		URL      string `json:"url" yaml:"url"`
		Desc     string `json:"desc" yaml:"desc"`
		ImageURL string `json:"imageUrl" yaml:"imageUrl"`
	}
)
