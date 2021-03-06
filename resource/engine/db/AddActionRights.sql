DROP PROCEDURE [dbo].[AddActionRights]
GO

CREATE PROCEDURE [dbo].[AddActionRights]
(
	@action_id AS int = 0,
	@rights_ids AS varchar(200) = '',
    @visit_id as bigint = 0,
    @action_status_id as int = 0 OUTPUT,
    @action_msg as varchar(200) = '' OUTPUT
)
AS
BEGIN
	SET NOCOUNT ON;

	DECLARE @this_user_id AS int = dbo.F_VisitUserID(@visit_id)
	DECLARE @this_module_id AS int = 2

	SET @action_msg = ''
	SET @action_status_id = 0

	DELETE action_rights WHERE action_id = @action_id

	INSERT INTO action_rights(
		action_id,
		rights_id,
		insert_visit_id,
		inserted_at)
	SELECT
		@action_id,
		value,
		@visit_id,
		getdate()
	FROM dbo.f_split(@rights_ids, ',')
END
GO
