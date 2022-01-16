CREATE TYPE "flight_status" AS ENUM (
  'ACTIVE',
  'SCHEDULED',
  'DELAYED',
  'DEPARTED',
  'ARRIVED',
  'CANCELED'
);

CREATE TYPE "reservation_status" AS ENUM (
  'REQUESTED',
  'PENDING',
  'CONFIRMED',
  'CANCELED',
  'CHECKED_IN',
  'ABANDONED'
);

CREATE TYPE "seat_class" AS ENUM (
  'ECONOMY',
  'ECONOMY_PLUS',
  'BUSINESS',
  'FIRST'
);

CREATE TYPE "seat_type" AS ENUM (
  'REGULAR',
  'EMERGENCY_EXIT',
  'ACCESSIBLE'
);

CREATE TYPE "account_status" AS ENUM (
  'ACTIVE',
  'CLOSED',
  'BLACKLISTED',
  'BLOCKED',
  'CANCELED'
);

CREATE TYPE "trip_type" AS ENUM (
  'ONE_WAY',
  'ROUND_TRIP'
);

CREATE TYPE "boarding_policy" AS ENUM (
  'ZONE_BASED',
  'GROUP_BASED'
);

CREATE TABLE "flights" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "flight_number" varchar,
  "boarding_policy" boarding_policy,
  "duration" integer,
  "distance" integer,
  "departure_airport" airports,
  "arrival_airport" airports
);

CREATE TABLE "airports" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar,
  "iata_code" varchar,
  "icao_code" varchar,
  "address" addresses
);

CREATE TABLE "airlines" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar,
  "iata_code" varchar
);

CREATE TABLE "aircrafts" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "tail_number" varchar,
  "manufacturer" varchar,
  "model" varchar,
  "capacity" integer,
  "range" integer,
  "airline_id" integer,
  "aircraft_type" varchar
);

CREATE TABLE "flight_instances" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "departure_gate" integer,
  "arrival_gate" integer,
  "departs_at" datetime,
  "arrives_at" datetime,
  "flight_status" flight_status,
  "flight_id" integer NOT NULL,
  "aircraft_id" integer NOT NULL
);

CREATE TABLE "flight_reservations" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "reservation_number" varchar,
  "reservation_status" reservation_status,
  "passengers" passengers[],
  "flight_instance_id" integer NOT NULL,
  "itenerary_id" integer NOT NULL
);

CREATE TABLE "iteneraries" (
  "id" integer,
  "origin" airports,
  "destination" airports,
  "airport_id" integer UNIQUE NOT NULL,
  "customer_id" integer UNIQUE NOT NULL
);

CREATE TABLE "accounts" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "password" varchar,
  "account_status" account_status,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "users" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "firstname" varchar,
  "lastname" varchar,
  "email" varchar,
  "phone" varchar,
  "address" addresses,
  "role_id" integer UNIQUE NOT NULL
);

CREATE TABLE "admins" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "password" varchar,
  "two_fa_secret" varchar,
  "two_fa_completed" boolean,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "crews" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "employee_id" varchar,
  "airline_id" integer UNIQUE NOT NULL,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "pilots" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "licence_number" varchar,
  "flight_hours" integer,
  "employee_id" varchar,
  "airline_id" integer UNIQUE NOT NULL,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "front_desks" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "employee_id" varchar,
  "airport_id" integer UNIQUE NOT NULL,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "addresses" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "street" varchar,
  "city" varchar,
  "state" varchar,
  "zipcode" varchar,
  "airport_id" integer UNIQUE NOT NULL,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "customers" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "address" addresses,
  "frequent_flyer_number" varchar,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "flights_airports" (
  "flight_id" integer NOT NULL,
  "airport_id" integer NOT NULL,
  PRIMARY KEY ("flight_id", "airport_id")
);

CREATE TABLE "passengers" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "firstname" varchar,
  "lastname" varchar,
  "age" integer,
  "passport_number" varchar,
  "flight_reservation_id" integer UNIQUE NOT NULL,
  "flight_seat_id" integer UNIQUE NOT NULL
);

CREATE TABLE "seats" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "seat_number" integer,
  "seat_row" varchar,
  "seat_type" seat_type,
  "seat_class" seat_class,
  "aircraft_id" integer UNIQUE NOT NULL
);

CREATE TABLE "flight_seats" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "flight_instance_id" integer UNIQUE NOT NULL,
  "seat_id" integer UNIQUE NOT NULL
);

CREATE TABLE "roles" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar
);

CREATE TABLE "permissions" (
  "id" integer UNIQUE PRIMARY KEY NOT NULL,
  "permission" varchar,
  "role_id" integer
);

CREATE TABLE "flight_crew" (
  "flight_id" integer UNIQUE NOT NULL,
  "crew_id" integer UNIQUE NOT NULL
);

ALTER TABLE "flight_instances" ADD FOREIGN KEY ("flight_id") REFERENCES "flights" ("id");

ALTER TABLE "aircrafts" ADD FOREIGN KEY ("airline_id") REFERENCES "airlines" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("id") REFERENCES "flights_airports" ("flight_id");

ALTER TABLE "airports" ADD FOREIGN KEY ("id") REFERENCES "flights_airports" ("airport_id");

ALTER TABLE "aircrafts" ADD FOREIGN KEY ("id") REFERENCES "flight_instances" ("aircraft_id");

ALTER TABLE "flight_reservations" ADD FOREIGN KEY ("flight_instance_id") REFERENCES "flight_instances" ("id");

ALTER TABLE "flight_reservations" ADD FOREIGN KEY ("itenerary_id") REFERENCES "iteneraries" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "accounts" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "admins" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "crews" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "pilots" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "front_desks" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "customers" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "addresses" ("user_id");

ALTER TABLE "airports" ADD FOREIGN KEY ("id") REFERENCES "addresses" ("airport_id");

ALTER TABLE "iteneraries" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "seats" ADD FOREIGN KEY ("aircraft_id") REFERENCES "aircrafts" ("id");

ALTER TABLE "flight_seats" ADD FOREIGN KEY ("flight_instance_id") REFERENCES "flight_instances" ("id");

ALTER TABLE "seats" ADD FOREIGN KEY ("id") REFERENCES "flight_seats" ("seat_id");

ALTER TABLE "passengers" ADD FOREIGN KEY ("flight_reservation_id") REFERENCES "flight_reservations" ("id");

ALTER TABLE "flight_seats" ADD FOREIGN KEY ("id") REFERENCES "passengers" ("flight_seat_id");

ALTER TABLE "airports" ADD FOREIGN KEY ("id") REFERENCES "iteneraries" ("airport_id");

ALTER TABLE "permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "flights" ADD FOREIGN KEY ("id") REFERENCES "flight_crew" ("flight_id");

ALTER TABLE "crews" ADD FOREIGN KEY ("id") REFERENCES "flight_crew" ("crew_id");

ALTER TABLE "crews" ADD FOREIGN KEY ("airline_id") REFERENCES "airlines" ("id");

ALTER TABLE "pilots" ADD FOREIGN KEY ("airline_id") REFERENCES "airlines" ("id");

ALTER TABLE "front_desks" ADD FOREIGN KEY ("airport_id") REFERENCES "airports" ("id");

