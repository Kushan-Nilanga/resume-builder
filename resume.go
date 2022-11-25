package main

type Header struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
	Website  string `yaml:"website"`
	Location string `yaml:"location"`
	Profile  string `yaml:"profile"`
}

// function to convert header struct to html string
func (h Header) ToHTML() string {
	var html string

	html = html + "<div class='header'>"
	html = html + "<h1>" + h.Name + "</h1>"
	html = html + "<ul>"
	html = html + "<li>" + h.Email + "</li>"
	if h.Phone != "" {
		html = html + "<li>" + h.Phone + "</li>"
	}
	html = html + "<li>" + h.Website + "</li>"
	html = html + "<li>" + h.Location + "</li>"
	html = html + "<li>" + h.Profile + "</li>"
	html = html + "</ul>"
	html = html + "</div>"

	return html
}

// struct to hold summary
type Summary struct {
	Summary string `yaml:"summary"`
}

// function to convert summary struct to html string
func (s Summary) ToHTML() string {
	var html string

	html = html + "<div class='summary'>"
	html = html + "<h2>Summary</h2>"
	html = html + "<hr/>"
	html = html + "<p>" + s.Summary + "</p>"
	html = html + "</div>"

	return html
}

// struct to hold skills
type Skills struct {
	Skills []string `yaml:"skills"`
}

// function to convert skills struct to html string
func (s Skills) ToHTML() string {
	var html string

	html = html + "<div class='skills'>"
	html = html + "<h2>Skills</h2>"
	html = html + "<hr/>"
	html = html + "<ul>"

	for _, skill := range s.Skills {
		html = html + "<li>" + skill + "</li>"
	}

	html = html + "</ul>"
	html = html + "</div>"

	return html
}

// struct to hold education
type Education struct {
	Education []struct {
		Institution string   `yaml:"institution"`
		Location    string   `yaml:"location"`
		Degree      string   `yaml:"degree"`
		Year        string   `yaml:"year"`
		Notes       []string `yaml:"notes"`
	} `yaml:"education"`
}

// function to convert education struct to html string
func (e Education) ToHTML() string {
	var html string

	html = html + "<div class='education'>"
	html = html + "<h2>Education</h2>"
	html = html + "<hr/>"

	for _, edu := range e.Education {
		html = html + "<div class='institution'>"
		html = html + "<ul class='heading'>"
		html = html + "<li class='primary'>" + edu.Institution + "</li>"
		html = html + "<li class='tertiary'>" + edu.Degree + "</li>"
		html = html + "<li class='secondary right'>" + edu.Location + "</li>"
		html = html + "<li class='tertiary right'>" + edu.Year + "</li>"
		html = html + "</ul>"

		html = html + "<ul>"

		for _, note := range edu.Notes {
			html = html + "<li>" + note + "</li>"
		}

		html = html + "</ul>"
		html = html + "</div>"
	}

	html = html + "</div>"

	return html
}

// struct to hold experience
type Experience struct {
	Experience []struct {
		Company  string   `yaml:"company"`
		Location string   `yaml:"location"`
		Position string   `yaml:"position"`
		Year     string   `yaml:"year"`
		Notes    []string `yaml:"notes"`
	} `yaml:"experience"`
}

// function to convert experience struct to html string
func (e Experience) ToHTML() string {
	var html string

	html = html + "<div class='experience'>"
	html = html + "<h2>Experience</h2>"
	html = html + "<hr/>"

	for _, exp := range e.Experience {
		html = html + "<div class='company'>"
		html = html + "<ul class='heading'>"
		html = html + "<li class='primary'>" + exp.Position + "</li>"
		html = html + "<li class='tertiary'>" + exp.Company + "</li>"
		html = html + "<li class='secondary right'>" + exp.Year + "</li>"
		html = html + "<li class='tertiary right'>" + exp.Location + "</li>"
		html = html + "</ul>"

		html = html + "<ul>"

		for _, note := range exp.Notes {
			html = html + "<li>" + note + "</li>"
		}

		html = html + "</ul>"
		html = html + "</div>"
	}

	html = html + "</div>"

	return html
}

// struct to hold certifications
type Certifications struct {
	Certifications []struct {
		Title      string   `yaml:"title"`
		Year       string   `yaml:"year"`
		Location   string   `yaml:"location"`
		Insitution string   `yaml:"institution"`
		Notes      []string `yaml:"notes"`
	} `yaml:"certifications"`
}

// function to convert certifications struct to html string
func (c Certifications) ToHTML() string {
	var html string

	html = html + "<div class='certifications'>"
	html = html + "<h2>Certifications</h2>"
	html = html + "<hr/>"

	for _, cert := range c.Certifications {
		html = html + "<div class='certification'>"
		html = html + "<ul class='heading'>"
		html = html + "<li class='primary'>" + cert.Title + "</li>"
		html = html + "<li class='tertiary'>" + cert.Insitution + "</li>"
		html = html + "<li class='secondary right'>" + cert.Year + "</li>"
		html = html + "<li class='tertiary right'>" + cert.Location + "</li>"
		html = html + "</ul>"

		html = html + "<ul>"

		for _, note := range cert.Notes {
			html = html + "<li>" + note + "</li>"
		}

		html = html + "</ul>"
		html = html + "</div>"
	}

	html = html + "</div>"

	return html
}

// struct to hold Projects
type Projects struct {
	Projects []struct {
		Title        string   `yaml:"title"`
		Year         string   `yaml:"year"`
		Location     string   `yaml:"location"`
		Technologies string   `yaml:"technologies"`
		Notes        []string `yaml:"notes"`
	} `yaml:"projects"`
}

// function to convert projects struct to html string
func (p Projects) ToHTML() string {
	var html string

	html = html + "<div class='projects'>"
	html = html + "<h2>Projects</h2>"
	html = html + "<hr/>"

	for _, proj := range p.Projects {
		html = html + "<div class='project'>"
		html = html + "<ul class='heading'>"
		html = html + "<li class='primary'>" + proj.Title + "</li>"
		html = html + "<li class='tertiary'>" + proj.Technologies + "</li>"
		html = html + "<li class='secondary right'>" + proj.Year + "</li>"
		html = html + "<li class='tertiary right'>" + proj.Location + "</li>"
		html = html + "</ul>"

		html = html + "<ul>"

		for _, note := range proj.Notes {
			html = html + "<li>" + note + "</li>"
		}

		html = html + "</ul>"
		html = html + "</div>"
	}

	html = html + "</div>"

	return html
}

// struct to hold References
type References struct {
	References []struct {
		Name     string `yaml:"name"`
		Position string `yaml:"position"`
		Company  string `yaml:"company"`
		Email    string `yaml:"email"`
		Phone    string `yaml:"phone"`
	} `yaml:"references"`
}

// function to convert references struct to html string
func (r References) ToHTML() string {
	var html string

	html = html + "<div class='references'>"
	html = html + "<h2>References</h2>"
	html = html + "<hr/>"

	for _, ref := range r.References {
		html = html + "<div class='reference'>"
		html = html + "<h3>" + ref.Name + "</h3>"
		html = html + "<h4>" + ref.Position + "</h4>"
		html = html + "<h4>" + ref.Company + "</h4>"
		html = html + "<h4>" + ref.Email + "</h4>"
		html = html + "<h4>" + ref.Phone + "</h4>"
		html = html + "</div>"
	}

	html = html + "</div>"

	return html
}

// Resume Data Structure
type Resume struct {
	Header         Header         `yaml:"header"`
	Summary        Summary        `yaml:"summary"`
	KeySkills      Skills         `yaml:"skills"`
	Education      Education      `yaml:"education"`
	Experience     Experience     `yaml:"experience"`
	Certifications Certifications `yaml:"certifications"`
	Projects       Projects       `yaml:"projects"`
	References     []string
}

// function to convert resume struct to html string
func (r Resume) ToHTML() string {
	var html string

	html = html + r.Header.ToHTML()
	html = html + r.Summary.ToHTML()
	html = html + r.KeySkills.ToHTML()
	html = html + r.Education.ToHTML()
	html = html + r.Experience.ToHTML()
	html = html + r.Certifications.ToHTML()
	html = html + r.Projects.ToHTML()

	return html
}
