SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[v_claim_documents]
as
	select 
		n.*,
        dbo.f_document_categories(N.id) as document_categories,
		service_type = cast(case when n.service_id = 0 then 'CLM' else c.service_type end as char(3)),
		service_no = cast(case when n.service_id = 0 then M.claim_no else c.service_no end as varchar(22)),
		invoice_no = cast(case when n.service_id = 0 then 'Claim' else c.invoice_no end as varchar(15)),
		c.invoice_date,
		sequence_no = cast(case when n.service_id = 0 then 0 else c.sequence_no end as tinyint)--
	from claim_documents n
		join claims m on n.claim_id = m.id
		left outer join services c on n.service_id = c.id



GO
