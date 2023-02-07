package controller_meter

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetIdMeter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var meterPayload views.PayloadRetriveId
		defer cancel()
		c.BindJSON(&meterPayload)

		if validationErr := validate.Struct(&meterPayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		var resultMeter models.Meter
		var finalPayload views.FinalRetriveId
		result := meterCollection.FindOne(ctx, bson.M{"meter_name": meterPayload.MeterName})
		result.Decode(&resultMeter)
		result.Decode(&finalPayload)
		err := bcrypt.CompareHashAndPassword([]byte(resultMeter.Password), []byte(meterPayload.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
				"data":    "Meter name or password is not valid",
			})
			return
		}

		c.JSON(http.StatusOK,
			bson.M{
				"status":  http.StatusOK,
				"message": "Success",
				"data":    finalPayload,
			},
		)
	}
}
