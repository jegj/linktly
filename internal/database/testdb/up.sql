-- Start of internal/database/migrations/000001_create_linktly_schema.up.sql
CREATE SCHEMA IF NOT EXISTS linktly AUTHORIZATION linktly_admin;

-- End of internal/database/migrations/000001_create_linktly_schema.up.sql

-- Start of internal/database/migrations/000002_create_extension_pg_uuidv7.up.sql
CREATE EXTENSION IF NOT EXISTS pg_uuidv7;

-- End of internal/database/migrations/000002_create_extension_pg_uuidv7.up.sql

-- Start of internal/database/migrations/000003_create_role_linktly_user.up.sql
CREATE ROLE linktly_user LOGIN;
ALTER ROLE linktly_user SET search_path TO public,linktly;
GRANT USAGE ON SCHEMA linktly TO linktly_user;

-- End of internal/database/migrations/000003_create_role_linktly_user.up.sql

-- Start of internal/database/migrations/000004_create_set_created_function.up.sql
BEGIN;
CREATE OR REPLACE FUNCTION linktly.set_created_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at := uuid_v7_to_timestamptz(NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
COMMIT;

-- End of internal/database/migrations/000004_create_set_created_function.up.sql

-- Start of internal/database/migrations/000005_create_table_accounts.up.sql
BEGIN;
CREATE TABLE IF NOT EXISTS linktly.accounts (
   id  UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v7(),
   name VARCHAR(255) NOT NULL,
   lastname VARCHAR(255) NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   password VARCHAR(255),
   api_token VARCHAR(255) DEFAULT NULL,
   role INT NOT null DEFAULT 2,
   refresh_token_jti VARCHAR(255) DEFAULT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP DEFAULT NULL,
   CONSTRAINT email_check CHECK (
     email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
   )
);

COMMENT ON COLUMN linktly.accounts.id is 'To get created_at use uuid_v7_to_timestamptz(id)';
COMMENT ON COLUMN linktly.accounts.role is '1->admin, 2->user, 3->guest';

CREATE OR REPLACE TRIGGER before_insert_linktly_accounts_set_created_at
BEFORE INSERT ON linktly.accounts 
FOR EACH ROW
EXECUTE FUNCTION linktly.set_created_at();

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE linktly.accounts TO linktly_user;
COMMIT;

-- End of internal/database/migrations/000005_create_table_accounts.up.sql

-- Start of internal/database/migrations/000006_create_table_folders.up.sql
BEGIN;
CREATE TABLE IF NOT EXISTS linktly.folders (
   id  UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v7(),
   name VARCHAR(255) NOT NULL,
   description TEXT,
   account_id  UUID REFERENCES linktly.accounts(id) NOT NULL,
   parent_folder_id UUID REFERENCES linktly.folders(id) DEFAULT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP DEFAULT NULL,
   undeletable BOOLEAN DEFAULT FALSE
);

COMMENT ON COLUMN linktly.folders.id is 'To get created_at use uuid_v7_to_timestamptz(id)';

CREATE OR REPLACE TRIGGER before_insert_linktly_folders_set_created_at
BEFORE INSERT ON linktly.folders 
FOR EACH ROW
EXECUTE FUNCTION linktly.set_created_at();

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE linktly.folders TO linktly_user;
COMMIT;


-- End of internal/database/migrations/000006_create_table_folders.up.sql

-- Start of internal/database/migrations/000007_create_table_links.up.sql
BEGIN;
CREATE TABLE IF NOT EXISTS linktly.links (
   id  UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v7(),
   name VARCHAR(255) NOT NULL,
   linktly_url TEXT NOT NULL,
   url TEXT NOT NULL,
   description TEXT,
   account_id  UUID REFERENCES linktly.accounts(id) NOT NULL,
   folder_id   UUID REFERENCES linktly.folders(id) NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP DEFAULT NULL,
   expires_at TIMESTAMP DEFAULT NULL
);

COMMENT ON COLUMN linktly.links.id is 'To get created_at use uuid_v7_to_timestamptz(id)';

CREATE OR REPLACE TRIGGER before_insert_linktly_links_set_created_at
BEFORE INSERT ON linktly.links 
FOR EACH ROW
EXECUTE FUNCTION linktly.set_created_at();

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE linktly.links TO linktly_user;
COMMIT;



-- End of internal/database/migrations/000007_create_table_links.up.sql


-- Starting dummy data 

-- Start of internal/database/testdb/dummy_data/auth_dummy.sql
INSERT INTO linktly.accounts(
	id, name, lastname, email, password, api_token, role, created_at, updated_at, refresh_token_jti)
VALUES ('0191e400-fe29-7434-9c6c-26fbc133ecd1', 'Javier', 'Galarza', 'jegj57@gmail.com', '$2a$15$Qk1tLiWfXhTNbTJ50huFcuNrLS1CKjQ5NNxv6zoMZihtEncjJU4Lu', NULL, 2, '2024-09-12 02:12:36.009'::timestamp, NULL, NULL);

-- End of internal/database/testdb/dummy_data/auth_dummy.sql

