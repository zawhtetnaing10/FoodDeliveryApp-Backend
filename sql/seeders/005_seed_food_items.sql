INSERT INTO food_items (name, image_url, description, price, restaurant_id, food_category_id, created_at, updated_at)
VALUES
-- The Round House (restaurant_id = 1)
-- Breakfast (food_category_id = 1)
    ('Monhingar', 'https://media.istockphoto.com/id/2160527391/photo/the-traditional-or-classic-fish-soup-with-white-rice-noodles-from-myanmar-is-called-mohinga.jpg?s=612x612&w=0&k=20&c=LYAPH2qZcphYH6zYDPKEA7F3OPRlMU-B-uvr0kRNXnw=', 'A hearty soup made with rice noodles and a flavorful fish broth, typically featuring catfish', 12.50, 1, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Special Nann Gyi Thoke', 'https://food-touring.com/wp-content/uploads/2014/03/nan-gyi-thoke-2.jpg?w=700&h=467', 'Made with thick round rice noodles mixed with specially prepared chicken curry and chili oil.', 9.75, 1, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Shan Noodles', 'https://www.196flavors.com/wp-content/uploads/2016/10/Shan-noodles-4-FP.jpg', 'A savory noodle dish from the Shan State, featuring noodles with a curried chicken and tomato sauce, peanuts, and a mix of spices.', 8.00, 1, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

-- Set Menus (food_category_id = 2)
    ('The Business Lunch Set', 'https://media.timeout.com/images/106044499/image.jpg', 'A two-course meal with a choice of soup or salad and a main course.', 25.00, 1, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Romantic Dinner for Two', 'https://www.rvwest.com/images/uploads/vdaydinner.jpg', 'Includes a shared appetizer, two main courses, and a dessert to share.', 55.00, 1, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Family Feast', 'https://www.thefrugalfarmer.net/wp-content/uploads/2017/11/Thanksgiving_Family_Feast-1024x614.jpg', 'A selection of signature dishes designed for a group of four.', 75.00, 1, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

-- Drinks (food_category_id = 3)
    ('Classic Mojito', 'https://www.afarmgirlsdabbles.com/wp-content/uploads/2023/06/mojito-afarmgirlsdabbles-01s.jpg', 'Refreshing drink with mint, lime, sugar, and rum.', 9.50, 1, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Fresh Orange Juice', 'https://images-prod.healthline.com/hlcmsresource/images/AN_images/orange-juice-1296x728-feature.jpg', 'Freshly squeezed orange juice, served chilled.', 4.00, 1, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Espresso', 'https://www.thespruceeats.com/thmb/HJrjMfXdLGHbgMhnM0fMkDx9XPQ=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/what-is-espresso-765702-hero-03_cropped-ffbc0c7cf45a46ff846843040c8f370c.jpg', 'A strong shot of espresso.', 3.00, 1, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- Yoon Kitchen (restaurant_id = 2)
-- Lunch (food_category_id = 7)
    ('Spicy Kimchi Fried Rice', 'https://cdn.apartmenttherapy.info/image/upload/f_jpg,q_auto:eco,c_fill,g_auto,w_1500,ar_1:1/k%2FPhoto%2FRecipes%2F2023-12-kimchi-fried-rice%2Fkimchi-fried-rice-197', 'Stir-fried rice with aged kimchi, pork belly, and a fried egg on top.', 14.00, 2, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Bulgogi Bowl', 'https://www.shutterstock.com/image-photo/bulgogi-korean-bbq-beef-fried-600nw-2491516443.jpg', 'Marinated grilled beef served over rice with fresh vegetables.', 16.50, 2, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Jajangmyeon', 'https://www.saveur.com/uploads/2019/12/09/OB7S7ZQJAVDRFGVXYHMD3FUS2Y.jpg?format=auto&optimize=high&width=1440', 'Thick noodles with a black bean sauce, pork, and vegetables.', 13.00, 2, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

-- Dinner (food_category_id = 8)
    ('Korean BBQ Platter', 'https://discoversg.com/wp-content/uploads/sites/32/2017/05/Screenshot-85.png', 'A large platter of marinated meats for grilling at your table.', 45.00, 2, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Army Stew (Budae Jjigae)', 'https://images.food52.com/7f4NvcwkKqH88aaSBPIR8ABMOdQ=/4ef1453b-cd48-4955-9e9b-d9d1263fdb62--2019-0111_army-base-stew_3x2_james-ransom_063.jpeg', 'A hearty stew with sausage, spam, kimchi, tofu, and instant noodles.', 30.00, 2, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Seafood Pancake (Haemul Pajeon)', 'https://koreancuisinerecipes.com/wp-content/uploads/2021/08/haemul-pajeon-seafood-pancake.png', 'A savory pancake with green onions and assorted seafood.', 15.50, 2, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

-- Chinese Food (food_category_id = 5)
    ('Szechuan Chicken', 'https://www.kitchensanctuary.com/wp-content/uploads/2021/05/Szechuan-Chicken-Square-FS-18.jpg', 'Stir-fried chicken with fiery Szechuan peppers and peanuts.', 15.00, 2, 5, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Kung Pao Shrimp', 'https://i.ytimg.com/vi/JKhnGBlWpmY/hq720.jpg?sqp=-oaymwEhCK4FEIIDSFryq4qpAxMIARUAAAAAGAElAADIQj0AgKJD&rs=AOn4CLAUWp-U_YZ3ca6BYP7_WSAYwfZ6-g', 'Succulent shrimp with cashews, bell peppers, and chili sauce.', 18.00, 2, 5, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Mapo Tofu', 'https://ryukoch.b-cdn.net/images/main/mapo-tofu.jpg', 'Soft tofu cubes in a spicy bean sauce with minced pork.', 13.50, 2, 5, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- Ô''Thentic Brasserie (restaurant_id = 3)
-- Lunch (food_category_id = 7)
    ('Quiche Lorraine', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSWHvmHLoI83aIaYhob68Q4r6W18jQjoqgBBg&s', 'Classic French tart with bacon and cheese.', 11.50, 3, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Croque Monsieur', 'https://hips.hearstapps.com/hmg-prod/images/croque-monsieur-66a219aa5f0b2.jpg?crop=1xw:0.8198221757322176xh;center,top&resize=1200:*', 'A toasted ham and cheese sandwich with béchamel sauce.', 12.00, 3, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Salade Niçoise', 'https://www.foodandwine.com/thmb/bkXN0Yi5YrBJzFx--W99t1V0AVs=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Salade-Nicoise-FT-RECIPE0823-c7e3617cc7d4455f90a26fffa7ac128b.jpg', 'A refreshing salad with tuna, green beans, potatoes, and olives.', 14.50, 3, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Dinner (food_category_id = 8)
    ('Duck Confit', 'https://assets.bonappetit.com/photos/5c7ec2040e14053d377dc617/16:9/w_2512,h_1413,c_limit/duck-confit-with-spicy-pickled-raisins.jpg', 'Slow-cooked duck leg with crispy skin.', 26.00, 3, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Beef Bourguignon', 'https://oldcutkitchen.com/wp-content/uploads/2023/12/IMG_7156.jpg', 'Slow-braised beef stew with mushrooms and red wine.', 24.00, 3, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Escargots de Bourgogne', 'https://images.sbs.com.au/dims4/default/628e6ba/2147483647/strip/true/crop/1200x675+0+63/resize/1280x720!/quality/90/?url=http%3A%2F%2Fsbs-au-brightspot.s3.amazonaws.com%2Fdrupal%2Ffood%2Fpublic%2Fimg_7379-snails.jpg', 'Snails baked with garlic, parsley butter, and served with bread.', 18.00, 3, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Steak (food_category_id = 6)
    ('Filet Mignon', 'https://h2qshop.com/cdn/shop/articles/garlic-butter-filet-mignon-973408.jpg?v=1718419589', 'Tenderloin steak with a choice of sauce.', 35.00, 3, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Ribeye Steak', 'https://jesspryles.com/wp-content/uploads/2020/05/brown-garlic-butter-steak-3-scaled.jpg', 'Grilled ribeye with a side of pommes frites.', 32.00, 3, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Sirloin Steak', 'https://www.allrecipes.com/thmb/OJ28fIFte6Pyg93ML8IM-APbu1Y=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/AR-14554-sirloin-steak-with-garlic-butter-hero-4x3-d12fa79836754fcf850388e4677bbf55.jpg', 'A classic sirloin steak cooked to your preference.', 28.00, 3, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- My Ship Floating Restaurant (restaurant_id = 4)
-- Lunch (food_category_id = 7)
    ('Fish and Chips', 'https://upload.wikimedia.org/wikipedia/commons/f/ff/Fish_and_chips_blackpool.jpg', 'Classic battered cod with thick-cut fries and tartar sauce.', 17.50, 4, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Grilled Seabass', 'https://www.allrecipes.com/thmb/6pOT1C_fNN4GwJDU8fDi7tHokBM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/45782-grilled-sea-bass-DDMFS-4x3-01f4cb00dd75490a95fadf81844179a9.jpg', 'Whole grilled seabass with lemon, herbs, and a side of roasted vegetables.', 22.00, 4, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Calamari Rings', 'https://www.allrecipes.com/thmb/u0nGegwZwzLcKcxMaIous2-gVt8=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Crispy-Calamari-Rings-SoupLovingNicole-2000-98de2d857149433ab9949fd1586f9051.jpg', 'Crispy fried calamari served with a spicy aioli dip.', 11.00, 4, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

-- Dinner (food_category_id = 8)
    ('Lobster Thermidor', 'https://upload.wikimedia.org/wikipedia/commons/4/48/Lobster_Thermidor_entree.jpg', 'A rich and creamy lobster dish baked with cheese.', 65.00, 4, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Steak Frites', 'https://www.heinens.com/content/uploads/2025/03/Steak-Frites-with-Chimichurri-1.png', 'Grilled sirloin steak with French fries and a peppercorn sauce.', 28.00, 4, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Grilled Salmon', 'https://www.pccmarkets.com/wp-content/uploads/2017/08/pcc-rosemary-grilled-salmon-flo.jpg', 'A grilled salmon fillet with asparagus and a light lemon butter sauce.', 20.00, 4, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

-- Drinks (food_category_id = 3)
    ('Ocean Breeze Cocktail', 'https://mybartender.com/wp-content/uploads/2025/05/Ocean-Breeze-Margarita-758x423.png', 'A tropical mix of rum, blue curaçao, and pineapple juice.', 10.50, 4, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Strawberry Smoothie', 'https://www.allrecipes.com/thmb/aMLtmuAFr01C66eai_OtGRF0Xu4=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/20792-b-and-ls-strawberry-smoothie-ddmfs-0321-3x4-hero-f9aad20d876448a49a3561bec1da6363.jpg', 'A creamy blend of fresh strawberries and yogurt.', 6.50, 4, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Iced Americano', 'https://images.ctfassets.net/v601h1fyjgba/1vlXSpBbgUo9yLzh71tnOT/a1afdbe54a383d064576b5e628035f04/Iced_Americano.jpg', 'A simple, refreshing blend of espresso and cold water over ice.', 4.50, 4, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- My Hill -PHS (restaurant_id = 5)
-- Breakfast (food_category_id = 1)
    ('Special Mohinga', 'https://media-cdn.tripadvisor.com/media/photo-s/0a/44/2d/38/monhingha-i-do-no-recommend.jpg', 'A traditional Burmese fish noodle soup with a light, flavorful broth.', 7.00, 5, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Ohn No Khao Swe', 'https://tarasmulticulturaltable.com/wp-content/uploads/2015/03/Ohn-No-Khao-Swe-Burmese-Chicken-Coconut-Noodle-Soup-3-of-4.jpg', 'Burmese coconut milk noodle soup with chicken.', 8.50, 5, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Egg Curry with Paratha', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhr_zUKuULMH4q1V4gs66DootVAgdJEh6E-8_7N2xkYEwMuPt2vmqQ7UH4Vnz9tDFd9UhzXCvXlRqlct3Y4NPzNnM1k_ULFe29DJjL09iSq7coGA-9LDrt-u4vHxbcO2KjrkVvm90FRS2U/s2048/9B24E34B-AABC-4D88-A6B7-D4D6377523AD.jpeg', 'Soft boiled eggs in a thick curry sauce, served with flaky bread.', 6.50, 5, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Lunch (food_category_id = 7)
    ('Laphet Thoke', 'https://cdn11.bigcommerce.com/s-jl3t5tg/product_images/uploaded_images/burmese-tea-salad.jpg', 'A tangy and savory fermented tea leaf salad.', 9.00, 5, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Shan Khao Swe', 'https://www.vforveganista.com/wp-content/uploads/2021/08/hero2-edited-2.jpg', 'Rice noodles in a light chicken or pork broth.', 7.50, 5, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Samosa Salad', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS3TvOuaRFYlzoLUyV1JVy-_l45g_sEmqpkUw&s', 'Crushed samosas mixed with vegetables and a sour, savory dressing.', 8.00, 5, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Burmese Food (food_category_id = 4)
    ('Burmese Pork Curry', 'https://boyeatsworld.com.au/wp-content/uploads/2017/04/yangon-food-burmese-pork-curry.jpg', 'A rich and oily pork curry with a deep red color and savory flavor.', 12.00, 5, 4, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Burmese Fried Rice', 'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjxU9tmn6n_i1mYSOlphRQ1Cc2dv2BzHxm9OSasQY6xfBoVVKC8ibs_lKBjca5UDZRBo_06UcDDTOPOCmAUCUh1W6f5tDq9AKYpHd-P7FWdYfINpqJhQyEzJK4w43Ix6gTR_NTcO52QMbWL/s1600/HtaminKyaw.jpg', 'Fried rice with a variety of Burmese vegetables and seasonings.', 10.50, 5, 4, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Wet Thar Hnat', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcShP7AIJjWxoSMkaTXC75zOOtQ6DOHZ_UWW2A&s', 'A traditional Burmese braised pork dish with potatoes.', 11.00, 5, 4, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- Signature (restaurant_id = 6)
-- Lunch (food_category_id = 7)
    ('Signature Burger', 'https://www.sunfedranch.com/wp-content/uploads/SFR-recipe-signaturesauceburger.jpg', 'Our special house-made patty with caramelized onions, cheese, and a secret sauce.', 18.00, 6, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Spaghetti Carbonara', 'https://www.allrecipes.com/thmb/Vg2cRidr2zcYhWGvPD8M18xM_WY=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/11973-spaghetti-carbonara-ii-DDMFS-4x3-6edea51e421e4457ac0c3269f3be5157.jpg', 'Classic pasta dish with creamy egg sauce, pancetta, and black pepper.', 16.50, 6, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Grilled Chicken Sandwich', 'https://thefamilydinnerproject.org/wp-content/uploads/2014/07/Easy-grilled-chicken-sandwich.jpg', 'Juicy grilled chicken breast on a ciabatta bun with fresh greens and aioli.', 15.00, 6, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Dinner (food_category_id = 8)
    ('Pan-Seared Duck Breast', 'https://www.seriouseats.com/thmb/VbNOO-kkYKZmE5RjUCMLZXFZNQE=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__serious_eats__seriouseats.com__2018__01__20180112-duck-breast-vicky-wasik-15-f6ffef51a20f45b19a672ab2acba365b.jpg', 'With a cherry reduction and a side of garlic mashed potatoes.', 28.00, 6, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Beef Wellington', 'https://www.foodandwine.com/thmb/2k2Kq24_fMvHCyLMPRSNrpg5QdE=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/beef-wellington-FT-RECIPE0321-c9a63fccde3b45889ad78fdad078153f.jpg', 'Tender beef fillet coated in pâté and duxelles, wrapped in puff pastry.', 45.00, 6, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Seafood Risotto', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTtYWvwecFqHhMD0DrSrK9_jCcTRVOaivuMjg&s', 'Creamy Italian rice with a medley of shrimp, scallops, and mussels.', 25.00, 6, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Drinks (food_category_id = 3)
    ('Signature Martini', 'https://drinkviacarota.com/cdn/shop/files/signature-martini-100ml-12-pack.jpg-03.jpg?v=1750752757&width=1946', 'A unique blend of vodka, elderflower liqueur, and fresh lime.', 12.00, 6, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Old Fashioned', 'https://www.simplyrecipes.com/thmb/s_de1Nuw4ULiHNECVHOCBY5u5Wk=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2020__01__Old-Fashioned-Cocktail-LEAD-5-1024x681-aa81a798a156453d80d1f7d41de893ff.jpg', 'A timeless cocktail with bourbon, bitters, and a twist of orange.', 11.00, 6, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Red Wine Sangria', 'https://cdn.loveandlemons.com/wp-content/uploads/2021/06/sangria.jpg', 'A refreshing mix of red wine, brandy, and seasonal fruits.', 9.00, 6, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- Yankin Heights Rooftop (restaurant_id = 7)
-- Drinks (food_category_id = 3)
    ('Sunset Spritz', 'https://doubledutchdrinks.com/cdn/shop/articles/Untitled_design_5_6c4faba8-4744-4c00-828b-f80c2c65cb58.jpg?v=1747837792', 'A light and bubbly cocktail with Aperol, prosecco, and soda water.', 10.00, 7, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Whiskey Sour', 'https://www.marthastewart.com/thmb/tbNcHSfi2JSUELHZ6VzkB42KdiU=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/MSL-RM-316574-whiskey-sour-hero-1223-fccd689c7cd944658cbd17c07ac39440.jpg', 'A classic blend of whiskey, lemon juice, sugar, and egg white.', 12.00, 7, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Mango Margarita', 'https://i2.wp.com/www.downshiftology.com/wp-content/uploads/2019/04/Mango-Margarita-main.jpg', 'A tropical twist on a classic, with fresh mango puree and tequila.', 11.50, 7, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Steak (food_category_id = 6)
    ('Porterhouse Steak', 'https://images.getrecipekit.com/20240103175246-broiled-porterhouse-steak-truffle-butter.jpg?width=650&quality=90&', 'A large cut of steak with both the tenderloin and strip loin.', 48.00, 7, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Wagyu Striploin', 'https://craigcookthenaturalbutcher.com.au/cdn/shop/files/craig-cook-wagyu-striploin.jpg?v=1707118119', 'Premium Wagyu beef striploin, known for its marbling and flavor.', 75.00, 7, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('T-Bone Steak', 'https://fireandsmokesociety.com/cdn/shop/articles/Grilled-T-Bone-Steak-with-Steak-King-Seasoning-Fire-and-Smoke-Society_1024x1024.jpg?v=1689775539', 'A T-shaped bone with meat on both sides.', 42.00, 7, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Set Menus (food_category_id = 2)
    ('The Rooftop Tasting Menu', 'https://images.virginexperiencedays.co.uk/images/product/large/nine-course-dinner-tasting-01151248.jpg?auto=compress%2Cformat&w=1440&q=80&fit=max', 'A curated five-course tasting experience of our best dishes.', 90.00, 7, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Light Bites and Cocktails', 'https://media-cdn.tripadvisor.com/media/photo-s/18/9b/29/59/cocktails-and-light-bites.jpg', 'A menu for sharing, with three appetizers and two signature cocktails.', 50.00, 7, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Chef''s Choice', 'https://www.thegate.ca/wp-content/uploads/2022/12/Cheese4Change.jpg', 'Let our chef surprise you with a seasonal three-course meal.', 65.00, 7, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- Omivore by Sharky''s (restaurant_id = 8)
-- Steak (food_category_id = 6)
    ('Dry-Aged Tomahawk', 'https://www.bbquality.nl/wp-content/uploads/2024/09/DSCF1719-1-840x630.jpeg', 'A large bone-in ribeye steak dry-aged for enhanced flavor.', 85.00, 8, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Bone-in Ribeye', 'https://food.fnr.sndimg.com/content/dam/images/food/fullset/2022/05/12/VB1309-valerie-bertinelli-grilled-bone-in-ribeye-with-homemade-herb-brush_4x3.jpg.rend.hgtvcom.1280.960.suffix/1652368000558.webp', 'A juicy and flavorful bone-in cut of ribeye steak.', 40.00, 8, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('New York Strip', 'https://thebigmansworld.com/wp-content/uploads/2023/12/new-york-strip-steak-recipe.jpg', 'A classic cut of steak known for its tenderness and flavor.', 38.00, 8, 6, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Set Menus (food_category_id = 2)
    ('The Meat Lover''s Platter', 'https://kurzweilscountrymeats.com/cdn/shop/products/edited1copy.jpg?v=1675637120', 'A selection of our best steaks and smoked meats.', 120.00, 8, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('The Surf and Turf', 'https://www.piedmontese.com/Content/Images/_Piedmontese/RecipePhotos/pin-Lemon-Garlic-Butter-Surf-n-Turf.jpg', 'A combination of steak and grilled lobster tail.', 75.00, 8, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('The Sharky''s Grill Platter', 'https://media-cdn.tripadvisor.com/media/photo-m/1280/21/84/a2/67/sharky-s-restaurant.jpg', 'Includes a variety of grilled vegetables, smoked sausage, and chicken.', 60.00, 8, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Drinks (food_category_id = 3)
    ('Spicy Margarita', 'https://ichef.bbci.co.uk/food/ic/food_16x9_1600/recipes/spicy_margarita_68770_16x9.jpg', 'A classic margarita with a spicy kick from jalapeños.', 11.50, 8, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Craft Beer Selection', 'https://miamibrewing.com/wp-content/uploads/2022/09/How-to-taste-the-best-craft-beer-like-a-pro.jpg', 'A flight of three local craft beers.', 15.00, 8, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Smoked Old Fashioned', 'https://www.rouses.com/wp-content/uploads/2022/05/Social-Smoked-Old-Fashioned-RC.jpg', 'A classic cocktail with a hint of hickory smoke.', 13.00, 8, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- York Bar (restaurant_id = 9)
-- Drinks (food_category_id = 3)
    ('York Bar Special', 'https://girlabouttravel.co.uk/wp-content/uploads/2025/03/york-cocktails.jpg', 'Our signature cocktail with gin, cucumber, and a secret house-made syrup.', 13.50, 9, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Non-alcoholic Mojito', 'https://www.sustainablecooks.com/wp-content/uploads/2018/06/Classic-Virgin-Mojito-Recipe-2.jpg', 'A refreshing mocktail with mint, lime, and soda water.', 7.00, 9, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Classic Martini', 'https://www.sainsburysmagazine.co.uk/uploads/media/2400x1800/07/7427-Martini.jpg?v=1-0', 'A classic gin or vodka martini, shaken or stirred.', 12.00, 9, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Dinner (food_category_id = 8)
    ('Mini Beef Sliders', 'https://deliciousmadeeasy.com/wp-content/uploads/2019/06/mini-burger-sliders-6-of-7-749x1024.jpg', 'Three mini burgers with cheese, lettuce, and tomato.', 15.00, 9, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Crispy Chicken Wings', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ6JiCUsDDxchOjQyAgxyE7TJ8fF3j6SzppHw&s', 'Fried chicken wings with a choice of BBQ, buffalo, or honey mustard sauce.', 12.00, 9, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Loaded Nachos', 'https://www.allrecipes.com/thmb/wGUnIfJKpAeYrdLL-qCsFDMUBXU=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/7489045loaded-beef-nachos-recipe-lutzflcat4x3-6c7ba4f55c514056894c920446cfd0b2.jpg', 'Crispy tortilla chips topped with melted cheese, jalapeños, and guacamole.', 14.00, 9, 8, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Lunch (food_category_id = 7)
    ('Philly Cheesesteak', 'https://www.spoonforkbacon.com/wp-content/uploads/2025/05/philly-cheesesteak-recipe-card.jpg', 'Thinly sliced beef steak with melted provolone cheese on a hoagie roll.', 16.00, 9, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('BLT Sandwich', 'https://www.allrecipes.com/thmb/3YVAm9kdFKD4SFHPqv30krrvm3A=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/AR-209578-blt-DDMFS-3x4-e1e67e3c348645ba85e304249fd43ff9.jpg', 'A classic sandwich with bacon, lettuce, tomato, and mayonnaise on toasted bread.', 11.00, 9, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Caesar Salad', 'https://cdn.loveandlemons.com/wp-content/uploads/2024/12/caesar-salad.jpg', 'Crisp romaine lettuce with croutons, Parmesan cheese, and Caesar dressing.', 10.00, 9, 7, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- 50th street Restaurant and Bar (restaurant_id = 10)
-- Set Menus (food_category_id = 2)
    ('Classic Western Breakfast', 'https://smokedmeats.com/cdn/shop/files/TWB__60043_1280x_1df208da-7943-4146-a4e4-632d3b3b300c_1280x.jpg?v=1698836065', 'A set including two eggs any style, bacon or sausage, toast, and coffee.', 15.00, 10, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('The Lunch Express', 'https://tb-static.uber.com/prod/enhanced-images/image-touchup-v1/bc7fd550d55a14154d772180a867ee14/ce47d08019a77104fd35b6612239bf5d.jpeg', 'A quick two-course meal with a choice of soup or salad and a main dish like pasta or chicken.', 22.50, 10, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('The Bistro Dinner Set', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTsP6hdEI4pOfHE9ZFFrvd7OU2JaFSMz1CmuQ&s', 'Includes a starter, a main course (steak or fish), and a dessert.', 45.00, 10, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Chinese Food (food_category_id = 5)
    ('Peking Duck', 'https://cravinghomecooked.com/wp-content/uploads/2024/02/peking-duck-1-19.jpg', 'Crispy roasted duck served with thin pancakes, hoisin sauce, and spring onions.', 35.00, 10, 5, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Steamed Pork Buns', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQAjJ2xkBgXEJeP27O7ZMmxKz15M1JZXEEtuw&s', 'Soft, fluffy buns filled with savory minced pork.', 10.00, 10, 5, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Sweet and Sour Pork', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTZLyIq7ZGWy29DO7j3OA7vmHbS1CrvK8F7IA&s', 'Crispy pork pieces tossed in a sweet and tangy sauce with vegetables.', 14.00, 10, 5, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Drinks (food_category_id = 3)
    ('Lychee Martini', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRm7IklT3Y-jz283Ov1qyEG62DrJ3a7FDatjg&s', 'A sweet and fruity martini with lychee liqueur and a splash of lime.', 11.00, 10, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Ginger Ale with Lemon', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQZ_1iY4asY5PSyiw53zCi-EOBmyA6vO-83eQ&s', 'A simple, refreshing drink with ginger ale and a slice of lemon.', 5.00, 10, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Pineapple Juice', 'https://media.diageocms.com/media/rkqfpsp4/barcom_serve_image_1540866_gordonspineapplejuice.jpg', 'Freshly squeezed pineapple juice, served over ice.', 6.00, 10, 3, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),

--------------------------------------------------------------------------------

-- Rangoon Tea House (restaurant_id = 11)
-- Burmese Food (food_category_id = 4)
    ('Burmese Tea Salad', 'https://www.veggiessavetheday.com/wp-content/uploads/2016/04/Burma-Superstar-Tea-Leaf-Salad-new-FI-1200.jpg', 'A famous Burmese dish made with fermented tea leaves, nuts, and garlic.', 9.50, 11, 4, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Burmese Tea with Milk', 'https://thumbs.dreamstime.com/b/cup-myanmar-milk-tea-foam-marble-shop-yangon-burmese-style-126255713.jpg', 'A strong and creamy black tea with condensed milk.', 3.50, 11, 4, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Burmese Rice Salad', 'https://ethnojunkie.com/wp-content/uploads/2020/04/EZJV7166.jpg', 'A spicy and flavorful rice salad with fish paste.', 10.00, 11, 4, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Set Menus (food_category_id = 2)
    ('Afternoon Tea Set', 'https://thedecorkart.com/cdn/shop/articles/10.jpg?v=1650710826', 'A traditional set with a selection of finger sandwiches, pastries, and tea.', 22.00, 11, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Tea House Breakfast Set', 'https://tiptreetearooms.com/cdn/shop/articles/IMG_2524_2048x.jpg?v=1756466193', 'Includes Mohinga, a side of fried tofu, and a cup of Burmese tea.', 15.00, 11, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('The Tasting Platter', 'https://images.squarespace-cdn.com/content/v1/5df2eb230d6bf521e606eb45/1666718139602-OIYOIAB767OSNUIRQA1Y/IMG_5942.jpg?format=2500w', 'A sampler of our most popular Burmese snacks and salads.', 18.00, 11, 2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
-- Breakfast (food_category_id = 1)
    ('Mohinga Bowl', 'https://www.lotusfoods.com/cdn/shop/files/Mohinga_by_crazythickasians_03_800x.jpg?v=1663915613', 'A large bowl of the classic Mohinga noodle soup, perfect for breakfast.', 8.50, 11, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Shan Khao Swe with Chicken', 'https://www.196flavors.com/wp-content/uploads/2016/10/Shan-noodles-4-FP.jpg', 'A popular breakfast noodle dish with a light broth and tender chicken.', 9.00, 11, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC'),
    ('Fried Tofu', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTADNN6dy60C_sv0cvyPTv-JQ0eNqFCKzVazw&s', 'Crispy fried tofu served with a sweet and savory sauce.', 6.00, 11, 1, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC');