DROP PROCEDURE [dbo].[GetDoctors]
GO

CREATE PROCEDURE [dbo].[GetDoctors]
-- ***************************************************************************************************
-- Last modified on
-- 
-- EXEC GetDoctors @id = 292056
-- *************************************************************************************************** 
(
	@id int = 0,
	@filter as varchar(100) = '',

	@action int = 0, -- 0:list, 1:lookup, 10:for editing, 20:for new record, 50:fetch updated data
    @page int = 1, 
	@pagesize int = 0, 
	@row_count int = 0 OUTPUT, 
	@page_count int = 0 OUTPUT, 
	@sort varchar(200) = 'name',
	@order varchar(10) = 'asc',

    @visit_id as bigint = 0
)
AS
BEGIN
    SET NOCOUNT ON;

	IF @action in (10,20,50)
	BEGIN
		SELECT * FROM v_doctors WHERE id = @id
	END ELSE
	BEGIN
		DECLARE @user_id AS int = 0
		DECLARE @where nvarchar(500), @where2 nvarchar(500)
		DECLARE @columns varchar(100) = '*' 
/*
		IF @lookup > 0  
		BEGIN
			SET @pagesize = 1000000
			SET @columns = 'name, country'
		END */

		SET @where = '[name] like @filter or [full_name] like @filter'

		--IF @id > 0 
		--   SET @where2 = '[provider_type] = ''D'' and [id] = ' + CAST(@id as varchar)
		--ELSE IF (@id = 0 and @pagesize in (25, 50, 75, 100))
		--	SET @where2 = '[provider_type] = ''D'''	
		--ELSE IF (@id = 0 and @lookup = 0) 
		   --SET @where2 = '[provider_type] = ''D'' and [id] = -1'
		--ELSE IF (@id = 0 and @lookup > 0) 
		   --SET @where2 = '[provider_type] = ''D'''
    
		EXEC RunSimpleDynamicQuery
			@source = 'v_doctors',
			@columns = @columns,
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
END
GO
