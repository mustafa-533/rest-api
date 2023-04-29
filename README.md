# RESTful JSON API with CRUD operations for Books, Categories, and Articles Entities

This project is a simple RESTful JSON API built with Golang and Gin that supports CRUD (Create, Read, Update, Delete) operations for Books, Categories, and Articles. Each entity is stored with a different persistence mechanism: **Books** are stored in a MySQL database, **Categories** are stored in a text file, and **Articles** are stored in memory.


### Getting Started
To get started with this project, you need to have Golang and Gin installed on your system. You also need to have a MySQL database set up to store the books.

## Running the API
Once MySQL database connection is configured, you can run the API by running the following command in your terminal:
> go run main.go

This will start the API on http://localhost:{{PORT}}. PORT is configurable through environment variable.  

Following env variables are required for this api

* **LOG_LEVEL**: Set the log level. Value could be one of the following debug, info, errro. By default its value is **debug**
* **LISTEN_HTTP_PORT**: Poert at which server should listen for the incoming requests.
* **MYSQL_URI**: MySQL connecton string

## API Endpoints
The following endpoints are available in this API:

#### Books
* **GET /books**: Get all books
* **GET /books/:id**: Get a book by ID
* **POST /books**: Create a new book
* **PUT /books/:id**: Update a book by ID
* **DELETE /books/:id**: Delete a book by ID

#### Categories
* **GET /categories**: Get all categories
* **GET /categories/:id**: Get a category by ID
* **POST /categories**: Create a new category
* **PUT /categories/:id**: Update a category by ID
* **DELETE /categories/:id**: Delete a category by ID

#### Articles
* **GET /articles**: Get all articles
* **GET /articles/:id**: Get an article by ID
* **POST /articles**: Create a new article
* **PUT /articles/:id**: Update an article by ID
* **DELETE /articles/:id**: Delete an article by ID

## Dockerizing the API
You can also run the API in a Docker container by building the Docker image using the following command:
> docker build -t my-rest-api .
