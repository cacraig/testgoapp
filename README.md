# Test Go API

Example of wiring up Go + Oracle 12g



## Local development requirements:
  
 (1) Install Oracle client: [Mac Oracle client install instructions](https://vanwollingen.nl/install-oracle-instant-client-and-sqlplus-using-homebrew-a233ce224bf)
  
 (2) Install Docker [https://docs.docker.com/docker-for-mac/install/](https://docs.docker.com/docker-for-mac/install/)
  
 (3) Sign up for dockerhub account & accept Oracle license. [https://hub.docker.com/_/oracle-database-enterprise-edition](https://hub.docker.com/_/oracle-database-enterprise-edition)  
 (4) Start Oracle docker container
 ```bash  
  docker pull store/oracle/database-enterprise:12.2.0.1  
  docker run -d -it --name odb -p 1521:1521/tcp store/oracle/database-enterprise:12.2.0.1
 ```
 
 
## To Run: 
```bash  

cd testgoapp  
export TNS_ADMINS=$(pwd)
go run main.go
# Visit localhost:3000

```
