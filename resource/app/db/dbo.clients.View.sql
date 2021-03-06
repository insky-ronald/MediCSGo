SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO


CREATE VIEW [dbo].[clients]
AS
	SELECT        
		COMP_ID AS id, 
		rtrim(COMP_TYPE) AS name_type, 
		rtrim(ACCT_CODE) AS sun_code, 
		rtrim(SPIN_ID) AS soa_prefix, 
		rtrim(NAT_CODE) AS prefix_code,
		--rtrim(FIRST_NAME) AS first_name, 
		--rtrim(MIDDLE_NAME) AS middle_name, 
		--rtrim(LAST_NAME) AS last_name, 
		rtrim(COMP_NAME) AS full_name, 
		ADDRESS_ID AS address_id, 
		CONTACT_ID AS contact_id, 
		--SEX AS gender, 
		--DOB AS dob, 
		RATE_CRCY as currency_code,
		HOME_CTRY AS home_country_code, 
		--ALPHA_ID AS alpha_id, 
		PROV_TYPE AS status_code, 
		CLIENT_LLR_AMT as large_loss_amount,
		HOSP_TYPE AS hospital_type, 
		DOC_TYPE AS doctor_specialisation_code, 
		--TITLE AS title, 
		--EMAIL AS email, 
		HOTLINE AS hotline, 
		InsertDate AS create_date, 
		InsertUser AS create_user, 
		UpdateDate AS update_date, 
		UpdateUser AS update_user
	FROM dbo.tb_names



GO
