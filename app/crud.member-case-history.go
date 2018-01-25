package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "member-case-history",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetCaseHistory",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetCaseHistory", "GetCaseHistory", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("member_id", "int", "in", 0, 0)
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "claim_no")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
