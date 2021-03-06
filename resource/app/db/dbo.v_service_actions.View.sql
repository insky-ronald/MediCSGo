SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [dbo].[v_service_actions]
AS
	SELECT  
		c.*,
		ac.action_type,
		a.action
			--ac.CLASS_NAME AS ActionClass, 
			--a.ACTION_NAME AS ActionName
	FROM dbo.service_actions c
	LEFT OUTER JOIN dbo.action_types ac ON c.action_type_code = ac.code
	LEFT OUTER JOIN dbo.ACTIONS a ON C.action_type_code = a.action_type AND c.action_code = a.code
	--ACTION_CLASS AND C.ACTION_CODE = a.ACTION_CODE

GO
