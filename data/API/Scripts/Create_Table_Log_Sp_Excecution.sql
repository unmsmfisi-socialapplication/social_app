/****** Object:  Table [dbo].[Execution_Log_GUD]    Script Date: 3/12/2023 22:32:36 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[API_SP_Execution_Log](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[CodeModel] [int] NULL,
	[UserExec] [nvarchar](100) NULL,
	[DateExec] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (
STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF
) ON [PRIMARY]
) ON [PRIMARY]
GO


