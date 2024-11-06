BEGIN;
DROP TRIGGER IF EXISTS before_insert_linktly_accounts_set_created_at ON linktly.accounts;
DROP TRIGGER IF EXISTS before_update_linktly_accounts_set_updated_at ON linktly.accounts;
DROP TABLE IF EXISTS linktly.accounts;
COMMIT;
