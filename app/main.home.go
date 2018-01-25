package main

import (
	"net/http"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"ibsi/dbase"
	"ibsi/system"
)

type PageData struct {
	CustomerService dbase.TDataTableRows `json:"data"`
}

// AddTemplateController is in ibsi.controller.template.go
func init() {
	template.NewController(template.Controller {
		Pid: "home",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/{home:home\\/?}")
			ts.Add("\\/?")
			// ts.Add("/app/home")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			
			cs := dbase.Connections["DBApp"].OpenDataTable("GetCustomerServiceData", dbase.TParameters{"visit_id": vid})
			
			p.Title = "Claim1Med"
			p.Nav.WindowTitle = "Claim1Med"
			
			// p.Nav.CustomData = PageData{
				// CustomerService: cs.GetRows(),
			// }
			p.Nav.CustomData = map[string]interface{}{
				"data": cs.GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "main", "Main", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "tasks"
					s.Action = "admin"
					s.Title = "Task Manager"
					s.Icon = "timetable"
					s.Url = "app/task-manager"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "members"
					s.Action = "admin"
					s.Title = "Claims Processing"
					s.Icon = "users"
					s.Url = "app/claims-entry"
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "enquiries"
					s.Action = "admin"
					s.Title = "Enquiries"
					s.Icon = "search"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "reports"
					s.Action = "admin"
					s.Title = "Reports"
					s.Icon = "table"
					s.Url = "engine/sys-reports"
				})
			})
			
			if system.AllowAction("security") {
				utils.NewNavigatorItem(p.Nav, "security", "Security", func(item *utils.NavigatorItem) {
				
					utils.NewMenuItem(item, func(s *utils.MenuItem) {
						s.ID = "users"
						s.Action = "sys-users"
						s.Title = "Users"
						s.Icon = "user"
						s.Url = "engine/sys-users"
					});
					
					utils.NewMenuItem(item, func(s *utils.MenuItem) {
						s.ID = "roles"
						s.Action = "sys-roles"
						s.Title = "Roles"
						s.Icon = "users"
						s.Url = "engine/sys-roles"
					});

					utils.NewMenuItem(item, func(s *utils.MenuItem) {
						s.ID = "actions"
						s.Action = "sys-actions"
						s.Title = "Actions"
						s.Icon = "star"
						s.Url = "engine/sys-actions"
						s.Params["app_only"] = 1
					});

					utils.NewMenuItem(item, func(s *utils.MenuItem) {
						s.ID = "rights";
						s.Action = "sys-rights";
						s.Title = "Rights";
						s.Icon = "settings";
						s.Url = "engine/sys-rights";
						s.Params["app_only"] = 1
					});

					// utils.NewMenuItem(item, func(s *utils.MenuItem) {
						// s.ID = "permissions";
						// s.Action = "sys-permissions";
						// s.Title = "Set Permissions";
						// s.Icon = "security";
						// s.Url = "engine/sys-roles";
					// });
					
				})
			}
		},
	})
}
