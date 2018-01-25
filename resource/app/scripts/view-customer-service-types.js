// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: view-customer-service-types.js
// uses edit-customer-service-types.js
//==================================================================================================
function CustomerServiceTypesView(params){
	var name = "app/customer-service-types";
	
	return new JDBGrid({
		params: params,
		options: {
			horzScroll: true
		},
		toolbarTheme:"svg",
		Painter: {
			css: "customer-service-types"
		},
		editForm: function(id, container, dialog) {
			CustomerServiceTypesEdit({
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
					title: "Delete Customer Service Type",
					message: ('Please confirm to delete customer service type "<b>{0}</b>"').format(grid.dataset.lookup(id, "service_description"))
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
						.setprops("service_description", {label:"Customer Service Type Description"})		
				});	

				grid.Events.OnInitColumns.add(function(grid) {
					grid.NewColumn({fname: "code", width: 150, allowSort: true, fixedWidth:true});
					grid.NewColumn({fname: "service_description", width: 380, allowSort: true, fixedWidth:false});
				});
				
				/* grid.Events.OnInitRow.add(function(grid, row) {	
					row.attr("x-status", grid.dataset.get("is_active"))
				}); */
			});
		}
	});	
};
