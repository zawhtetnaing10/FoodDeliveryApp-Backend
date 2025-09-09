-- name: CreateOrder :one
INSERT INTO orders(user_id, delivery_address_id, payment_method_id, total_cost, order_number, created_at, updated_at)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW() AT TIME ZONE 'UTC',
    NOW() AT TIME ZONE 'UTC'
)
RETURNING *;

-- name: GetOrderById :many
SELECT o.id AS order_id,
 o.order_number AS order_number,
 o.total_cost AS order_total_cost,
 o.created_at AS order_created_at,
 o.updated_at AS order_updated_at,
 f.id AS food_item_id,
 f.name AS food_item_name,
 f.image_url AS food_item_image_url,
 f.description AS food_item_description,
 f.price AS food_item_price,
 f.created_at AS food_item_created_at,
 f.updated_at AS food_item_updated_at,
 fho.quantity AS food_item_quantity,
 da.id AS delivery_address_id,
 da.street_address AS delivery_address,
 da.created_at AS delivery_address_created_at,
 da.updated_at AS delivery_address_updated_at,
 pm.id AS payment_method_id,
 pm.card_number AS payment_method_card_number,
 pm.expiry_date AS payment_method_expiry_date,
 pm.cvv AS payment_method_cvv,
 pm.name_on_card AS payment_method_name_on_card,
 pm.created_at AS payment_method_created_at,
 pm.updated_at AS payment_method_updated_at
FROM orders o 
INNER JOIN food_items_has_orders fho
ON fho.order_id = o.id
INNER JOIN food_items f
ON f.id = fho.food_item_id
INNER JOIN delivery_addresses da
ON o.delivery_address_id = da.id
INNER JOIN payment_methods pm
ON o.payment_method_id = pm.id
WHERE o.id = $1;


-- name: GetOrdersForUser :many
SELECT o.id AS order_id,
 o.order_number AS order_number,
 o.total_cost AS order_total_cost,
 o.created_at AS order_created_at,
 o.updated_at AS order_updated_at,
 f.id AS food_item_id,
 f.name AS food_item_name,
 f.image_url AS food_item_image_url,
 f.description AS food_item_description,
 f.price AS food_item_price,
 f.created_at AS food_item_created_at,
 f.updated_at AS food_item_updated_at,
 fho.quantity AS food_item_quantity,
 da.id AS delivery_address_id,
 da.street_address AS delivery_address,
 da.created_at AS delivery_address_created_at,
 da.updated_at AS delivery_address_updated_at,
 pm.id AS payment_method_id,
 pm.card_number AS payment_method_card_number,
 pm.expiry_date AS payment_method_expiry_date,
 pm.cvv AS payment_method_cvv,
 pm.name_on_card AS payment_method_name_on_card,
 pm.created_at AS payment_method_created_at,
 pm.updated_at AS payment_method_updated_at
FROM orders o 
INNER JOIN food_items_has_orders fho
ON fho.order_id = o.id
INNER JOIN food_items f
ON f.id = fho.food_item_id
INNER JOIN delivery_addresses da
ON o.delivery_address_id = da.id
INNER JOIN payment_methods pm
ON o.payment_method_id = pm.id
WHERE o.user_id = $1;

-- name: CalculateTotalCost :one
SELECT SUM(f.price * t.quantity)::numeric
FROM food_items AS f
CROSS JOIN LATERAL
    jsonb_to_recordset(
        sqlc.arg(items)::jsonb
    ) AS t(id INT, quantity INT)
WHERE 
f.id = t.id;