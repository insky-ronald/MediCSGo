// ****************************************************************************************************
// Last modified on
// 
// ****************************************************************************************************
//==================================================================================================
// File name: edit-claim-types.js
//==================================================================================================
function ClaimTypesEdit(params){
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
		postBack: "app/claim-types",
		init: function(editor) {
			editor.Events.OnInitData.add(function(sender, data) {
			data.Columns
				.setprops("code", {label:"Code", key:true, required:true})
				.setprops("claim_type", {label:"Claim Type Description", required:true})
			});
			
			editor.Events.OnInitEditor.add(function(sender, editor) {
				editor.NewGroupEdit("General", function(editor, tab) {
					editor.AddGroup("Claim Type", function(editor) {
						editor.AddEdit({ID:"code"});
						editor.AddEdit({ID:"claim_type"});
					});
				});
			});
		}
	});
}; 
