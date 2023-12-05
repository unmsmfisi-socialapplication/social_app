CREATE TABLE Response_Matrix (
    ID INT PRIMARY KEY,
    Id_model NVARCHAR(100),
    Text NVARCHAR(500),
    Prediction NVARCHAR(50)
);



CREATE TABLE Execution_Log_GUD (
    ID INT PRIMARY KEY,
    CodeModel INT,
    UserExec NVARCHAR(100),
    DateExec DATETIME
);
