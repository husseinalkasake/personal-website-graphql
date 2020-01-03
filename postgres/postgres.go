package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	"github.com/lib/pq"
)

// Db is the database struct
type Db struct {
	*sql.DB
}

// New creates a new database using a connection string
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string to database
func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, "Toronto1")
}

// About struct
type About struct {
	ID        int
	Header    string
	Paragraph []string
}

// Education struct
type Education struct {
	ID               int
	Title            string
	Institution      string
	Location         string
	Date             string
	RelevantProjects []string
	RelevantCourses  []string
}

// Experience struct
type Experience struct {
	ID           int
	Title        string
	Organization string
	Location     string
	DateRange    string
	Summary      []string
}

// Project struct
type Project struct {
	ID             int
	Title          string
	Subtitle       sql.NullString
	WorkInProgress bool
	Date           string
	Summary        []string
	MoreInfo       []string
	Source         sql.NullString
	Videos         []string
	Images         []string
}

// GetAbout gets the about section
func (d *Db) GetAbout() []About {
	// Prepare query
	stmt, err := d.Prepare("SELECT * FROM about")
	if err != nil {
		fmt.Println("GetAbout Preperation Err: ", err)
	}

	// Make query
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("GetAbout Query Err: ", err)
	}

	// Receive response
	var about About
	abouts := []About{}
	for rows.Next() {
		err = rows.Scan(
			&about.ID,
			&about.Header,
			pq.Array(&about.Paragraph),
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		abouts = append(abouts, about)
	}

	return abouts
}

// GetEducation gets all education experience
func (d *Db) GetEducation() []Education {
	// Prepare query
	stmt, err := d.Prepare("SELECT * FROM education")
	if err != nil {
		fmt.Println("GetEducation Preperation Err: ", err)
	}

	// Make query
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("GetEducation Query Err: ", err)
	}

	// Receive response
	var education Education
	educations := []Education{}
	for rows.Next() {
		err = rows.Scan(
			&education.ID,
			&education.Title,
			&education.Institution,
			&education.Location,
			&education.Date,
			pq.Array(&education.RelevantProjects),
			pq.Array(&education.RelevantCourses),
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		educations = append(educations, education)
	}

	return educations
}

// GetExperience gets all work experience
func (d *Db) GetExperience() []Experience {
	// Prepare query
	stmt, err := d.Prepare("SELECT * FROM experience")
	if err != nil {
		fmt.Println("GetExperience Preperation Err: ", err)
	}

	// Make query
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("GetExperience Query Err: ", err)
	}

	// Receive response
	var experience Experience
	experiences := []Experience{}
	for rows.Next() {
		err = rows.Scan(
			&experience.Title,
			&experience.Organization,
			&experience.Location,
			&experience.ID,
			&experience.DateRange,
			pq.Array(&experience.Summary),
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		experiences = append(experiences, experience)
	}

	return experiences
}

// GetAllProjects gets all projects
func (d *Db) GetAllProjects() []Project {
	// Prepare query
	stmt, err := d.Prepare("SELECT * FROM projects")
	if err != nil {
		fmt.Println("GetAllProjects Preperation Err: ", err)
	}

	// Make query
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("GetAllProjects Query Err: ", err)
	}

	// Receive response of projects
	var project Project
	projects := []Project{}
	for rows.Next() {
		err = rows.Scan(
			&project.Title,
			&project.ID,
			&project.Subtitle,
			&project.WorkInProgress,
			&project.Date,
			pq.Array(&project.Summary),
			pq.Array(&project.MoreInfo),
			&project.Source,
			pq.Array(&project.Videos),
			pq.Array(&project.Images),
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		projects = append(projects, project)
	}

	return projects
}
