SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[GetUsers] 
-- ***************************************************************************************************
-- Last modified on
-- 14-OCT-2014 ihms.0.0.1.0
-- *************************************************************************************************** 
(
	@group_id as varchar(10) = '',
	@lookup int = 0,
	@filter as varchar(100) = '',

    @page int = 1, 
	@pagesize int = 0, 
	@row_count int = 0 OUTPUT, 
	@page_count int = 0 OUTPUT, 
	@sort varchar(200) = 'user_name',
	@order varchar(10) = 'asc',

    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

    --DECLARE @user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @user_id AS int = 0
    DECLARE @where nvarchar(500), @where2 nvarchar(500)
	DECLARE @columns varchar(1024) = 'id,user_name,name,designation,phone_no,is_active,is_supervisor,case when password is null then 0 else 1 end as has_password, cast(dbo.F_UserGroups(user_name, 1) as varchar(255)) as group_codes, cast(dbo.F_UserGroups(user_name, 0) as varchar(255)) as groups'
	--DECLARE @columns varchar(1024) = 'id,user_name,name,password,designation,phone_no,is_active,is_supervisor,cast(dbo.F_UserGroups(user_name, 1) as varchar(255)) as group_codes, cast(dbo.F_UserGroups(user_name, 0) as varchar(255)) as groups'
	IF @lookup > 0  
	BEGIN
		SET @pagesize = 1000000
		SET @columns = 'user_name,name'
		SET @sort = 'name'
	END

	SET @where = '[user_name] like @filter or [name] like @filter'

	IF LEN(@group_id) > 0
		SET @where = '(' + @where + ') or (' +'[user_name] in (select [user_name] from usergroups where group_id = '''+@group_id+''')'+ ')'
		--SET @where2 = '[user_name] in (select [user_name] from usergroups where group_id = '''+@group_id+''')'
    
	EXEC RunSimpleDynamicQuery
        @source = 'users',
        @columns = @columns,
		--@columns = 'id,user_name,name,designation,phone_no,is_active,is_supervisor,case when password is null then 0 else 1 end as has_password, cast(dbo.F_UserGroups(user_name, 1) as varchar(255)) as group_codes, cast(dbo.F_UserGroups(user_name, 0) as varchar(255)) as groups',

        @filter = @filter,
        @where = @where,
		@where2 = @where2,
        
        @page = @page, 
    	@pagesize = @pagesize, 
    	@row_count = @row_count OUTPUT, 
    	@page_count = @page_count OUTPUT,
        @sort = @sort,
        @order = @order

END


GO
