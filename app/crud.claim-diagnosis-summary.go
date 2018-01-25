package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "claim-diagnosis-summary",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetClaimDiagnosisSummary",
		// Rights: map[string]bool{
			// "add": true,
			// "edit": true,
			// "delete": false,
		// },
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = true
			crud["edit"] = true
			crud["delete"] = false
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimDiagnosisSummary", "GetClaimDiagnosisSummary", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
