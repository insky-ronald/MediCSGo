// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: view-noc-types.js
// uses edit-noc-types.js
//==================================================================================================
function NOCTypesView(params){
	var name = "app/noc-types";
	
	return new JDBGrid({
		params: params,
		options: {
			horzScroll: true
		},
		toolbarTheme:"svg",
		Painter: {
			css: "noc-types"
		},
		editForm: function(id, container, dialog) {
			NOCTypesEdit({
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
					title: "Delete Notification of Claims Type",
					message: ('Please confirm to delete notification of claims type "<b>{0}</b>"').format(grid.dataset.lookup(id, "service_description"))
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
						.setprops("service_description", {label:"Notification of Claims Type Description"})		
						.setprops("display_name", {label:"Display Name"})
				});	

				grid.Events.OnInitColumns.add(function(grid) {
					grid.NewColumn({fname: "code", width: 150, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "service_description", width: 380, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "display_name", width: 300, allowSort: true, fixedWidth:false});
				});
				
				grid.Events.OnInitRow.add(function(grid, row) {	
					row.attr("x-status", grid.dataset.get("is_active"))
				});
			});
		}
	});	
};
