$AzureSQLServerName = "tcp:sv-socialdb-sql.database.windows.net"
$AzureSQLDatabaseName = "socialdb_sql"


$Cred =  Get-AutomationPSCredential -Name 'Credential_sql'

$SQLOutput = $(Invoke-Sqlcmd `
                -ServerInstance $AzureSQLServerName `
                -Username $Cred.UserName `
                -Password $Cred.GetNetworkCredential().Password `
                -Database $AzureSQLDatabaseName `
                -Query "EXEC [dbo].[process_bi_sentimental]")

Write-Output $SQLOutput
