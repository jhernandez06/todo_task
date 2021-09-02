INSERT INTO users 
(id,first_name,last_name,email,active,rol,password_hash,created_at,updated_at) 
VALUES 
(uuid_generate_v1(),'Javier', 'Hernandez','jhernandez@wawand.co','true','admin',crypt('javier', gen_salt('bf')), NOW(), NOW());