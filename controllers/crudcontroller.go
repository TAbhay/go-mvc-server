package controllers

import (
	"context"
	"go-mvc-server/configs"
	"go-mvc-server/models"
	"go-mvc-server/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var clusterCollection *mongo.Collection = configs.GetCollection(configs.DB, "clusters")
var validateCluster = validator.New()

func CreateCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cluster models.Cluster
		defer cancel()

		// Validate the request body
		if err := c.BindJSON(&cluster); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validateCluster.Struct(&cluster); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Check if the URL already exists
		var existingCluster models.Cluster
		err := clusterCollection.FindOne(ctx, bson.M{"url": cluster.URL}).Decode(&existingCluster)
		if err == nil {
			// Cluster with the same URL already exists
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Cluster with the specified URL already exists"}})
			return
		}
		if err != mongo.ErrNoDocuments {
			// Some other error occurred
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newCluster := models.Cluster{
			ID:    primitive.NewObjectID(),
			Name:  cluster.Name,
			Token: cluster.Token,
			URL:   cluster.URL,
		}

		result, err := clusterCollection.InsertOne(ctx, newCluster)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
func GetClusters() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var clusters []models.Cluster
		defer cancel()

		results, err := clusterCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleCluster models.Cluster
			if err = results.Decode(&singleCluster); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			clusters = append(clusters, singleCluster)
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": clusters}})
	}
}

func GetCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		clusterID := c.Param("id")
		var cluster models.Cluster
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(clusterID)

		err := clusterCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&cluster)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": cluster}})
	}
}

func UpdateCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		clusterID := c.Param("id")
		var cluster models.Cluster
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(clusterID)

		// Validate the request body
		if err := c.BindJSON(&cluster); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validateCluster.Struct(&cluster); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"name": cluster.Name, "token": cluster.Token, "url": cluster.URL}

		result, err := clusterCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Get updated cluster details
		var updatedCluster models.Cluster
		if result.MatchedCount == 1 {
			err := clusterCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedCluster)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedCluster}})
	}
}

func FilterClusters() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := c.Query("url")
		name := c.Query("name")

		filter := bson.M{}
		if url != "" {
			filter["url"] = url
		}
		if name != "" {
			filter["name"] = name
		}

		var clusters []models.Cluster
		results, err := clusterCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleCluster models.Cluster
			if err = results.Decode(&singleCluster); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			clusters = append(clusters, singleCluster)
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": clusters}})
	}
}

func DeleteCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		clusterID := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(clusterID)

		result, err := clusterCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Cluster with specified ID not found!"}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Cluster successfully deleted!"}})
	}
}
