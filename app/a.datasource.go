package main

import (
	"ibsi/dbase"
)

func init() {
	dbase.Connections["DBApp"].NewCommand("System_ManageSession", "System_ManageSession", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("session_id", "string", "in", 50, "")
		cmd.NewParameter("user_id", "int", "in", 0, 0)
		cmd.NewParameter("user_name", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
	})
	
	dbase.Connections["DBApp"].NewCommand("GetCustomerServiceData", "GetCustomerServiceData", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
	
	dbase.Connections["DBApp"].NewCommand("GetClaimMemberInfo", "GetClaimMemberInfo", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("member_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetMemberMedicalNotes", "GetMemberMedicalNotes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 1)
		cmd.NewParameter("claim_id", "int", "in", 0, 1)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetService", "GetService", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("service_type", "string", "in", 3, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetServiceSubType", "GetServiceSubType", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("service_type", "string", "in", 3, "")
		cmd.NewParameter("sub_type", "string", "in", 4, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetGopCalculationDates", "GetGopCalculationDates", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetGopEstimates", "GetGopEstimates", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetServiceSubTypes", "GetServiceSubTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("service_type", "string", "in", 3, "")
		cmd.NewParameter("sub_type", "string", "in", 4, "")
	})
	
	// dbase.Connections["DBApp"].NewQuery("lookup_service_type_name", func(cmd dbase.ICommand) {
			// cmd.NewParameter("module", "string", "in", 3, "")
			// cmd.NewParameter("code", "string", "in", 4, "")
		// }, 
		// func() string {
			// return `
				// select
					// code,
					// sub_type,
					// display_name
				// from service_sub_types
				// where service_type = @module and code = @code
			// `
		// },
	// )
	
	dbase.Connections["DBApp"].NewCommand("GetMemberClaims", "GetMemberClaims", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("member_id", "int", "in", 0, 0)
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "claim_no")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})	
	
	dbase.Connections["DBApp"].NewCommand("GetClaimServices", "GetClaimServices", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("service_type", "string", "in", 3, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
	
	dbase.Connections["DBApp"].NewCommand("lookup_claim_types", "GetClaimTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("lookup", "int", "in", 0, 1)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
