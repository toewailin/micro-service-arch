#!/bin/bash

# Step 1: Create the main project folder
echo "Creating the project directory..."
mkdir micro-service-arch
cd micro-service-arch

# Step 2: Initialize Go module for the main project
echo "Initializing Go module for the main project..."
go mod init micro-service-arch

# Step 3: Create the main folder structure
echo "Creating main folder structure..."
mkdir api-gateway product-service order-service user-service faq-service

# Step 4: Initialize Go module for each microservice and create subfolders and files

# API Gateway
echo "Creating API Gateway files..."
mkdir -p api-gateway
cd api-gateway
touch main.go
cd ..

# Product Service
echo "Creating Product Service files..."
mkdir -p product-service/{cmd,config,handlers,models,routes,services,db}
cd product-service
go mod init product-service
touch cmd/main.go
touch config/config.go
touch handlers/product_handler.go
touch models/product.go
touch routes/product_routes.go
touch services/product_service.go
touch db/product.db
cd ..

# Order Service
echo "Creating Order Service files..."
mkdir -p order-service/{cmd,config,handlers,models,routes,services,db}
cd order-service
go mod init order-service
touch cmd/main.go
touch config/config.go
touch handlers/order_handler.go
touch models/order.go
touch routes/order_routes.go
touch services/order_service.go
touch db/order.db
cd ..

# User Service
echo "Creating User Service files..."
mkdir -p user-service/{cmd,config,handlers,models,routes,services,db}
cd user-service
go mod init user-service
touch cmd/main.go
touch config/config.go
touch handlers/user_handler.go
touch models/user.go
touch routes/user_routes.go
touch services/user_service.go
touch db/user.db
cd ..

# FAQ Service
echo "Creating FAQ Service files..."
mkdir -p faq-service/{cmd,config,handlers,models,routes,services,db}
cd faq-service
go mod init faq-service
touch cmd/main.go
touch config/config.go
touch handlers/faq_handler.go
touch models/faq.go
touch routes/faq_routes.go
touch services/faq_service.go
touch db/faq.db
cd ..

# Step 5: Create Docker Compose File
echo "Creating docker-compose.yml..."
touch docker-compose.yml

echo "Project setup completed!"

# Step 6: Open VS Code
code .
