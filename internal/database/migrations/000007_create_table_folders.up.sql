BEGIN;
CREATE TABLE IF NOT EXISTS linktly.folders (
   id  UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v7(),
   name VARCHAR(255) NOT NULL,
   description TEXT,
   account_id  UUID REFERENCES linktly.accounts(id) ON DELETE CASCADE ON UPDATE NO ACTION,
   parent_folder_id UUID REFERENCES linktly.folders(id) ON DELETE CASCADE ON UPDATE NO ACTION DEFAULT NULL,
   created_at TIMESTAMP WITH TIME ZONE,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
   undeletable BOOLEAN DEFAULT FALSE
);

COMMENT ON COLUMN linktly.folders.id is 'To get created_at use uuid_v7_to_timestamptz(id)';

CREATE OR REPLACE TRIGGER before_insert_linktly_folders_set_created_at
BEFORE INSERT ON linktly.folders 
FOR EACH ROW
EXECUTE FUNCTION linktly.set_created_at();

CREATE OR REPLACE TRIGGER before_update_linktly_folders_set_updated_at
BEFORE UPDATE ON linktly.folders
FOR EACH ROW
EXECUTE FUNCTION linktly.set_updated_at();

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE linktly.folders TO linktly_user;
COMMIT;
