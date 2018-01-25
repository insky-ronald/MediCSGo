// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: edit-cost-containment-types.js
//==================================================================================================
function CostContainmentTypesEdit(params){
	return new FormEditor({
		id: params.id,
		dialog: params.dialog,
		container: params.container,
		labelWidth: 150,
		containerPadding: defaultValue(params.containerPadding, 10),
		pageControlTheme: defaultValue(params.pageControlTheme, "main"),
		fillContainer: defaultValue(params.fillContainer, false),
		showToolbar: params.showToolbar,
		url: params.url,
		postBack: "app/cost-containment-types",
		init: function(editor) {
			editor.Events.OnInitData.add(function(sender, data) {
			data.Columns
				.setprops("code", {label:"Code", key:true, required:true})
				.setprops("service_description", {label:"Cost Containment Type", required:true})
				.setprops("display_name", {label:"Display Name"})
				.setprops("is_active", {label:"Status"})
			});
			
			editor.Events.OnInitEditor.add(function(sender, editor) {
				editor.NewGroupEdit("General", function(editor, tab) {
					editor.AddGroup("Recovery Type", function(editor) {
						editor.AddEdit({ID:"code"});
						editor.AddEdit({ID:"service_description"});
						editor.AddEdit({ID:"display_name"});
					});
					editor.AddGroup("Option", function(editor) {
						editor.AddRadioButton("is_active", {
							key: "id",
							value: "value",
							data: [
								{id: "1", value: "Active"},
								{id: "0", value: "Inactive"}
							]
						});
					});
				});
			});
		}
	});
}; 
