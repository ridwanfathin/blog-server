# blog-server

## Description

RESTful API Web server for simple blog using Echo framework, Gorm, Redis, and Postgresql.

## Installation

* Clone the project:

  ```shell
  git clone https://github.com/ridwanfathin/blog-server
  ```

* Open the directory:

```shell
cd echo-server
```

* Define `GOPATH`:

```shell
export GOPATH=[your project path]
```

* Get external libraries

```shell
go get "github.com/lib/pq"
go get "github.com/jinzhu/gorm"
go get "github.com/labstack/echo"
go get "github.com/garyburd/redigo/redis"
```

* Change environment settings

Change `host`, `user`, `password`, `dbname` file in `src/postgres/pq.go`

* Build binaries:

```shell
make build
```

* Run the server:
  * Linux & Mac

  ```shell
  make run
  ```

  * Windows

  ```shell
  make run
  ```

## API Documentation

[Documentation](https://documenter.getpostman.com/view/2407396/RWTeVMrU#be82841b-3d06-f6f8-6e97-64876360879f)