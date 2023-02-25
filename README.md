# Projeto Go Pokemon (não é Pokemon Go)
- Projeto para estudos visando buscar dados de uma API, usando a linguagem Go, e persistir esses dados em um banco, MSSQL.

## Instalar banco pelo docker
- docker run --name sqlservergo -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=m1Nh@s3n4Go" -p 1433:1433 -d mcr.microsoft.com/mssql/server