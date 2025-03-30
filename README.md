[![es](https://img.shields.io/badge/lang-es-red)](https://github.com/MarceloDanielToledo/BackendGO_AWS/blob/main/README.es.md)

# 📌 Project  
Go API that simulates the Twitter backend, using MongoDB and deployed on AWS with Lambda.  

# 🏗️ Architecture  
The API follows a serverless approach and consists of the following AWS services:

    - 🔹 AWS Lambda: Executes business logic without the need for servers.  
    - 🔹 API Gateway: Exposes the API and routes requests to Lambda.  
    - 🔹 MongoDB: NoSQL database used for persistence.  
    - 🔹 AWS Secrets Manager: Stores credentials and sensitive configurations.  
    - 🔹 Amazon S3: Stores files such as avatars and banners.  
    - 🔹 Amazon CloudWatch: Logs and metrics for monitoring.  

# 📂 Project Structure  

Organization of the codebase and its purpose

    - 📁 awsgo: AWS SDK initialization.  
    - 📁 bd: Logic for MongoDB access (CRUD).  
    - 📁 handlers: Handles HTTP requests and connects to routers.  
    - 📁 routers: Defines API routes.  
    - 📁 jwt: Generates and validates JWT tokens.  
    - 📁 secretmanager: Manages credentials in AWS Secrets Manager.  
    - 📁 models: Defines data structures and models.  
    - 📄 main.go: Application entry point.  
    - 📄 go.mod: Go module dependencies.  
    - 📄 go.sum: Dependency verification hashes.  
