# PetBlog 
![123](https://github.com/t0rch13/petblog/assets/110306539/88d1ad16-c19d-46bc-a5b1-66fad959a71b)
Makes Batyrkhan, Sara, Adilet
# Description
Petblog is a simple and user-friendly app that helps pet owners organize their daily lives. With Petblog you can:
# Features 
- CRUD operation: 
- Authentication: Implements authentication mechanisms to ensure secure access.
- Authorization Middleware: Controls access to certain features based on user roles.
- Petblog uses: golang as the programming language for the backend and vanilla html for the frontend.
Petblog also uses middleware to process requests and responses, as well as to verify
authentication and authorization of users, passwords are hashed in the database PostgreSQL
# Installation
To install and run Petblog you will need the following components:
- [golang] - programming language for the backend
- [git] - version control system
- [postgresql] - database for storing information about users and tasks
# Data Migrations
Data migrations are managed using the `migrate` utility. Before running the application, ensure that you have performed all necessary migrations by executing the migration commands against your PostgreSQL database.
# To install Petblog, follow these steps:
1.Clone the Petblog repository using the command:
```
git clone https://github.com/t0rch13/petblog.git
```
2.Go to the project folder using the command:
```
cd petblog
```
3.Launch the application using the command:
```
go run main.go
```
4.Open your browser and go to:
```
https://localhost:8000
```

- Worked on server parts Batyrkhan
- Worked on the client parts Sara
- Worked on the creation and organization of files Adilet
# Acknowledgments  
- I want to thank everyone who helped me create Petblog, especially:

- [golang] - for an excellent programming language
- [git] - for a convenient version control system
- [postgresql] - for a simple and fast database
- [docker] - for convenient containerisation
