/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/chrisevett/floridaman/cmd"
	"github.com/chrisevett/floridaman/pkg/database"
	i "github.com/chrisevett/floridaman/pkg/incident"
)

func main() {

	db, err := database.DbClient()

	if err != nil {
		fmt.Println("Failed to establish database connection: ", err)
		panic("Fatal")
	}

	// add static data

	// fuck it lets truncate for now

	db.Exec("TRUNCATE TABLE incident_states;")

	var incidentStates = []i.IncidentState{{Id: 1, State: "Investigating", Description: "We are currently investigating the issue"},
		{Id: 2, State: "Identified", Description: "We have identified the issue and are working on a resolution"},
		{Id: 3, State: "Monitoring", Description: "We have implemented a correction and are monitoring the outcome"},
		{Id: 4, State: "Resolved", Description: "The incident has been resolved"}}

	db.Create(&incidentStates)
	// todo : only do this if our entrypoint is cli
	cmd.Execute()
}
