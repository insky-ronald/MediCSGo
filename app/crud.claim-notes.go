package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "claim-notes",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetClaimNotes",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = true
			crud["edit"] = true
			crud["delete"] = true
			// DatabaseUtils.GetActionPermission("admin", Crud)
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimNotes", "GetClaimNotes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("type", "string", "in", 1, "")
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("service_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 100, "")
		cmd.NewParameter("order", "string", "in", 100, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
