
CREATE table vehicle (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "type" varchar(100) NOT NULL,
  "brand" varchar(100) NOT NULL,
  "build_date" varchar(20) NOT NULL,

  "last_longitude" decimal NOT NULL,
  "last_latitude" decimal NOT NULL,

  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

