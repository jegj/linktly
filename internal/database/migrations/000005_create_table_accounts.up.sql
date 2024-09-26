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
