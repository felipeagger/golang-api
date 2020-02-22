CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	username varchar(150) NOT NULL,
	age int4 NULL,
	bio varchar(250) NULL,
	link varchar(500) NULL,
	avatar varchar(500) NULL,
	created_at timestamp NULL DEFAULT now(),
	score numeric(15,2) NULL,
	CONSTRAINT users_id_pk PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.users OWNER TO postgres;
GRANT ALL ON TABLE public.users TO postgres;
