BEGIN;
DROP TRIGGER IF EXISTS before_insert_linktly_folders_set_created_at ON linktly.folders;
DROP TABLE IF EXISTS linktly.folders;
COMMIT;

