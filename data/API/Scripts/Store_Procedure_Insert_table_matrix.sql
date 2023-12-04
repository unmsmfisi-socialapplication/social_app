/****** Object:  StoredProcedure [dbo].[GetUltimateData]    Script Date: 3/12/2023 21:59:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[SP_GetDataAPI]	@TypeModel INT
AS
BEGIN

----------------------------------------------
--[1] Declaramos Variables
----------------------------------------------
DECLARE @DateUpdate DATETIME
DECLARE @User NVARCHAR(250)
DECLARE @DateUltimateEX DATETIME

----------------------------------------------
--[2] Insertamos en la Tabla de Log
----------------------------------------------
set @User = SYSTEM_USER;
set @DateUpdate = GETDATE();

-------------------------------------------------
--[3] Insertamos la ejecucion de la tabal log
-------------------------------------------------
INSERT INTO dbo.API_SP_Execution_Log
SELECT @TypeModel, @User, @DateUpdate;

------------------------------------
--[4] Borramos los datos de matriz
------------------------------------
DELETE FROM dbo.Response_Matrix;  

------------------------------------
--[5] Insertando datos del modelo
------------------------------------
-- Modelo de SPAM
IF @TypeModel = 1
BEGIN
	SELECT  @DateUltimateEX = DateExec
    FROM dbo.API_SP_Execution_Log
    WHERE CodeModel = 1
    ORDER BY DateExec DESC;

    INSERT INTO dbo.Response_Matrix (Id_model, text, Prediction)
	SELECT 1 AS model_id, text, prediction
	FROM dbo.Master_spam
	WHERE DATEDIFF(hour, @DateUltimateEX, timestamp)>= 0;
END

---- Modelo de POST
ELSE IF @TypeModel = 2
BEGIN
    SELECT  @DateUltimateEX = DateExec
    FROM dbo.API_SP_Execution_Log
    WHERE CodeModel = 2
    ORDER BY DateExec DESC;

    INSERT INTO dbo.Response_Matrix (Id_model, text, Prediction)
	SELECT 2 AS model_id, text, prediction
	FROM dbo.Master_post
	WHERE DATEDIFF(hour, @DateUltimateEX, timestamp)>= 0;
END

-- Modelo de SENTIMENT
ELSE IF @TypeModel = 3
BEGIN
    SELECT  @DateUltimateEX = DateExec
    FROM dbo.API_SP_Execution_Log
    WHERE CodeModel = 3
    ORDER BY DateExec DESC;

    INSERT INTO dbo.Response_Matrix (Id_model, text, Prediction)
	SELECT 2 AS model_id, text, prediction
	FROM dbo.Master_post
	WHERE DATEDIFF(hour, @DateUltimateEX, timestamp)>= 0;
END


END