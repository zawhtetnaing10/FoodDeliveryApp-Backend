package handlers

import (
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
