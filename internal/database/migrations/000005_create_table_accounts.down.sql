BEGIN;
DROP TRIGGER IF EXISTS before_insert_linktly_accounts_set_created_at ON linktly.accounts;
DROP FUNCTION IF EXISTS linktly.set_created_at();
DROP TABLE IF EXISTS linktly.accounts;
COMMIT;
