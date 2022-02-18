CREATE TYPE "gender" AS ENUM (
  'male',
  'female'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" gender NOT NULL,
  "balance" bigint NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "joined_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "drivers" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" gender NOT NULL,
  "balance" bigint NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "current_cab_id" int UNIQUE NOT NULL,
  "joined_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cabs" (
  "id" SERIAL PRIMARY KEY,
  "driver_id" int NOT NULL,
  "brand" varchar NOT NULL,
  "model" varchar NOT NULL,
  "color" varchar NOT NULL,
  "plate" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "trips" (
  "id" SERIAL PRIMARY KEY,
  "origin" varchar NOT NULL,
  "destination" varchar NOT NULL,
  "rider_id" int NOT NULL,
  "driver_id" int NOT NULL,
  "start_time" timestamptz NOT NULL DEFAULT (now()),
  "end_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "status" int NOT NULL DEFAULT 1,
  "amount" bigint NOT NULL,
  "cab_id" int NOT NULL,
  "driver_rating" int NOT NULL DEFAULT 0
);

ALTER TABLE "drivers" ADD FOREIGN KEY ("current_cab_id") REFERENCES "cabs" ("id");

ALTER TABLE "cabs" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("rider_id") REFERENCES "users" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("cab_id") REFERENCES "cabs" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "drivers" ("id");

CREATE INDEX ON "drivers" ("username");

CREATE INDEX ON "cabs" ("id");

CREATE INDEX ON "cabs" ("driver_id");

CREATE INDEX ON "cabs" ("plate");

CREATE INDEX ON "trips" ("id");

CREATE INDEX ON "trips" ("rider_id");

CREATE INDEX ON "trips" ("driver_id");

CREATE INDEX ON "trips" ("rider_id", "driver_id");

CREATE INDEX ON "trips" ("cab_id");

CREATE INDEX ON "trips" ("origin");

CREATE INDEX ON "trips" ("destination");

CREATE INDEX ON "trips" ("origin", "destination");

COMMENT ON COLUMN "users"."balance" IS 'must be positive';

ALTER TABLE public.drivers DISABLE TRIGGER ALL;

INSERT INTO public.drivers
(id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at)
VALUES(0, 'owner', 'password', 'owner', 'male', 0, 'owner@simpleUber.com', 0, now());

ALTER TABLE public.drivers ENABLE TRIGGER ALL;

INSERT INTO public.cabs
(id, driver_id, brand, model, color, plate, created_at)
VALUES(0, 0, 'aston martin', 'rapid s', 'blcak', '00A000-00', now());