# golang-api
Golang API

CRUD in Golang Using Gin Web Framework, Postgres and Redis for Cache.

# Subir a Aplicacao com Docker:
  Acesse a raiz do repositorio e rode: 
  
```  
  make docker  
```

  Parar a Aplicacao: make dockerdown

# Dependencias

Gin-gonic, go-lib/pq(Postgres), go-redis

# Requisitos :

Deixar as Porta (8800, 5432, 6379) do seu host local livre, pois serão essas portas que a aplicacao ira utilizar.

# Endereços e Rotas

Golang API = http://127.0.0.1:8800/

GET http://127.0.0.1:8800/users
GET http://127.0.0.1:8800/users/:username
POST http://127.0.0.1:8800/users
PUT http://127.0.0.1:8800/users/:username
DELETE http://127.0.0.1:8800/users/:username

# Links/Observações

Para Utilizar Docker é necessario ter instalado:

```  
  Docker: https://www.docker.com/

  Docker-Compose: https://docs.docker.com/compose/
  
```  

# Referencias

https://github.com/gin-gonic/gin
