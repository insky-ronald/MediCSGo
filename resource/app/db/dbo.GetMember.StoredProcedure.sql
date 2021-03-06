SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetMember] 
 
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
    @id int = 0,
    @visit_id as bigint = 0
)
as
begin
	SET NOCOUNT ON

	select	
		m.*,
		h.notes,
		md.name_type,
		md.first_name,
		md.middle_name,
		md.last_name,
		md.full_name, 
		md.gender, 
		md.dob, 
		md.home_country_code,
		md.nationality_code,
		md.alpha_id,
		md.address_id,
		md.contact_id,
		p.policy_no,
		ph.full_name as policy_holder,
		p.issue_date as policy_issue_date,
		p.start_date as policy_start_date,
		p.end_date as policy_end_date,
        p.product_code,
		r.product_name,
		pd.plan_name,
		mm.full_name as main_member,
		c.full_name as client_name,
		cast(1 as bit) as auto_gen_certificate,
		cast(0 as bit) as create_policy,
		cast(0 as bit) as member_is_policy_holder
	from members m
		LEFT OUTER JOIN plan_history H on m.id = h.member_id and m.sequence_no = h.sequence_no
		LEFT OUTER JOIN policies p on m.policy_id = p.id
		LEFT OUTER JOIN names ph on p.name_id = ph.id
		LEFT OUTER JOIN names md on m.name_id = md.id -- CLIENT MD ON M.COMP_ID = MD.COMP_ID
		LEFT OUTER JOIN plans pd on m.plan_code = pd.code-- PLANDEF PD ON M.PLAN_CODE = PD.PLAN_CODE
		LEFT OUTER JOIN products r on p.product_code = r.code
		LEFT OUTER JOIN names mm on m.main_name_id = mm.id -- CLIENT MM ON M.MAIN_ID = MM.COMP_ID
		LEFT OUTER JOIN names c on m.client_id = c.id
	where m.id = @id
end

GO
