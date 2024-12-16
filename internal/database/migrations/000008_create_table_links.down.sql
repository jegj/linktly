BEGIN;
DROP TRIGGER IF EXISTS before_insert_linktly_links_set_created_at ON linktly.links;
DROP TABLE IF EXISTS linktly.links;
COMMIT;
