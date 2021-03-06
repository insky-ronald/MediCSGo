DROP FUNCTION [dbo].[f_user_roles]
GO

CREATE FUNCTION [dbo].[f_user_roles]
/****************************************************************************************************
 Last modified on

****************************************************************************************************/
(
  @user_id as int = 0
) RETURNS varchar(100)
AS
BEGIN
	declare @roles varchar(100)

	select
		@roles = COALESCE(@roles + ',', '') + LTRIM(STR(u.role_id))
	from user_roles u
	where u.user_id = @user_id

	RETURN @roles
END
