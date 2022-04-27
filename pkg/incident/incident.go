package incident

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Incident struct {
	gorm.Model
	Id              int `gorm:"primaryKey"`
	Description     string
	StartTime       datatypes.Date
	EndTime         datatypes.Date
	Tags            []string
	Events          []IncidentEvent // array of ids? better way to represent this so its time series, i bet thats a postgres feature
	IncidentStateId int
	IncidentState   string
	SeverityId      int
	Severity        string
	System          AffectedSystem
	Checklists      []Checklist
}

// todo : checklist populate type Checklist struct

type IncidentState struct {
	gorm.Model
	Id          int `gorm:"uniqueIndex;autoIncrement:false"`
	State       string
	Description string
}

type Severity struct {
	gorm.Model
	Id          int `gorm:"uniqueIndex;autoIncrement:false"`
	Severity    string
	Description string
}

type AffectedSystem struct {
	gorm.Model
	Id          int `gorm:"uniqueIndex;autoIncrement:false"`
	System      string
	Description string
}

type IncidentEvent struct {
	gorm.Model
	Id            int
	OccuranceDate datatypes.Date
	EventType     IncidentEventType

	// todo : support a binary like an image for this
	// this should be a pointer to object storage

}

// we'll need a separate data type for page, mention, image, huddle (do we record those)
type IncidentEventType struct {
	gorm.Model
	Id          int `gorm:"uniqueIndex;autoIncrement:false"`
	EventType   string
	Description string
}

type Checklist struct {
	gorm.Model
	Id       int
	Incident Incident          // todo make sure this is A FK
	Template ChecklistTemplate // todo make sure this is also a FK

}

type ChecklistTemplate struct {
	gorm.Model
	Id int
}

type ChecklistItem struct {
	gorm.Model
	Id          int
	Description string
	State       bool
}

// createincident will return an incident id and later a checklist
func CreateIncident(db *gorm.DB) (Incident, error) {

	var incidentStates = Incident{Description: "I cant see my toes", StartTime: datatypes.Date(time.Now())}

	db.Create(&incidentStates)

	return Incident{}, nil
}

// we should get some data from slack about which incident this is, like from the channel
func UpdateIncident(db *gorm.DB, id int, desiredState string) (Incident, error) {

	// transition state to desired state

	// err if not valid

	// get checklist associated with state

	// return incident with new status and checklist

}

// just return all the data we have I guess
func ResolveIncident(db *gorm.DB, incidentId int) (Incident, error) {

	return Incident{}, nil
}
