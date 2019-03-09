### Seek for optimal golang packages to develop Web application

#### Used Library

-   [Iris Framework](https://github.com/kataras/iris)
-   [Gorm](https://github.com/jinzhu/gorm) to manage Database.
-   [Validator](https://github.com/asaskevich/govalidator)
-   [CLI](https://github.com/urfave/cli) used for:

    1- Database Migration.
    
    2- Database Seed.
    
    3- Run The Web Application.

### Project structure

```
├── README.md
├── app
│   ├── console
│   │   └── migrate.go
│   ├── database.go
│   ├── http
│   │   ├── controllers
│   │   │   ├── city.go
│   │   │   ├── task.go
│   │   │   └── test.go
│   │   ├── middleware
│   │   └── routes.go
│   ├── jobs
│   └── models
│       ├── city.go
│       └── task.go
├── app.go
├── public
└── test.db
```
