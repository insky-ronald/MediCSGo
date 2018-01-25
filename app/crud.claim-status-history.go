package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "claim-status-history",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetClaimStatusHistory",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimStatusHistory", "GetClaimStatusHistory", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
