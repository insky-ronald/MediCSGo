package main

import (
	"ibsi/dbase"

func init() {
	dbase.Connections["DBApp"].NewCommand("Sys_GetTables", "Sys_GetTables", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetUser", "GetUser", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetUsers", "GetUsers", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddUser", "AddUser", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("user_name", "string", "in", 10, "")
		cmd.NewParameter("name", "string", "in", 60, "")
		cmd.NewParameter("designation", "string", "in", 60, "")
		cmd.NewParameter("phone_no", "string", "in", 20, "")
		cmd.NewParameter("is_active", "bool", "in", 0, 0)
		cmd.NewParameter("is_supervisor", "bool", "in", 0, 0)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetGroup", "GetGroup", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetGroups", "GetGroups", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetUserGroups", "GetUserGroups", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("user_name", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimsEnquiry", "GetClaimsEnquiry", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClaim", "GetClaim", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetMember", "GetMember", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetTaskManager", "GetTaskManager", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetPlan", "GetPlan", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetProduct", "GetProduct", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClient", "GetClient", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetContacts", "GetContacts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetBanks", "GetBanks", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddContacts", "AddContacts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("title", "string", "in", 20, "")
		cmd.NewParameter("full_name", "string", "in", 50, "")
		cmd.NewParameter("department", "string", "in", 50, "")
		cmd.NewParameter("position", "string", "in", 50, "")
		cmd.NewParameter("phone", "string", "in", 20, "")
		cmd.NewParameter("mobile", "string", "in", 20, "")
		cmd.NewParameter("fax", "string", "in", 20, "")
		cmd.NewParameter("email", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("AddBanks", "AddBanks", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("bank_name", "string", "in", 34, "")
		cmd.NewParameter("sort_code", "string", "in", 11, "")
		cmd.NewParameter("swift_code", "string", "in", 11, "")
		cmd.NewParameter("bank_address1", "string", "in", 34, "")
		cmd.NewParameter("bank_address2", "string", "in", 34, "")
		cmd.NewParameter("bank_address3", "string", "in", 34, "")
		cmd.NewParameter("bank_country_code", "string", "in", 2, "")
		cmd.NewParameter("beneficiary_name", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_bank_account", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_address1", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_address2", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_address3", "string", "in", 34, "")
		cmd.NewParameter("beneficiary_country_code", "string", "in", 2, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetCurrencies", "GetCurrencies", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "currency")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddCurrencies", "AddCurrencies", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 3, "")
		cmd.NewParameter("currency", "string", "in", 60, "")
		cmd.NewParameter("is_active", "int", "in", 0, 0)
		cmd.NewParameter("no_roundoff", "int", "in", 0, 0)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetNationalities", "GetNationalities", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 10, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddNationalities", "AddNationalities", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 10, "")
		cmd.NewParameter("nationality", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimDiagnosis", "GetClaimDiagnosis", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimDiagnosisEdit", "GetClaimDiagnosisEdit", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GeClaimNotes", "GeClaimNotes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("service_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetHospitals", "GetHospitals", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClinics", "GetClinics", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetPharmacies", "GetPharmacies", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetProviderDiscount", "GetProviderDiscount", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 1)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClients", "GetClients", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetProducts", "GetProducts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 10, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "product_name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetPolicies", "GetPolicies", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "policy_no")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetClaimTypes", "GetClaimTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("lookup_claim_types", "GetClaimTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("lookup", "int", "in", 0, 1)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddClaimTypes", "AddClaimTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("claim_type", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetInvoiceTypes", "GetInvoiceTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddInvoiceTypes", "AddInvoiceTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetCaseFeeTypes", "GetCaseFeeTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddCaseFeeTypes", "AddCaseFeeTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetCustomerServiceTypes", "GetCustomerServiceTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddCustomerServiceTypes", "AddCustomerServiceTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetGOPTypes", "GetGOPTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddGOPTypes", "AddGOPTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetNOCTypes", "GetNOCTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddNOCTypes", "AddNOCTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetRecoveryTypes", "GetRecoveryTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddRecoveryTypes", "AddRecoveryTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetCostContainmentTypes", "GetCostContainmentTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddCostContainmentTypes", "AddCostContainmentTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetFlagTypes", "GetFlagTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 4, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddFlagTypes", "AddFlagTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 4, "")
		cmd.NewParameter("service_description", "string", "in", 60, "")
		cmd.NewParameter("display_name", "string", "in", 100, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetFlagSubTypes", "GetFlagSubTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddFlagSubTypes", "AddFlagSubTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 3, "")
		cmd.NewParameter("flag_code", "string", "in", 4, "")
		cmd.NewParameter("flag_sub_type", "string", "in", 50, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetServiceStatusCodes", "GetServiceStatusCodes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 7, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "module")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	// DBConnection.NewCommand("ClaimNotesTypes", "
		// select *
		// from notes
		// order by id
	// ", CommandType.Text)

	 // dbase.Connections["DBApp"].NewCommand("GetActionTypes", "GetActionTypes", "procedure", func(cmd dbase.ICommand) {
		// cmd.NewParameter("code", "string", "in", 3, "")
	 // })

	 dbase.Connections["DBApp"].NewCommand("GetAuditlogTypes", "GetAuditlogTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddAuditlogTypes", "AddAuditlogTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 3, "")
		cmd.NewParameter("description", "string", "in", 100, "")
		cmd.NewParameter("log_type", "string", "in", 1, "")
		cmd.NewParameter("is_active", "int", "in", 0, 1)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("AddNoteType", "AddNoteType", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("note_type", "string", "in", 60, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetNoteTypes", "GetNoteTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("filter", "string", "in", 100, "")

		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "country")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddNoteSubType", "AddNoteSubType", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("note_type", "string", "in", 3, "")
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("note_sub_type", "string", "in", 60, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	dbase.Connections["DBApp"].NewCommand("GetNoteSubTypes", "GetNoteSubTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("note_type", "string", "in", 3, "")
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("filter", "string", "in", 100, "")

		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "country")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetCustomerServiceData", "GetCustomerServiceData", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
