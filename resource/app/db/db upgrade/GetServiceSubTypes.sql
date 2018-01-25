DROP PROCEDURE [dbo].[GetServiceSubTypes]
GO

CREATE PROCEDURE [dbo].[GetServiceSubTypes]
-- ***************************************************************************************************
-- Last modified on
-- 
-- *************************************************************************************************** 
(
	@service_type char(3) = '',
	@sub_type varchar(4) = ''
    --@visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	select
		code,
		sub_type,
		display_name
	from service_sub_types
	where service_type = @service_type and code = @sub_type

END
