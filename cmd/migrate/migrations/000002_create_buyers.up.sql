CREATE TABLE IF NOT EXISTS t_organization_buyers (
  id bigserial PRIMARY KEY NOT NULL,
  logo varchar(255),
  name varchar(255) NOT NULL,
  is_active boolean DEFAULT true,
  created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT now()
)
