SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetClaimMemberInfo] 
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
    @claim_id int = 0,
	@member_id int = 0,
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	declare @info table (
		member_id int,
		claim_id int,
		history_id int,
		name_id int,
		main_name_id int,
		certificate_no varchar(25),
		member_name varchar(100),
		main_member varchar(100),
		--is_vip varchar(5),
		--is_blacklist varchar(5),
		is_vip bit,
		is_blacklist bit,
		is_vip_blackllist varchar(10),
		dob datetime,
		age int,
		sex varchar(10),
		plan_code char(15),
		plan_id int,
		plan_type char(10),
			PLAN_LMT_TEM varchar(60),
			INV_DTL_TEM varchar(60),
		plan_name varchar(60),
        sub_product varchar(15),
		plan_sub_product varchar(100),
		plan_inception_date datetime,
		plan_start_date datetime,
		plan_end_date datetime,		
		status char(1),
        member_active bit,
		policy_id int,
		policy_name_id int,
		policy_no varchar(20),
		policy_issue_date datetime,
		policy_start_date datetime,
		policy_end_date datetime,
		policy_holder varchar(100),
		product_code char(10),
		product_name varchar(60),
		float_id int,
		client_id int,
		client_name varchar(100),
        client_active bit,
        client_inactive_date datetime,
		base_currency_code char(3),
		eligibility_currency_code char(3),
		client_currency_code char(3),
		claim_reference1 varchar(25),
		claim_reference2 varchar(25),
		claim_reference3 varchar(25),
		confirm_policy bit,
		confirm_plan bit,
		notes varchar(max),
		waiting_period_date datetime,
		waiting_period_days int,
		waiting_period_months int,
		waiting_period_days_months varchar(100)
	)


	IF @member_id = 0
		SELECT @member_id = member_id from claims WHERE id = @claim_id

	INSERT INTO @info
	EXEC ssp_claim_member_info @IP_ID = @member_id, @CLAIM_NO = @claim_id

	UPDATE @Info set claim_id = @claim_id

	SELECT * FROM @info
END

GO
