// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: view-service-status-codes.js
// uses edit-service-status-codes.js
//==================================================================================================
function ServiceStatusCodesView(params){
	var name = "app/service-status-codes";
	
	return new JDBGrid({
		params: params,
		options: {
			horzScroll: true
		},
		toolbarTheme:"svg",
		Painter: {
			css: "service-status-codes"
		},
		editForm: function(id, container, dialog) {
			ServiceStatusCodesEdit({
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
					title: "Delete Service Status Code",
					message: ('Please confirm to delete service status code "<b>{0}</b>"').format(grid.dataset.lookup(id, "sub_status"))
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
						.addColumn("pagesize", 25, {numeric:true})
						.addColumn("sort", "module")
						.addColumn("order", "asc")
						.addColumn("filter", "")
				});
				
				grid.Events.OnInitData.add(function(grid, data) {
					data.Columns
						.setprops("code", {label:"Code", numeric:false, key:true})
						.setprops("module", {label:"Service Type"})	
						.setprops("main_status", {label:"Main Status"})
						.setprops("sub_status_code", {label:"Sub-Status Code"})
						.setprops("sub_status", {label:"Sub-Status Description"})
				});	

				grid.Events.OnInitColumns.add(function(grid) {
					grid.NewColumn({fname: "module", width: 200, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "main_status", width: 180, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "sub_status_code", width: 150, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "sub_status", width: 250, allowSort: true, fixedWidth:false});
				});
				
				grid.Events.OnInitRow.add(function(grid, row) {	
					row.attr("x-status", grid.dataset.get("is_system"))
				});
			});
		}
	});	
};
