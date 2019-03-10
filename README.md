### Seek for optimal golang packages to develop Web application

#### Used Library

-   [Iris Framework](https://github.com/kataras/iris)
-   [Gorm](https://github.com/jinzhu/gorm) to manage Database.
-   [Validator](https://github.com/asaskevich/govalidator)
-   [hide](https://github.com/emvi/hide) to hide really records ID
-   [Platform-Agnostic Security Tokens](https://github.com/o1egl/paseto) to secure API tokens
-   [Blake2B](golang.org/x/crypto/blake2b) is to check integrity of files and sent data
-   [CLI](https://github.com/urfave/cli) used for:

    1- Database Migration.

    2- Database Seed.

    3- Run The Web Application.

### Project structure

```sh
.
├── README.md
├── app
│   ├── bootstrap.go
│   ├── console
│   │   ├── kernel.go
│   │   └── migrate.go
│   ├── http
│   │   ├── controllers
│   │   │   ├── city.go
│   │   │   ├── task.go
│   │   │   └── test.go
│   │   ├── middleware
│   │   │   └── hideID.go
│   │   └── routes.go
│   ├── jobs
│   └── models
│       ├── city.go
│       └── task.go
├── app.go
├── public
└── test.db
```
