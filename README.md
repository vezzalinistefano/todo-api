# TODO RestFul API

I made ChatGPT generate me a project idea to learn golang, here is what i was given:

>Build a RESTful API server that handles CRUD operations for a simple data model.
>
>Here are the steps you can follow to build this project:
>
>1. Define a simple data model. For example, you could define a Todo struct with fields for the todo item's ID, title, description, and status.
>2. Use a Go package like mux or gin to set up an HTTP server with endpoints for creating, reading, updating, and deleting todos. You can define these endpoints as functions that handle HTTP requests and return JSON responses.
>3. Use an in-memory data store to store the todos. You can use a Go slice or map to store the todos, or you can use a package like gorm or sqlx to interface with a SQL database.
>4. Define handlers for the HTTP endpoints that interact with the data store to perform CRUD operations on the todos.
>5. Test your API using a tool like curl or Postman. You should be able to create new todos, retrieve a list of all todos, retrieve a single todo by ID, update an existing todo, and delete a todo.
>6. Add authentication and authorization to your API using a package like jwt-go or oauth2.
>7. Deploy your API to a cloud platform like Heroku or AWS Elastic Beanstalk.
>
>That's it! This project should give you a good understanding of how to build a RESTful API server in Go, and how to handle basic CRUD operations using HTTP requests. Good luck!
