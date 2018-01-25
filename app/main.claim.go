package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"ibsi/dbase"
	// "ibsi/system"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	
	type PageData struct {
		CustomerService dbase.TDataTableRows `json:"data"`
	}
	
	template.NewController(template.Controller {
		Pid: "claim",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			// ts.Add("/app/claim/{id\\/?}")
			ts.Add("/app/claim/{keyid:new}/{memberid:[0-9]+\\/?}")
			ts.Add("/app/claim/{keyid:[0-9]+\\/?}")
			// ts.Add("/app/claim/{new}/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			vars := mux.Vars(r)
			
			var id, memberId int64 = 0, 0
			
			if vars["keyid"] != "new" {
				id = utils.StrToInt64(vars["keyid"])
			}
			
			dbClaim := dbase.Connections["DBApp"].OpenDataTable("GetClaim", dbase.TParameters{"id": id, "visit_id": vid})
			
			// newRecord := 1
			
			if id == 0 {
				memberId = utils.StrToInt64(vars["memberid"])
				
				p.Title = "New Claim"
				p.Nav.PageTitle = "New Claim"
				p.Nav.WindowTitle = "New Claim"
			} else {
				memberId = dbClaim.Get("member_id").(int64)
				
				p.Title = "Claim: " + dbClaim.Get("claim_no").(string)
				p.Nav.PageTitle = "Claim: " + dbClaim.Get("claim_no").(string)
				p.Nav.WindowTitle = dbClaim.Get("claim_no").(string)
			}

			dbMmember := dbase.Connections["DBApp"].OpenDataTable("GetClaimMemberInfo", dbase.TParameters{"claim_id":id, "member_id":memberId, "visit_id":vid})
			
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(id == 0, 1, 0),
				"claim_type": utils.Ifx(id == 0, "", dbClaim.Get("claim_type")),
				"memberId": memberId,
				"data": dbClaim.GetRows(),
				"member": dbMmember.GetRows(),
				"medical_notes": dbase.Connections["DBApp"].OpenDataTable("GetMemberMedicalNotes", dbase.TParameters{"id":memberId, "claim_id":id, "visit_id":vid}).GetRows(),
				"countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "claim", "Claim", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Title = "Details"
					s.Icon = "table-edit"
					s.Action = "admin"
					s.Url = "app/claim-details"
					s.Css = "*"
					s.Run = "ClaimDetailsView"
					s.Params["claim_id"] = id
					// If ClaimType = "MED"
						// .Description = "Medical Claim"
					// End if
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "history"
					s.Title = "Case History"
					s.Icon = "history"
					s.Action = "admin"				
					s.Url = "app/member-case-history"
					s.Params["claim_id"] = id
					s.Params["member_id"] = memberId
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "documents"
					s.Title = "Documents"
					s.Icon = "documents"
					s.Action = "admin"
					s.Url = "app/claim-documents"
					s.Params["claim_id"] = id
					s.Params["service_id"] = 0
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "notes"
					s.Title = "Notes"
					s.Icon = "notes"
					s.Action = "admin"				
					s.Url = "app/claim-notes"
					s.Params["type"] = "C"
					s.Params["claim_id"] = id
					s.Params["service_id"] = 0
				})
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "log"
					s.Title = "Audit Log"
					s.Icon = "timetable"
					s.Action = "admin"				
					s.Url = "app/claim-audit-logs"
					s.Params["claim_id"] = id
					s.Params["service_id"] = 0
				})
			})
			
			if id == 0 {
				return
			}
			
			utils.NewNavigatorItem(p.Nav, "services", "Services", func(item *utils.NavigatorItem) {
		
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "invoices"
					s.Title = "Invoices"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "inv"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "gop"
					s.Title = "Guarantee of Payments"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "gop"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "noc"
					s.Title = "Notification of Claims"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "noc"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "casefee"
					s.Title = "Case Fees"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "cas"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "recovery"
					s.Title = "Recovery of Claims"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "rec"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "costcont"
					s.Title = "Cost Containment"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "cos"
				})
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "flag"
					s.Title = "Flags"
					s.Icon = "table"
					s.Action = "admin"				
					s.Url = "app/claim-invoices"
					s.Css = "*"
					s.Run = "ListClaimServices"
					s.Params["claim_id"] = id
					s.Params["module"] = "flg"
				})
			})
		},
	})
	
	
}
