# GoBodyguard

This is an ACL-based authorization self-contained and independent microservice which shall be a drop-in into other microservices-based systems.

## References

* https://www.codeproject.com/articles/1056853/lightning-fast-access-control-lists-in-csharp

## Installation

* go get github.com/karuppaiah/gobodyguard
* cd $GOPATH/src/github.com/karuppaiah/gobodyguard
* install dep(https://github.com/golang/dep)
* dep ensure
* cd migrations
* dbmate up
* cd ..
* go run auth.go or use fresh(https://github.com/gravityblast/fresh)

## Package structure

* app : all auth related functionality are present here
    1. authdata.go : models
    2. authdatastorage.go : repository with GORM(https://github.com/jinzhu/gorm)
    3. authdatastore.go: Interface for repository
    4. authlogic.go: Usecase layer for auth(logic done here)
    5. authopns.go: Interface for usecase layer, exposed to presentation layer
    6. config.go: Singleton config to pick from Environment variables
    7. constants.go: All constants in this package
    8. dbinit.go: Database singleton initialization
    9. httpauth.go: Presentation layer with Gin (https://github.com/gin-gonic/gin)
    10. util.go: All helper functions
* migrations : Migration script  executed using (https://github.com/amacneil/dbmate)
* Proto : This service implementation in gRPC
* auth.go : main file to be executed
* .env: For running docker-compose
* Gopkg.toml: Auto generated file by dep package (don't modify or delete it)
* Gopkg.lock: Auto generated file by dep package ( don't modifu or delete it)
* runner.conf : file used by fresh package(if used)
* LICENSE : license file


# TODO :
- [ ] Test all API's
- [ ] Load Data on startup and keep it in memory
- [ ] Docker image and publish in hub.docker.com
- [ ] Write terraform for infrastructure
- [ ] Write ansible for deploy and config
- [ ] Deploy in AWS cloud/Google cloud
- [ ] redis storage for data storage
- [ ] Write unit testing
- [ ] Flutter/web assembly frontend
- [ ] Utilites for message Q like kafka(producer and consumer)
- [ ] Stress testing scripts