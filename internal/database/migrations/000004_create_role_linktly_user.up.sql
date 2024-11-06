CREATE ROLE linktly_user LOGIN;
ALTER ROLE linktly_user SET search_path TO public,linktly;
GRANT USAGE ON SCHEMA linktly TO linktly_user;
GRANT ALL ON SCHEMA linktly TO linktly_user;
