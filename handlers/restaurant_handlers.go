package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type restaurantResponse struct {
	Id            int64              `json:"id"`
	Name          string             `json:"name"`
	ImageUrl      string             `json:"image_url"`
	AverageRating float64            `json:"average_rating"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	Categories    []categoryResponse `json:"restaurant_categories"`
}

type restaurantDetailsResponse struct {
	Id             int64                   `json:"id"`
	Name           string                  `json:"name"`
	ImageUrl       string                  `json:"image_url"`
	AverageRating  float64                 `json:"average_rating"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      time.Time               `json:"updated_at"`
	FoodCategories []*foodCategoryResponse `json:"food_categories"`
}

type foodCategoryResponse struct {
	Id        int64              `json:"id"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	FoodItems []foodItemResponse `json:"food_items"`
}

type foodItemResponse struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type categoryResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Get All Restaurants with Categories
func (cfg *ApiConfig) GetAllRestaurantWithCategories(w http.ResponseWriter, r *http.Request) {
	// Check Token
	token, err := GetBearerToken(r.Header)
	if err != nil {
		cfg.LogError("Cannot read the authorization token", err)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to get all restaurants")
		return
	}

	_, jwtErr := ValidateJWT(token, cfg.TokenSecret)
	if jwtErr != nil {
		cfg.LogError("Validate jwt failed", jwtErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to get all restaurants")
		return
	}

	// Fetch data
	restaurantsDb, fetchErr := cfg.Db.GetAllRestaurantsWithCategories(r.Context())
	if fetchErr != nil {
		cfg.LogError("Failed to fetch restaurants", fetchErr)
		RespondWithError(w, http.StatusInternalServerError, "Failed to get restaurants. Please try again.")
		return
	}

	restaurantsMap := map[int64]*restaurantResponse{}
	result := []*restaurantResponse{}

	for _, restaurantDb := range restaurantsDb {
		restaurant, ok := restaurantsMap[restaurantDb.RestaurantID]
		if !ok {
			// Doesn't exist yet
			// Create new restaurant object
			avgRating, parseErr := strconv.ParseFloat(restaurantDb.RestaurantAverageRating, 64)
			if parseErr != nil {
				cfg.LogError(fmt.Sprintf("Cannot parse average rating: %v", restaurantDb.RestaurantAverageRating), parseErr)
				RespondWithError(w, http.StatusInternalServerError, "There was something wrong while fetching the restaurants. Please try again.")
				return
			}
			restaurantResponse := restaurantResponse{
				Id:            restaurantDb.RestaurantID,
				Name:          restaurantDb.RestaurantName,
				ImageUrl:      restaurantDb.RestaurantImageUrl,
				AverageRating: avgRating,
				CreatedAt:     restaurantDb.RestaurantCreatedAt,
				UpdatedAt:     restaurantDb.RestaurantUpdatedAt,
				Categories: []categoryResponse{
					{
						Id:        restaurantDb.RestaurantCategoryID,
						Name:      restaurantDb.RestaurantCategoryName,
						CreatedAt: restaurantDb.RestaurantCategoryCreatedAt,
						UpdatedAt: restaurantDb.RestaurantCategoryUpdatedAt,
					},
				},
			}

			restaurantsMap[restaurantDb.RestaurantID] = &restaurantResponse
			result = append(result, &restaurantResponse)

		} else {
			// Restaurant already exists
			// Create new category and append it to the restaurant's category
			newCategory := categoryResponse{
				Id:        restaurantDb.RestaurantCategoryID,
				Name:      restaurantDb.RestaurantCategoryName,
				CreatedAt: restaurantDb.RestaurantCategoryCreatedAt,
				UpdatedAt: restaurantDb.RestaurantCategoryUpdatedAt,
			}

			restaurant.Categories = append(restaurant.Categories, newCategory)
		}
	}

	RespondWithJson(w, http.StatusOK, result)
}

// Get Restaurant Details
func (cfg *ApiConfig) GetRestaurantDetails(w http.ResponseWriter, r *http.Request) {
	// Check Token
	token, err := GetBearerToken(r.Header)
	if err != nil {
		cfg.LogError("Cannot read the authorization token", err)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to get all restaurants")
		return
	}

	_, jwtErr := ValidateJWT(token, cfg.TokenSecret)
	if jwtErr != nil {
		cfg.LogError("Validate jwt failed", jwtErr)
		RespondWithError(w, http.StatusUnauthorized, "You are not authorized to get all restaurants")
		return
	}

	// Path parameter
	restaurantIdStr := r.PathValue("restaurant_id")
	if restaurantIdStr == "" {
		cfg.LogError("Cannot find restaurant_id in path", errors.New("cannot find restaurant id"))
		RespondWithError(w, http.StatusBadRequest, "Restaurant id must be provided.")
		return
	}
	restaurantId, convErr := strconv.Atoi(restaurantIdStr)
	if convErr != nil {
		cfg.LogError("Error converting restaurant_id", convErr)
		RespondWithError(w, http.StatusInternalServerError, "Error converting restaurant id")
		return
	}

	// Fetch Data
	dbResult, dbErr := cfg.Db.GetResturantWithFoodCategoryAndFoodItems(r.Context(), int64(restaurantId))
	if dbErr != nil {
		cfg.LogError("Error fetching restaurant details", dbErr)
		RespondWithError(w, http.StatusNotFound, "Restaurant not found.")
		return
	}

	resRating, parseErr := strconv.ParseFloat(dbResult[0].RestaurantAverageRating, 64)
	if parseErr != nil {
		cfg.LogError("Error parsing average rating", parseErr)
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch restaurant details")
		return
	}

	// Response
	response := restaurantDetailsResponse{
		Id:            dbResult[0].RestaurantID,
		Name:          dbResult[0].RestaurantName,
		ImageUrl:      dbResult[0].RestaurantImageUrl,
		AverageRating: resRating,
		CreatedAt:     dbResult[0].RestaurantCreatedAt,
		UpdatedAt:     dbResult[0].RestaurantUpdatedAt,
	}

	foodCategoriesMap := map[int64]*foodCategoryResponse{}
	foodCategories := []*foodCategoryResponse{}

	// Populate food categories
	for _, singleDbResult := range dbResult {
		foodCategory, ok := foodCategoriesMap[singleDbResult.FoodCategoryID]

		// Parse food item price
		foodItemPrice, parseErr := strconv.ParseFloat(singleDbResult.FoodItemPrice, 64)
		if parseErr != nil {
			cfg.LogError("Error parsing food item price", parseErr)
			RespondWithError(w, http.StatusInternalServerError, "Failed to fetch restaurant details")
			return
		}

		if !ok {
			// Create Category and add food item to it.
			newFoodCategory := foodCategoryResponse{
				Id:        singleDbResult.FoodCategoryID,
				Name:      singleDbResult.FoodCategoryName,
				CreatedAt: singleDbResult.FoodCategoryCreatedAt,
				UpdatedAt: singleDbResult.FoodCategoryUpdatedAt,
				FoodItems: []foodItemResponse{
					{
						Id:          singleDbResult.FoodItemID,
						Name:        singleDbResult.FoodItemName,
						ImageUrl:    singleDbResult.FoodItemImageUrl,
						Description: singleDbResult.FoodItemDescription,
						Price:       foodItemPrice,
						CreatedAt:   singleDbResult.FoodItemCreatedAt,
						UpdatedAt:   singleDbResult.FoodItemUpdatedAt,
					},
				},
			}

			// Add that to map and also foodCategories
			foodCategoriesMap[newFoodCategory.Id] = &newFoodCategory
			foodCategories = append(foodCategories, &newFoodCategory)

		} else {
			// Already found. Add the food item to the found food category
			newFoodItem := foodItemResponse{
				Id:          singleDbResult.FoodItemID,
				Name:        singleDbResult.FoodItemName,
				ImageUrl:    singleDbResult.FoodItemImageUrl,
				Description: singleDbResult.FoodItemDescription,
				Price:       foodItemPrice,
				CreatedAt:   singleDbResult.FoodItemCreatedAt,
				UpdatedAt:   singleDbResult.FoodItemUpdatedAt,
			}
			foodCategory.FoodItems = append(foodCategory.FoodItems, newFoodItem)
		}
	}

	// Add the populated food categories to restaurant
	response.FoodCategories = foodCategories

	// Return the result
	RespondWithJson(w, http.StatusOK, response)
}
