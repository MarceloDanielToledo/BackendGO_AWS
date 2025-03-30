[![en](https://img.shields.io/badge/lang-en-red)](https://github.com/MarceloDanielToledo/BackendGO_AWS/blob/main/README.md)

# 📌 Proyecto  

 API en Go que simula el backend de Twitter, utilizando MongoDB y desplegada en AWS con Lambda.  

# 🏗️ Arquitectura  

La API sigue un enfoque serverless y está compuesta por los siguientes servicios en AWS:  

    - 🔹 AWS Lambda: Ejecuta la lógica de negocio sin necesidad de servidores.  
    - 🔹 API Gateway: Expone la API y enruta solicitudes a Lambda.  
    - 🔹 MongoDB: Base de datos NoSQL utilizada para la persistencia.  
    - 🔹 AWS Secrets Manager: Almacena credenciales y configuraciones sensibles.  
    - 🔹 Amazon S3: Almacena archivos como avatares y banners de usuario.  
    - 🔹 Amazon CloudWatch: Registra logs y métricas para monitoreo.  

# 📂 Estructura del Proyecto  

 Organización del código y su propósito.  

    - 📁 awsgo: Inicialización del SDK de AWS.  
    - 📁 bd: Lógica de acceso a MongoDB (CRUD).  
    - 📁 handlers: Manejo de solicitudes HTTP y conexión con routers.  
    - 📁 routers: Definición de rutas de la API.  
    - 📁 jwt: Generación y validación de tokens JWT.  
    - 📁 secretmanager: Gestión de credenciales en AWS Secrets Manager.  
    - 📁 models: Definición de estructuras y modelos de datos.  
    - 📄 main.go: Punto de entrada de la aplicación.  
    - 📄 go.mod: Módulo de dependencias de Go.  
    - 📄 go.sum: Hashes de verificación de dependencias.  
