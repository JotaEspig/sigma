# 😎 SIGMA
SIGMA (Sistema Intuitivo de Gestão de Matérias Acadêmicas) is an academic system designed to be more intuitive and easier than SIGAA (Sistema Integrado de Gestão de Atividades Acadêmicas)

## 👀 Usage: 
You can run either on Windows or Linux. You can use Docker or not, but if you are not using, you must have installed in your computer: Golang and PostgreSQL. And you must have set environment variables according to how your postgresql is configured. 
Ex.: DB_USERNAME for the database user

### 🐳 Run with docker:
To run with Docker is simple, you can just compose the file docker-compose.yml and then you can access [127.0.0.1:8080](http://127.0.0.1:8080). If you want to configure something, just change the docker-compose.yml or the Dockerfile.

### 🕵️‍♀️ Without docker: 
If you have GO and PostgreSQL installed in your pc, you can use this method. Add the environment variables (listed in docker-compose) and check if postgresql is running. Now you can use:  [`go run .`] to run the app
