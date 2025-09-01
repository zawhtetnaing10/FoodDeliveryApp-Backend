package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type restaurantResponse struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	ImageUrl      string    `json:"image_url"`
	AverageRating float64   `json:"average_rating"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Get All Restaurants
func (cfg *ApiConfig) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
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

	restaurantsFromDb, getRestaurantsErr := cfg.Db.GetAllRestaurants(r.Context())
	if getRestaurantsErr != nil {
		cfg.LogError("Get restaurants failed", getRestaurantsErr)
		RespondWithError(w, http.StatusInternalServerError, "There's something wrong while fetching the restaurants. Please try again.")
		return
	}

	response := []restaurantResponse{}
	for _, dbRestaurant := range restaurantsFromDb {

		avgRating, parseErr := strconv.ParseFloat(dbRestaurant.AverageRating, 64)
		if parseErr != nil {
			cfg.LogError(fmt.Sprintf("Cannot parse average rating: %v", dbRestaurant.AverageRating), parseErr)
			RespondWithError(w, http.StatusInternalServerError, "There was something wrong while fetching the restaurants. Please try again.")
			return
		}

		restaurantResponse := restaurantResponse{
			Id:            dbRestaurant.ID,
			Name:          dbRestaurant.Name,
			ImageUrl:      dbRestaurant.ImageUrl,
			AverageRating: avgRating,
			CreatedAt:     dbRestaurant.CreatedAt,
			UpdatedAt:     dbRestaurant.UpdatedAt,
		}

		response = append(response, restaurantResponse)
	}

	RespondWithJson(w, http.StatusOK, response)
}
