/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/chrisevett/floridaman/pkg/database"
	i "github.com/chrisevett/floridaman/pkg/incident"
	"github.com/spf13/cobra"
)

// dbinitCmd represents the dbinit command
var dbinitCmd = &cobra.Command{
	Use:   "dbinit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbinit called")

		db, err := database.DbClient()
		// check erorr
		if err != nil {
			fmt.Println("Error ", err)
			panic("did on opsie db client init dbinitcmd")
		}

		db.Exec("TRUNCATE TABLE incident_states;")

		var incidentStates = []i.IncidentState{{Id: 1, State: "Investigating", Description: "We are currently investigating the issue"},
			{Id: 2, State: "Identified", Description: "We have identified the issue and are working on a resolution"},
			{Id: 3, State: "Monitoring", Description: "We have implemented a correction and are monitoring the outcome"},
			{Id: 4, State: "Resolved", Description: "The incident has been resolved"}}

		db.Create(&incidentStates)

		/*
				err = db.AutoMigrate(&i.AffectedSystem{}, &i.Severity{}, &i.Incident{}, &i.IncidentEvent{},
			&i.IncidentEventType{}, &i.IncidentState{}, &i.ChecklistItem{},
			&i.ChecklistTemplate{}, &i.Checklist{})
		*/
	},
}

func init() {
	rootCmd.AddCommand(dbinitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbinitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbinitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
