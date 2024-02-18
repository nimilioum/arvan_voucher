DO $$ BEGIN
CREATE TYPE transaction_type as ENUM (
    'deposit',
    'withdraw'
);
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
