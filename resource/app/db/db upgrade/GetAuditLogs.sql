DROP PROCEDURE [dbo].[GetAuditLogs]
GO

CREATE PROCEDURE [dbo].[GetAuditLogs]
-- ***************************************************************************************************
-- Last modified on
-- 06-OCT-2017
-- *************************************************************************************************** 
(
	@claim_id int = 0,
	@service_id int = 0,
    @page int = 1, 
	@pagesize int = 0, 
	@row_count int = 0 OUTPUT, 
	@page_count int = 0 OUTPUT, 
	@sort varchar(200) = 'create_date',
	@order varchar(10) = 'desc',
    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	DECLARE @where2 varchar(200) = ''
    
	IF @claim_id > 0
		EXEC RunDynamicQueryBuilder2 @id=0, @type='n', @operator='=', @column_name='claim_id', @static_value = @claim_id, @where = @where2 OUTPUT

	IF @service_id > 0
		EXEC RunDynamicQueryBuilder2 @id=0, @type='n', @operator='=', @column_name='service_id', @static_value = @service_id, @where = @where2 OUTPUT

	EXEC RunSimpleDynamicQuery
        @source = 'v_audit_logs',
		@columns = '*',
        @filter = '',
        @where = '',
        @where2 = @where2,
        @page = @page, 
    	@pagesize = @pagesize, 
    	@row_count = @row_count OUTPUT, 
    	@page_count = @page_count OUTPUT,
        @sort = @sort,
        @order = @order

END
GO
