// ****************************************************************************************************
// File name: edit-claim-details.js
// Last modified on
// 
// ****************************************************************************************************
function InvoiceDetailsEdit(params){
	// console.log(params)
	new JPageControl({
		owner: this,
		container: params.container,
		Painter: {
			theme: "data-entry",
			autoHeight: false
		},
		init: function(pg) {			
			pg.NewTab("Invoice", {
				OnCreate: function(tab) {
					tab.content.css("border", "1px solid #92846A").css("overflow-y", "auto");
					new SimpleEditor({
						id: "edit_claim",
						dataset: params.dataset,
						container: tab.content,
						initData: function(editor, data) {
							data.Columns
								.setprops("id", {label:"ID", numeric:true, key: true, readonly:true})
								.setprops("claim_no", {label:"Claim No.", readonly:true})
								.setprops("service_no", {label:"Service No.", readonly:true})
								.setprops("payee_type", {label:"Payee Type", required:true})
								.setprops("reference_no1", {label:"Inter-Company Invoice No."})
								.setprops("invoice_no", {label:"Invoice No.", required:true})
								.setprops("invoice_date", {label:"Invoice Date", type:"date", required:true})
								.setprops("invoice_received_date", {label:"Date Received", type:"date", required:true})
								.setprops("invoice_entry_date", {label:"Entry Date", type:"date", readonly:true, 
									getText: function(column, value) {
										return ("{0} by {1}").format(column.formatDateTime(), column.dataset.get("invoice_entry_user"));
									}})

								// .setprops("sub_product", {label:"Sub-Product", readonly:true})
								// .setprops("hcm_reference", {label:"Assistance Ref."})
								// .setprops("reference_no2", {label:"Reference 2"})
								// .setprops("reference_no3", {label:"Reference 3"})
								
								.setprops("create_user", {label:"User", readonly:true})
								.setprops("create_date", {label:"Date Created", type:"date", format:"datetime", readonly:true, 
									getText: function(column, value) {
										// return ("{0} by {1}").format(column.formatDateTime(), column.dataset.get("create_user"));
										// return ("{0} by {1}").format(column.formatDateTime("d MMMM yyyy"), column.dataset.get("create_user"));
										// return ("{0} by {1}").format(column.formatDateTime("MMMM d, yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
										// return ("{0} by {1}").format(column.formatDateTime2("MMMM d, yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
										return ("{0} by {1}").format(column.formatDateTime2("dd/MM/yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
										// return ("{0} by {1}").format(column.formatDateTime2("dd MMM yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
									}})
								.setprops("update_user", {label:"User", readonly:true})
								.setprops("update_date", {label:"Last Updated", type:"date", format:"datetime", readonly:true,
									getText: function(column, value) {
										// return ("{0} by {1}").format(column.formatDateTime(), column.dataset.get("update_user"));
										// return ("{0} by {1}").format(column.formatDateTime2("MMMM d, yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
										return ("{0} by {1}").format(column.formatDateTime2("dd/MM/yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
										// return ("{0} by {1}").format(column.formatDateTime2("dd MMM yyyy hh:mm:ss tt"), column.dataset.get("create_user"));
									}})
								
								.setprops("claim_currency_code", {label:"Currency", required:true})
								.setprops("claim_currency_rate_date", {label:"Rate Date", type:"date", required:true})
								.setprops("claim_currency_to_base", {label:"Rate to Base Currency", numeric:true, required:true})
								.setprops("claim_currency_to_client", {label:"Rate to Client Currency", numeric:true, required:true})
								.setprops("claim_currency_to_eligibility", {label:"Rate to Eliigibility Currency", numeric:true, required:true})
								
								.setprops("discount_type", {label:"Type"})
								.setprops("discount_percent", {label:"Discount Rate (%)", numeric:true, type:"money", format:"00"})
								// .setprops("plan_name", {label:"Type", required:true})
								
								// .setprops("country_of_incident", {label:"Country of Incident", required:true})

								// .setprops("is_prexisting", {label:"Pre-Existing"})
								// .setprops("is_accident", {label:"Accident"})
								// .setprops("first_symptom_date", {label:"First Symptom", type:"date", required:true})
								// .setprops("first_consultation_date", {label:"First Consultation", type:"date", required:true})
						},
						initEditor: function(editor) {
							editor.AddGroup("General", function(editor) {
								editor.AddEdit("service_no");
								// editor.AddRadioButton("payee_type", {
								editor.AddListBox("payee_type", {
									key: "id",
									value: "value",
									data: [
										{id:"R", value:"Reimbursement"},
										{id:"C", value:"Cashless"},
										{id:"N", value:"Not paid by ISOS"}
									]
								});
								editor.AddEdit("reference_no1");
							});
							
							editor.AddGroup("Provider Invoice", function(editor) {
								editor.AddEdit("invoice_no");
								editor.AddEdit("invoice_date");
								editor.AddEdit("invoice_received_date");
								editor.AddEdit("invoice_entry_date");
							});
								
							editor.AddGroup("Discount", function(editor) {
								editor.AddListBox("discount_type", {
									key: "id",
									value: "value",
									data: [
										{id:"0", value:"No Discount"},
										{id:"1", value:"Invoice Header by Percentage"},
										{id:"3", value:"Invoice Line by Percentage"},
										{id:"4", value:"Invoice Line by Amount"}
									]
								});
								editor.AddEdit("discount_percent");
							});
							
							editor.AddGroup("Currency & Exchange Rates", function(editor) {
								editor.AddEdit("claim_currency_code");
								editor.AddEdit("claim_currency_rate_date");
								editor.AddEdit("claim_currency_to_base");
								editor.AddEdit("claim_currency_to_client");
								editor.AddEdit("claim_currency_to_eligibility");
							});
								
							// editor.AddGroup("Location", function(editor) {
								// editor.AddEdit("country_of_incident");
							// });
								
							// editor.AddGroup("Medical Condition", function(editor) {
								// editor.AddRadioButton("is_prexisting", {
									// key: "id",
									// value: "value",
									// data: [
										// {id:true, value:"Yes"},
										// {id:false, value:"No"}
									// ]
								// });
								// editor.AddRadioButton("is_accident", {
									// key: "id",
									// value: "value",
									// data: [
										// {id:true, value:"Yes"},
										// {id:false, value:"No"}
									// ]
								// });
							// });
								
							// editor.AddGroup("Medical History Dates", function(editor) {
								// editor.AddEdit("first_symptom_date");
								// editor.AddEdit("first_consultation_date");
							// });
							
							editor.AddGroup("Update Info", function(editor) {
								editor.AddEdit("create_date");
								editor.AddEdit("update_date");
							});							
						}
					});
				}
			});
			pg.NewTab("Payment", {
				OnCreate: function(tab) {
					tab.content.css("border", "1px solid #92846A").css("overflow-y", "auto");
				}
			});
			pg.NewTab("Status", {
				OnCreate: function(tab) {
					tab.content.css("border", "1px solid #92846A").css("overflow-y", "auto");
				}
			});
			pg.NewTab("Actions", {
				OnCreate: function(tab) {
					tab.content.css("border", "1px solid #92846A").css("overflow-y", "auto");
				}
			});
		}
	});
};
