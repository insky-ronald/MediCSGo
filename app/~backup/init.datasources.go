package main

import (
	"ibsi/dbase"
)

func init() {
	//***************************************************************************************************
	//	SYSTEM 
	//***************************************************************************************************
	dbase.Connections["DBSecure"].NewCommand("AddVisit", "System_AddVisit", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("visit_id", "int", "inout", 0, 0)
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("session_id", "string", "in", 48, "")
		cmd.NewParameter("method", "string", "in", 10, "")
		cmd.NewParameter("local_ip", "string", "in", 20, "")
		cmd.NewParameter("remote_ip", "string", "in", 20, "")
		cmd.NewParameter("remote_host", "string", "in", 100, "")
		cmd.NewParameter("user_agent", "string", "in", 100, "")
		cmd.NewParameter("referrer_url", "string", "in", 200, "")
		cmd.NewParameter("request_url", "string", "in", 200, "")
	})

	dbase.Connections["DBSecure"].NewCommand("Login", "System_Login", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("user_id", "int", "inout", 0, -20)
		cmd.NewParameter("user_name", "string", "in", 200, "")
		cmd.NewParameter("password", "string", "in", 200, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBSecure"].NewCommand("Logout", "System_Logout", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	//====================================================================================================
	//	USERS
	//====================================================================================================
	dbase.Connections["DBSecure"].NewCommand("AddUser", "AddUser", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("organisation_id", "int", "in", 0, 0)
		cmd.NewParameter("user_name", "string", "in", 200, "")
		cmd.NewParameter("last_name", "string", "in", 60, "")
		cmd.NewParameter("middle_name", "string", "in", 60, "")
		cmd.NewParameter("first_name", "string", "in", 60, "")
		cmd.NewParameter("gender", "string", "in", 1, "")
		cmd.NewParameter("dob", "datetime", "in", 0, nil)
		cmd.NewParameter("email", "string", "in", 200, "")
		cmd.NewParameter("status_code_id", "int", "in", 0, 0)
		cmd.NewParameter("roles", "string", "in", 100, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetUsers", "GetUsers", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("ids", "string", "in", 200, "")
		cmd.NewParameter("filter", "string", "in", 200, "")
		cmd.NewParameter("mode", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "user_name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetUserSessionInfo", "System_GetUserSessionInfo", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("visit_id", "int", "in", 0, nil)
	})

	dbase.Connections["DBSecure"].NewCommand("GetMyRights", "System_GetMyRights", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("action_id", "int", "in", 0, nil)
		cmd.NewParameter("visit_id", "int", "in", 0, nil)
		cmd.NewParameter("error_log_id", "int", "in", 0, 0)
		cmd.NewParameter("verbose", "int", "in", 0, 1)
	})

	dbase.Connections["DBSecure"].NewCommand("GetUserDetails", "System_GetUserDetails", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
	})

	dbase.Connections["DBSecure"].NewCommand("GetUserActions", "System_GetUserActions", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("user_id", "int", "in", 0, 0)
		cmd.NewParameter("role_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, nil)
		cmd.NewParameter("error_log_id", "int", "inout", 0, 0)
		cmd.NewParameter("verbose", "int", "in", 0, 1)
	})

	dbase.Connections["DBSecure"].NewCommand("lookup_users", "GetUsersLookup", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("ids", "string", "in", 100, "")
		cmd.NewParameter("filter", "string", "in", 200, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "full_name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
	}) 

	//====================================================================================================
	//ROLES
	//====================================================================================================
	dbase.Connections["DBSecure"].NewCommand("AddRole", "AddRole", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("role", "string", "in", 100, "")
		cmd.NewParameter("description", "string", "in", 200, "")
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("position", "int", "in", 0, 0)
		cmd.NewParameter("status_code_id", "int", "in", 0, 0)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetRoles", "GetRoles", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("ids", "string", "in", 200, "")
		cmd.NewParameter("filter", "string", "in", 200, "")
		cmd.NewParameter("mode", "int", "in", 0, 0)
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "position")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetUserRoles", "System_GetUserRoles", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("user_id", "int", "in", 0, 0)
		cmd.NewParameter("error_log_id", "int", "inout", 0, 0)
		cmd.NewParameter("verbose", "int", "in", 0, 1)
	})

	//====================================================================================================
	//	 ACTIONS
	//====================================================================================================
	dbase.Connections["DBSecure"].NewCommand("AddAction", "AddAction", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("code", "string", "in", 20, "")
		cmd.NewParameter("action_name", "string", "in", 100, "")
		cmd.NewParameter("description", "string", "in", 200, "")
		cmd.NewParameter("action_type_id", "int", "in", 0, 0)
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("position", "int", "in", 0, 0)
		cmd.NewParameter("status_code_id", "int", "in", 0, 0)
		cmd.NewParameter("rights", "string", "in", 100, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetActions", "GetActions", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("ids", "string", "in", 200, "")
		cmd.NewParameter("filter", "string", "in", 200, "")
		cmd.NewParameter("mode", "int", "in", 0, 0)
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "position")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	//====================================================================================================
	//	RIGHTS
	//====================================================================================================
	dbase.Connections["DBSecure"].NewCommand("AddRights", "AddRights", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("code", "string", "in", 20, "")
		cmd.NewParameter("rights", "string", "in", 100, "")
		cmd.NewParameter("description", "string", "in", 200, "")
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("position", "int", "in", 0, 0)
		cmd.NewParameter("status_code_id", "int", "in", 0, 0)

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetRights", "GetRights", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("ids", "string", "in", 200, "")
		cmd.NewParameter("filter", "string", "in", 200, "")
		cmd.NewParameter("mode", "int", "in", 0, 0)
		cmd.NewParameter("application_id", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "position")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetActionRights", "GetActionRights", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("action_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	//====================================================================================================
	//	PERMISSIONS
	//====================================================================================================
	dbase.Connections["DBSecure"].NewCommand("AddPermission", "AddPermission", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("role_id", "int", "in", 0, 0)
		cmd.NewParameter("action_id", "int", "in", 0, 0)
		cmd.NewParameter("permissions", "string", "in", 100, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetManagePermissions", "GetManagePermissions", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("role_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetPermissions", "GetPermissions", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("role_id", "int", "in", 0, 0)
		cmd.NewParameter("action_id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetAllowAction", "GetAllowAction", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("action_code", "string", "in", 20, "")
		cmd.NewParameter("allow", "bool", "inout", 1, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBSecure"].NewCommand("GetMyPermission", "GetMyPermission", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("action_code", "string", "in", 20, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	//====================================================================================================
	// 
	//====================================================================================================
	dbase.Connections["DBSecure"].NewCommand("GetText", "GetText", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBSecure"].NewCommand("AddText", "AddText", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("label", "string", "in", 100, "")
		cmd.NewParameter("text", "string", "in", -1, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})

	//====================================================================================================
	// 
	//====================================================================================================
	dbase.Connections["DBApp"].NewCommand("GetOwner", "GetOwner", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
