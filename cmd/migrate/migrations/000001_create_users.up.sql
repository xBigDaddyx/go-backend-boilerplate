CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS t_organization_users (
  id bigserial PRIMARY KEY,
	name varchar(255) NOT NULL,
	email citext UNIQUE NOT NULL,
  password bytea NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT now()
)
