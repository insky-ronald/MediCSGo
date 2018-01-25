SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[gop_medical_services]
AS
	SELECT        
		INVOICE_ID AS id,		ROOM_EXP AS room_expense,		MISC_EXP AS misc_expense,		REMARKS AS notes,		GOP_NAME AS gop_name,		ADMIT_FIRST_CALL AS admission_first_call,		ADMIT_DOC_RCVD AS admission_document_received,		ADSENDING_DOC AS admission_sending_document,		ADDOC_REC2 AS admission_document_received2,		ADDOC_REC3 AS admission_document_received3,		ADMIN_INITIAL_GOP AS admission_initial_gop,		ADMIT_TAT_FIRST_CALL AS admission_tat_first_call,		ADMIT_TAT_DOC_COMP AS admission_tat_complete_document,		DISCH_FIRST_CALL AS discharge_first_call,		DISCH_DOC_RCVD AS discharge_document_received,		DISSENDING_DOC AS discharge_sending_document,		DISDOC_REC2 AS discharge_document_received2,		DISDOC_REC3 AS discharge_document_received3,		DISCH_FINAL_GOP AS discharge_final_gop,		DISCH_TAT_FIRST_CALL AS discharge_tat_first_call,		DISCH_TAT_DOC_COMP AS discharge_tat_complete_document,		LINK_INVOICE_ID as link_invoice_id,		InsertDate AS create_date,		InsertUser AS create_user,		UpdateDate AS update_date,		UpdateUser AS update_user
	FROM dbo.tb_gop_medical_services


GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[40] 4[49] 2[5] 3) )"
      End
      Begin PaneConfiguration = 1
         NumPanes = 3
         Configuration = "(H (1 [50] 4 [25] 3))"
      End
      Begin PaneConfiguration = 2
         NumPanes = 3
         Configuration = "(H (1 [50] 2 [25] 3))"
      End
      Begin PaneConfiguration = 3
         NumPanes = 3
         Configuration = "(H (4 [30] 2 [40] 3))"
      End
      Begin PaneConfiguration = 4
         NumPanes = 2
         Configuration = "(H (1 [56] 3))"
      End
      Begin PaneConfiguration = 5
         NumPanes = 2
         Configuration = "(H (2 [66] 3))"
      End
      Begin PaneConfiguration = 6
         NumPanes = 2
         Configuration = "(H (4 [50] 3))"
      End
      Begin PaneConfiguration = 7
         NumPanes = 1
         Configuration = "(V (3))"
      End
      Begin PaneConfiguration = 8
         NumPanes = 3
         Configuration = "(H (1[56] 4[18] 2) )"
      End
      Begin PaneConfiguration = 9
         NumPanes = 2
         Configuration = "(H (1 [75] 4))"
      End
      Begin PaneConfiguration = 10
         NumPanes = 2
         Configuration = "(H (1[66] 2) )"
      End
      Begin PaneConfiguration = 11
         NumPanes = 2
         Configuration = "(H (4 [60] 2))"
      End
      Begin PaneConfiguration = 12
         NumPanes = 1
         Configuration = "(H (1) )"
      End
      Begin PaneConfiguration = 13
         NumPanes = 1
         Configuration = "(V (4))"
      End
      Begin PaneConfiguration = 14
         NumPanes = 1
         Configuration = "(V (2))"
      End
      ActivePaneConfig = 0
   End
   Begin DiagramPane = 
      Begin Origin = 
         Top = 0
         Left = 0
      End
      Begin Tables = 
         Begin Table = "tb_gop_medical_services"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 440
               Right = 324
            End
            DisplayFlags = 280
            TopColumn = 13
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 2655
         Alias = 2925
         Table = 1170
         Output = 720
         Append = 1400
         NewValue = 1170
         SortType = 1350
         SortOrder = 1410
         GroupBy = 1350
         Filter = 1350
         Or = 1350
         Or = 1350
         Or = 1350
      End
   End
End
' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'VIEW',@level1name=N'gop_medical_services'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=1 , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'VIEW',@level1name=N'gop_medical_services'
GO
