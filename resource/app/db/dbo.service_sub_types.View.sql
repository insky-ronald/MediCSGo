SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[service_sub_types]
AS
	SELECT        
		MODULE_ID AS service_type,		SERVICE_TYPE AS code,		SERVICE_NAME AS sub_type,		DISPLAY_NAME AS display_name,		ACTIVE AS is_active,		InsertDate AS create_date,		InsertUser AS create_sers,		UpdateDate AS update_date,		UpdateUser AS update_user
	FROM dbo.tb_sub_claim_types


GO
