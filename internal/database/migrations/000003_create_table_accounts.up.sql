BEGIN;
CREATE TABLE linktly.accounts (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255),
   lastname VARCHAR(255),
   email VARCHAR(255) UNIQUE NOT NULL,
   password VARCHAR(255),
   api_token VARCHAR(255) DEFAULT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   CONSTRAINT email_check CHECK (
     email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
   )
);

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE linktly.accounts TO linktly_user;
COMMIT;
