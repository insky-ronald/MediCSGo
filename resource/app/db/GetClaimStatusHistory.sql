DROP PROCEDURE [dbo].[GetClaimStatusHistory]
GO
/****** Object:  StoredProcedure [dbo].[GetClaimDiagnosis]    Script Date: 9/27/2017 1:42:53 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetClaimStatusHistory] 
-- ***************************************************************************************************
-- Last modified on
-- 27-SEP-2017
-- *************************************************************************************************** 
(
    @claim_id int = 0, 
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	SELECT
		id = s.STATUS_ID,
		claim_id = s.CLAIM_NO,
		status_code = s.STATUS,
		status = s.STATUS_NAME,
		status_date = s.STATUS_DATE,
		user_name = s.STATUS_USER,
		user_full_name = u.name
	FROM vw_claim_status_history s
	LEFT OUTER JOIN users u on s.STATUS_USER = u.user_name
	WHERE s.CLAIM_NO = @claim_id
	ORDER BY s.STATUS_DATE desc
END
GO
