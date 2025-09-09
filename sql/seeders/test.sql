SELECT o.id AS order_id,
 o.order_number AS order_number,
 
 f.id AS food_item_id,
 f.name AS food_item_name,
 
 fho.quantity AS food_item_quantity,

 da.id AS delivery_address_id,
 pm.id AS payment_method_id,
 pm.card_number AS payment_method_card_number

FROM orders o 
INNER JOIN food_items_has_orders fho
ON fho.order_id = o.id
INNER JOIN food_items f
ON f.id = fho.food_item_id
INNER JOIN delivery_addresses da
ON o.delivery_address_id = da.id
INNER JOIN payment_methods pm
ON o.payment_method_id = pm.id
WHERE o.user_id = 1;