SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: trigger_set_update_at_timestamp(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.trigger_set_update_at_timestamp() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
END;
$$;


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: operations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.operations (
    uuid uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone DEFAULT now(),
    created_by integer NOT NULL,
    updated_by integer NOT NULL
);


--
-- Name: policys; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.policys (
    uuid uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    principal_uuid uuid NOT NULL,
    resource_uuid uuid NOT NULL,
    operation_uuid uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone DEFAULT now(),
    created_by integer NOT NULL,
    updated_by integer NOT NULL
);


--
-- Name: principals; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.principals (
    uuid uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone DEFAULT now(),
    created_by integer NOT NULL,
    updated_by integer NOT NULL
);


--
-- Name: resources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.resources (
    uuid uuid DEFAULT public.uuid_generate_v1() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone DEFAULT now(),
    created_by integer NOT NULL,
    updated_by integer NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    sub character varying(50) NOT NULL,
    name character varying(250) NOT NULL,
    given_name character varying(250) NOT NULL,
    family_name character varying(250),
    profile character varying(550),
    picture character varying(550),
    email character varying(350) NOT NULL,
    email_verified boolean DEFAULT false NOT NULL,
    gender character varying(550)
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: operations operations_pkey_; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.operations
    ADD CONSTRAINT operations_pkey_ PRIMARY KEY (uuid);


--
-- Name: policys policy_pkey_; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.policys
    ADD CONSTRAINT policy_pkey_ PRIMARY KEY (uuid);


--
-- Name: principals principals_pkey_; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.principals
    ADD CONSTRAINT principals_pkey_ PRIMARY KEY (uuid);


--
-- Name: resources resources_pkey_; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resources
    ADD CONSTRAINT resources_pkey_ PRIMARY KEY (uuid);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_uniq; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_uniq UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: operations trigger_set_operations_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_set_operations_updated_at BEFORE UPDATE ON public.operations FOR EACH ROW EXECUTE PROCEDURE public.trigger_set_update_at_timestamp();


--
-- Name: policys trigger_set_policys_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_set_policys_updated_at BEFORE UPDATE ON public.policys FOR EACH ROW EXECUTE PROCEDURE public.trigger_set_update_at_timestamp();


--
-- Name: principals trigger_set_principals_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_set_principals_updated_at BEFORE UPDATE ON public.principals FOR EACH ROW EXECUTE PROCEDURE public.trigger_set_update_at_timestamp();


--
-- Name: resources trigger_set_resources_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_set_resources_updated_at BEFORE UPDATE ON public.resources FOR EACH ROW EXECUTE PROCEDURE public.trigger_set_update_at_timestamp();


--
-- Name: users trigger_users_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE PROCEDURE public.trigger_set_update_at_timestamp();


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20190227232621'),
    ('20190401054713');
