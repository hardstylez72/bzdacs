# BZDACS - simple lightweight access control system
## Goal - allows setting certain policy for certain user in terms of api routes

## [Demo](http://bzdacs.eblog.cyou/)

# Swagger
[UI](http://bzdacs.eblog.cyou/api/swagger/index.html)
and
[Spec](http://bzdacs.eblog.cyou/api/swagger/source)

### Try on local
via docker-compose
> docker-compose -f docker-compose.build.yml up

### [Client for Golang](https://github.com/hardstylez72/bzdacs-go) 
Easy to use:
```go
...
import acsmw "github.com/hardstylez72/bzdacs-go"

func extractAuthorizationParams(req *http.Request) *acsmw.InputParams {
	return &acsmw.InputParams{
		Login:       "<Your user login>",
		Route:       req.RequestURI,
		Namespace:   "<Your namespace>",
		System:      "<Your system>",
		RouteMethod: req.Method,
	}
}

func main (
	router := NewYourFavoriteRouter()
	router.Use(acsmw.AccessCheck(config.GetBackend().Host, extractAuthorizationParams))
	server.Serve(router)
)
...
```
## Main domains
|Domain|Definition|Attributes|
--- | --- | --- 
|Route|path to api| path, method|
|Group|aggregates routes| code|
|User|aggregates groups and routes| login|
|System|like database name in DBMS| name|
|Namespace|like schema name in DBMS| name|
|Tag|additional label to route| name|



## Stack
- [Postgres](https://github.com/postgres/postgres)
- [VueJS 2.x](https://github.com/vuejs/vue) ([typescript](https://github.com/microsoft/TypeScript))
- [Golang](https://github.com/golang/go)
- happy with [docker](https://github.com/docker)
- ready for [k8s](https://github.com/kubernetes/kubernetes)


![](./docs/schema.png)




<details>
<summary>To contributors</summary>

build
> sh build.sh

tests (integration)
> go test ./test/... 

server
> go run -config ./config/config.local.yaml

ui
> npm run --prefix ./ui serve
</details>

### 

 
 
