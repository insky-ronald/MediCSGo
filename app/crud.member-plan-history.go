package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "member-plan-history",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetMemberPlanHistory",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetMemberPlanHistory", "GetMemberPlanHistory", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("member_id", "int", "in", 0, 1)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
