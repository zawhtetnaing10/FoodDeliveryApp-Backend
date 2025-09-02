SELECT r.id as restaurant_id, 
    r.name as restaurant_name,
    r.image_url as restaurant_image_url,
    r.average_rating as restaurant_average_rating, 
    r.created_at as restaurant_created_at,
    r.updated_at as restaurant_updated_at,
    fi.id as food_item_id, 
    fi.name as food_item_name, 
    fi.image_url as food_item_image_url,
    fi.description as food_item_description,
    fi.price as food_item_price,
    fi.created_at as food_item_created_at,
    fi.updated_at as food_item_updated_at,
    fc.id as food_category_id, 
    fc.name as food_category_name,
    fc.created_at as food_category_created_at,
    fc.updated_at as food_category_updated_at
FROM restaurants r
INNER JOIN food_items fi
ON r.id = fi.restaurant_id
INNER JOIN food_categories fc
ON fc.id = fi.food_category_id
WHERE r.id = $1;