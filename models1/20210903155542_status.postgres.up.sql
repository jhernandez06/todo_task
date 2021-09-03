ALTER TABLE users RENAME COLUMN active TO status_user;

ALTER TABLE users
ALTER COLUMN active TYPE VARCHAR;
UPDATE users SET status_user = 'active' 
WHERE rol = 'admin';

