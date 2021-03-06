SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetProduct] 
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
    @code varchar(10) = '', 
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
    
	SELECT
		r.*,
		n.full_name as client_name,
		f.float_name
	FROM products r
	join names n on r.client_id = n.id
	left outer join floats f on r.float_id = f.id
	WHERE r.code = @code
END



GO
