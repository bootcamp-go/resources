-- DCL (Data Control Language)
-- drop user
DROP USER IF EXISTS 'melisprint_user'@'localhost';

-- create user 
CREATE USER 'melisprint_user'@'localhost' IDENTIFIED BY 'melisprint_pass';

-- grant privileges
GRANT ALL PRIVILEGES ON melisprint.* TO 'melisprint_user'@'localhost';

-- flush privileges
FLUSH PRIVILEGES;