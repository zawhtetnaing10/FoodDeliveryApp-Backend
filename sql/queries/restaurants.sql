-- name: GetAllRestaurants :many
SELECT * FROM restaurants;

-- name: GetAllRestaurantsWithCategories :many
SELECT restaurants.id as restaurant_id,
restaurants.name as restaurant_name,
restaurants.image_url as restaurant_image_url,
restaurants.average_rating as restaurant_average_rating,
restaurants.created_at as restaurant_created_at,
restaurants.updated_at as restaurant_updated_at,
restaurant_categories.id as restaurant_category_id,
restaurant_categories.name as restaurant_category_name,
restaurant_categories.created_at as restaurant_category_created_at,
restaurant_categories.updated_at as restaurant_category_updated_at 
FROM restaurants 
INNER JOIN restaurant_categories_has_restaurants
ON restaurants.id = restaurant_categories_has_restaurants.restaurant_id
INNER JOIN restaurant_categories
ON restaurant_categories.id = restaurant_categories_has_restaurants.restaurant_category_id;
