package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "claim-documents",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetClaimDocuments",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = true
			crud["delete"] = true
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimDocuments", "GetClaimDocuments", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("service_id", "int", "in", 0, 0)
		cmd.NewParameter("document_source", "string", "in", 1, "")
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
