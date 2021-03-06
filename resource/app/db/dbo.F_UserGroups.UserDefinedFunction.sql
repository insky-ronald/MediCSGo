SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE FUNCTION [dbo].[F_UserGroups]
(
  @user_name varchar(10),
  @return_codes bit = 1
) RETURNS varchar(max)
AS
BEGIN
	declare @groups varchar(max)

    SELECT 
        @groups = COALESCE(@groups + ',' ,'') + case when @return_codes = 1 then ug.group_id else g.group_name end
    FROM usergroups ug
	JOIN groups g on ug.group_id = g.id
	where ug.user_name = @user_name
    order by g.group_name
    
  	RETURN @groups
END

/*
select * from groups
select top 10 * from usergroups
select dbo.F_UserGroups('JOHN', 1)
select dbo.F_UserGroups('JOHN', 0)

*/
GO
