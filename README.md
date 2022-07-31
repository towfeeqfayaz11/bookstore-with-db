```
> go mod init github.com/towfeeq/bookstore-with-db
> go get "github.com/jinzhu/gorm"       --> ORM for sql db
> go get "github.com/jinzhu/gorm/dialects/mysql"
> go get "github.com/gorilla/mux"


------------------------------------------------------------------------------------
bookstore-with-db
├── README.md
├── cmd
│   └── main
│       └── man.go                  --> entrypoint to app
├── go.mod
├── go.sum
└── pkg
    ├── configs
    │   └── config.go               --> contains functions to initialize funtion to db
    ├── controllers
    │   └── controller.go           --> contains functions to handle all the routes
    ├── models
    │   └── model.go                --> contains functions to communcate with db
    ├── routes
    │   └── route.go                --> contains all the routes or api endpoints
    └── utils
        └── util.go                 --> contains utility functions e.g; for marshalling / unmarshalling

8 directories, 9 files
------------------------------------------------------------------------------------

               routes               controllers          models                     DB

         |---> /book/          -->  getAllBooks    -->   getAllBooks    ---|
         |---> /book/          -->  createBook     -->   createBook     ---|
Users  --|---> /book/{id}      -->  getBookById    -->   getBookById    ---|------> books
         |---> /book/{id}      -->  updateBook     -->   updateBook     ---|
         |---> /book/{id}      -->  deleteBook     -->   deleteBook     ---|
```
 
