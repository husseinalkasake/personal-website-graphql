package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
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

// "jwL8X;{sbJ~MB,^P"

// ConnString returns a connection string to database
func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, "HailToTheThief@20XX")
}

// Project struct
type Project struct {
	ID             int
	Title          string
	Subtitle       string
	WorkInProgress bool
	Date           string
	Summary        []string
	MoreInfo       []string
	Source         string
	Videos         []string
	Images         []string
}

// GetAllProjects is called within our projects query
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
			&project.ID,
			&project.Title,
			&project.Subtitle,
			&project.WorkInProgress,
			&project.Date,
			&project.Summary,
			&project.MoreInfo,
			&project.Source,
			&project.Videos,
			&project.Images,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		fmt.Println("PROJECT", project)
		projects = append(projects, project)
	}

	return projects

}
