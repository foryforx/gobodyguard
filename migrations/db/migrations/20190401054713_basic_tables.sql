-- migrate:up

-- Function to automatically set update_at field
CREATE OR REPLACE FUNCTION trigger_set_update_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- principals table
CREATE table principals (
	uuid  uuid NOT NULL DEFAULT uuid_generate_v1(),  
        CONSTRAINT principals_pkey_ PRIMARY KEY (uuid),
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone NULL,
    created_by text NOT NULL, 
    updated_by text NOT NULL
);

CREATE TRIGGER trigger_set_principals_updated_at
BEFORE UPDATE ON principals
FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_at_timestamp();

-- resources table
CREATE table resources (
	uuid  uuid NOT NULL DEFAULT uuid_generate_v1(),  
        CONSTRAINT resources_pkey_ PRIMARY KEY (uuid),
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone NULL,
    created_by text NOT NULL, 
    updated_by text NOT NULL
);

CREATE TRIGGER trigger_set_resources_updated_at
BEFORE UPDATE ON resources
FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_at_timestamp();

-- operations table
CREATE table operations (
	uuid  uuid NOT NULL DEFAULT uuid_generate_v1(),  
        CONSTRAINT operations_pkey_ PRIMARY KEY (uuid),
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone NULL,
    created_by text NOT NULL, 
    updated_by text NOT NULL
);

CREATE TRIGGER trigger_set_operations_updated_at
BEFORE UPDATE ON operations
FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_at_timestamp();

-- policies table
CREATE table policys (
	uuid  uuid NOT NULL DEFAULT uuid_generate_v1(),  
        CONSTRAINT policy_pkey_ PRIMARY KEY (uuid),
    principal_uuid uuid NOT NULL,
    resource_uuid uuid NOT NULL,
    operation_uuid uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone NULL,
    created_by text NOT NULL, 
    updated_by text NOT NULL
);

CREATE TRIGGER trigger_set_policys_updated_at
BEFORE UPDATE ON policys
FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_at_timestamp();

-- migrate:down
DROP TABLE policys;
DROP TRIGGER IF EXISTS trigger_set_policies_updated_at ON policys;
DROP TABLE operations;
DROP TRIGGER IF EXISTS trigger_set_operations_updated_at ON operations;
DROP TABLE resources;
DROP TRIGGER IF EXISTS trigger_set_resources_updated_at ON resources;
DROP TABLE principals;
DROP TRIGGER IF EXISTS trigger_set_principals_updated_at ON principals;
-- DROP FUNCTION IF EXISTS trigger_set_update_at_timestamp;
