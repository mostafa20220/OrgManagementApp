package main

/*

Objective

The objective of this project is to develop a comprehensive API using Golang. This API
will handle tasks such as token management, CRUD operations for organizations, and
user invitations with limited permissions. Additionally, it will seamlessly integrate with
MongoDB using Docker and Docker Compose.

Deliverables

● For project execution, 'docker compose up' must serve as the sole requirement
● Ensure that the endpoints precisely conform to the schema outlined in this
document.
● Use Bearer Token for authorization
● Users who are invited to an organization should only be able to read organization
data.
● Please share the code with us by placing it in a public repository on your personal
GitHub account or any other source control platform.
● Upon completion, kindly share the link to your repository and any relevant
resources by responding to the task email.
*/

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/config"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/pkg/database/mongodb/repository"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/pkg/utils"
)

func main(){

	// Initialize configuration, database, and other dependencies
	config := config.LoadConfig("config/app-config.yaml")
	dbClient, err := repository.NewMongoDBClient(config.Database.MongoURI)
	if err != nil {
		panic(err)
	}

	// Set up the Gin router and routes
	router := gin.Default()
	routes.SetupRoutes(router, dbClient)

	// Start the server
	router.Run(":8080")
}