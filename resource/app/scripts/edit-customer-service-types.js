// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: edit-customer-service-types.js
//==================================================================================================
function CustomerServiceTypesEdit(params){
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
		postBack: "app/customer-service-types",
		init: function(editor) {
			editor.Events.OnInitData.add(function(sender, data) {
			data.Columns
				.setprops("code", {label:"Code", key:true, required:true})
				.setprops("service_description", {label:"Customer Service Type", required:true})
			});
			
			editor.Events.OnInitEditor.add(function(sender, editor) {
				editor.NewGroupEdit("General", function(editor, tab) {
					editor.AddGroup("Customer Service Type", function(editor) {
						editor.AddEdit({ID:"code"});
						editor.AddEdit({ID:"service_description"});
					});
				});
			});
		}
	});
}; 
