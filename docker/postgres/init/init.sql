CREATE DATABASE IF NOT EXISTS "Gofre"
WITH
    OWNER = postgres ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8' LOCALE_PROVIDER = 'libc' TABLESPACE = pg_default CONNECTION
LIMIT
    = -1 IS_TEMPLATE = False;

-- auth micro service
CREATE SCHEMA IF NOT EXISTS auth AUTHORIZATION postgres;

CREATE TABLE
    IF NOT EXISTS auth.reset_tokens (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY (
            INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1
        ),
        user_id integer NOT NULL,
        hash_token character varying(255) COLLATE pg_catalog."default" NOT NULL,
        expires_at timestamp
        with
            time zone NOT NULL,
            CONSTRAINT reset_tokens_pkey PRIMARY KEY (id),
            CONSTRAINT reset_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES auth.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
    )
CREATE TABLE
    IF NOT EXISTS auth.users (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY (
            INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1
        ),
        name character varying(255) COLLATE pg_catalog."default" NOT NULL,
        username character varying(255) COLLATE pg_catalog."default" NOT NULL,
        last_name character varying(255) COLLATE pg_catalog."default",
        cell_phone character varying(255) COLLATE pg_catalog."default",
        email character varying(255) COLLATE pg_catalog."default" NOT NULL,
        password character varying(255) COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp
        with
            time zone NOT NULL,
            updated_at timestamp
        with
            time zone NOT NULL,
            CONSTRAINT users_pkey PRIMARY KEY (id)
    )
-- categories micro service






