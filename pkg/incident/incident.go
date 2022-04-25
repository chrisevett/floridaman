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
	IncidentStateId int
}

// todo : checklist populate type Checklist struct

type IncidentState struct {
	gorm.Model
	Id          int `gorm:"uniqueIndex;autoIncrement:false"`
	State       string
	Description string
}

/*
- has
    - ID
    - start time
    - end time
    - events
    - tags
    - many systems
    - kb link
    - state
    - severity

*/

// createincident will return an incident id and later a checklist
func CreateIncident(db *gorm.DB) (Incident, error) {

	var incidentStates = Incident{Description: "I cant see my toes", StartTime: datatypes.Date(time.Now())}

	db.Create(&incidentStates)

	return Incident{}, nil
}

// just return all the data we have I guess
func ResolveIncident(db *gorm.DB, incidentId int) (Incident, error) {

	return Incident{}, nil
}
