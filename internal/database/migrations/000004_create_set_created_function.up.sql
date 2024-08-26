BEGIN;
CREATE OR REPLACE FUNCTION linktly.set_created_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at := uuid_v8_to_timestamptz(NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
COMMIT;
