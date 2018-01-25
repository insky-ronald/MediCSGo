// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: view-auditlog-types.js
// uses edit-auditlog-types.js
//==================================================================================================
function AuditLogTypesView(params){
	var name = "app/auditlog-types";
	
	return new JDBGrid({
		params: params,
		options: {
			horzScroll: true
		},
		toolbarTheme:"svg",
		Painter: {
			css: "auditlog-types"
		},
		editForm: function(id, container, dialog) {
			AuditLogTypesEdit({
				url: ("?code={0}").format(id),
				container: container,
				containerPadding: 0,				
				showToolbar: false,
				pageControlTheme: "data-entry",
				fillContainer: true,
				dialog: dialog,
			})
		},
		init: function(grid) {
			grid.Methods.add("deleteConfirm", function(grid, id) {
				return {
					title: "Delete Audit Log Type",
					message: ('Please confirm to delete audit log type "<b>{0}</b>"').format(grid.dataset.lookup(id, "description"))
				}
			});
			
			grid.Events.OnInitGrid.add(function(grid) {
				grid.optionsData.url = name;
				grid.options.showToolbar = true;
				grid.options.horzScroll = false;
				grid.options.showPager = true;
				grid.options.showSummary = false;
				grid.options.cardView = false;
				grid.options.autoScroll = true;
				grid.options.allowSort = true;
				grid.options.showSelection = false;
				grid.options.showBand = false;
				grid.options.simpleSearch = true;
				grid.options.simpleSearchField = "filter";
				
				grid.Events.OnInitDataRequest.add(function(grid, dataParams) {
					dataParams
						.addColumn("page", 1, {numeric:true})
						.addColumn("pagesize", 50, {numeric:true})
						.addColumn("sort", "code")
						.addColumn("order", "asc")
						.addColumn("filter", "")
				});
				
				grid.Events.OnInitData.add(function(grid, data) {
					data.Columns
						.setprops("code", {label:"Code", numeric:false, key:true})
						.setprops("description", {label:"Description"})	
						.setprops("log_type_desc", {label:"Log Type"})
				});	

				grid.Events.OnInitColumns.add(function(grid) {
					grid.NewColumn({fname: "code", width: 100, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "description", width: 400, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "log_type_desc", width: 100, allowSort: true, fixedWidth:false});
				});
				
				grid.Events.OnInitRow.add(function(grid, row) {	
					row.attr("x-status", grid.dataset.get("is_active"))
				});
			});
		}
	});	
};
