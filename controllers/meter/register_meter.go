package controller_meter

import (
	"context"
	"net/http"
	"time"
	"tutorial/configs"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

var meterCollection *mongo.Collection = configs.GetCollection(configs.DB, "meter")

func RegisterMeter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var meter models.Meter
		defer cancel()
		c.BindJSON(&meter)

		if validationErr := validate.Struct(&meter); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		count, err_ := meterCollection.CountDocuments(ctx, bson.M{"meter_name": meter.MeterName})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count >= 1 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Meter name has been taken"})
			return
		}

		bytes, errors := bcrypt.GenerateFromPassword([]byte(meter.Password), 14)
		if errors != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Password tidak valid"})
		}

		newMeter := models.Meter{
			MeterName: meter.MeterName,
			Password:  string(bytes),
			CreatedAt: time.Now(),
		}

		result, err := meterCollection.InsertOne(ctx, newMeter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		finalView := views.MeterView{
			MeterId:   result.InsertedID,
			MeterName: meter.MeterName,
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})
	}
}
