package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"loro-handler/controllers"
	"loro-handler/middleware"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.Use(middleware.SetCorsHeader)
	r.Use(middleware.RecordUaAndTime)
	r.Use(middleware.ErrorHandler)
	controllers := controllers.GetControllers()
	r.POST("/user", controllers.UserController.GetUserInfo)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
