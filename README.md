# SIGMA
SIGMA (Sistema Intuitivo de Gestão de Matérias Acadêmicas) is an academic system designed to be more intuitive and easier than SIGAA (Sistema Integrado de Gestão de Atividades Acadêmicas)

## Usage: 
You can run either on Windows or Linux. You can use Docker or not, but if you are not using, you must have installed in your computer: Golang and PostgreSQL. And you must have set environment variables according to how your postgresql is configured. 
Ex.: DB_USERNAME for the database user

### With docker:
To run with Docker is simple, you can just compose the file docker-compose.yml. And if you want to configure something, just change the docker-compose.yml or the Dockerfile. After this, you must create the tables in the database. The tables are located in /services/database/SIGMA.sql

### Without docker: 
If you have GO and PostgreSQL installed in your pc, you can use this method. First, create a database called whatever you want, then add a environment variable called DB_DB containing the name of the database. Then run the SIGMA.sql. Now you can use:  [`go run .`] to run the app
