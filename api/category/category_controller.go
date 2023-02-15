package category

import (
	"github.com/rs/zerolog"
	"os"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"myAppEcho/configs"
	"myAppEcho/utils"
	"net/http"
	"time"
)

var categoryCollection *mongo.Collection = configs.GetCollection(configs.DB, "categories")
var validate = validator.New()

var logger = zerolog.New(os.Stdout)

func CreateCategory (c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var category Category
	defer cancel()

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&category); validationErr != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newCategory := Category{
		Id: 	primitive.NewObjectID(),
		Title:	category.Title,
		Color: 	category.Color,
	}

	result, err := categoryCollection.InsertOne(ctx, newCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusCreated, utils.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetAllCategory (c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	var categories []Category
	defer cancel()
	results, err := categoryCollection.Find(ctx, bson.M{})

	if(err != nil){
		logger.Error().Msg("Internal server error!!");
		return c.JSON(http.StatusInternalServerError, utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	defer results.Close(ctx)
	for results.Next(ctx){
		var singleCategory Category
		if err = results.Decode(&singleCategory); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
		categories = append(categories, singleCategory)
		
	}
	return c.JSON(http.StatusOK, utils.Response{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": categories}}) 
}