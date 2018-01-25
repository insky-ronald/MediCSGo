package main

import (
	"net/http"
	"ibsi/utils"
	"ibsi/template"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	template.NewController(template.Controller {
		Pid: "home",
		Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/{home:home\\/?}")
			// ts.Add("/app/home")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			p.Title = "GO! Claims1MED"
			p.Nav.WindowTitle = "Claims1MED"
			
			utils.NewNavigatorItem(p.Nav, "main", "Main", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "customers"
					s.Action = "customers"
					s.Title = "Customers"
					s.Icon = "users"
					s.Css = "*"
					s.Scripts = "*"
					// s.Url = "app/customers"
				})
				
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "suppliers"
					// s.Action = "suppliers"
					// s.Title = "Suppliers"
					// s.Icon = "users"
					// s.Url = "app/suppliers"
				// })
			})
			
			// if (Security.AllowAction("security")) {
				// utils.NewNavigatorItem(p.Nav, "security", "Security", func(item *utils.NavigatorItem) {
				
					// utils.NewMenuItem(item, func(s *utils.MenuItem) {
						// s.ID = "users"
						// s.Action = "sys-users"
						// s.Title = "Users"
						// s.Icon = "user"
						// s.Url = "engine/sys-users"
					// });
					
					// utils.NewMenuItem(item, func(s *utils.MenuItem) {
						// s.ID = "roles"
						// s.Action = "sys-roles"
						// s.Title = "Roles"
						// s.Icon = "users"
						// s.Url = "engine/sys-roles"
					// });

					// utils.NewMenuItem(item, func(s *utils.MenuItem) {
						// s.ID = "actions"
						// s.Action = "sys-actions"
						// s.Title = "Actions"
						// s.Icon = "star"
						// s.Url = "engine/sys-actions"
						// s.Params["app_only"] = 1
					// });

					// utils.NewMenuItem(item, func(s *utils.MenuItem) {
						// s.ID = "rights";
						// s.Action = "sys-rights";
						// s.Title = "Rights";
						// s.Icon = "settings";
						// s.Url = "engine/sys-rights";
						// s.Params["app_only"] = 1
					// });

					// utils.NewMenuItem(item, func(s *utils.MenuItem) {
						// s.ID = "permissions";
						// s.Action = "sys-permissions";
						// s.Title = "Set Permissions";
						// s.Icon = "security";
						// s.Url = "engine/sys-roles";
					// });
					
				// })
			// }
		},
	})
}
