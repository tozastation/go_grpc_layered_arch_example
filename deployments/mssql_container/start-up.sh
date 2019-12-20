#!/bin/bash

wait_time=15s

# wait for SQL Server to come up
echo importing data will start in $wait_time...
sleep $wait_time
echo importing data...

for filepath in "/init-data/*.sql"
do
  echo "import: " $filepath
  /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "$SA_PASSWORD" -i $filepath
done