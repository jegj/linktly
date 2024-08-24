CREATE ROLE linktly_user LOGIN;
ALTER ROLE linktly_user SET search_path TO linktly;
GRANT USAGE ON SCHEMA linktly TO linktly_user;

