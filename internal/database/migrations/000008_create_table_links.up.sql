BEGIN;
CREATE TABLE IF NOT EXISTS linktly.links (
   id  UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v7(),
   name VARCHAR(255) NOT NULL,
   linktly_code VARCHAR(255) UNIQUE NOT NULL,
   url TEXT NOT NULL,
   description TEXT,
   account_id  UUID REFERENCES linktly.accounts(id) NOT NULL,
   folder_id   UUID REFERENCES linktly.folders(id) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
   expires_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

COMMENT ON COLUMN linktly.links.id is 'To get created_at use uuid_v7_to_timestamptz(id)';

CREATE OR REPLACE TRIGGER before_insert_linktly_links_set_created_at
BEFORE INSERT ON linktly.links 
FOR EACH ROW
EXECUTE FUNCTION linktly.set_created_at();

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE linktly.links TO linktly_user;
COMMIT;


