-- Step 1: Create necessary enum types
DO $$ BEGIN
    CREATE TYPE appointment_status_enum AS ENUM ('scheduled', 'completed', 'canceled');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE day_of_week_enum AS ENUM ('Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE payment_status_enum AS ENUM ('pending', 'completed', 'failed');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE payment_method_enum AS ENUM ('credit_card', 'paypal', 'bank_transfer');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Continue with table creation
-- Step 1: Create the users table
CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,             -- Unique identifier for each user
    "name" varchar NOT NULL,           -- Name of the user
    "email" varchar UNIQUE NOT NULL,   -- Unique email for each user
    "password" varchar NOT NULL,       -- Password (hashed) for authentication
    "user_type" varchar NOT NULL,  -- Type of user (e.g., admin, vendor, customer)
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Timestamp for when the user was created
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  -- Timestamp for when the user was last updated
);


CREATE TABLE "contacts" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "phone" varchar UNIQUE,
  "address" varchar
);

CREATE TABLE "vendors" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "biography" text,
  "profile_picture" varchar,
  "active" bool DEFAULT false
);

CREATE TABLE "customers" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid
);

CREATE TABLE "appointments" (
  "id" uuid PRIMARY KEY,
  "customer_id" uuid,
  "vendor_id" uuid,
  "date" TIMESTAMP(3) NOT NULL,
  "time_slot_id" uuid,
  "status" appointment_status_enum
);

CREATE TABLE "time_slots" (
  "id" uuid PRIMARY KEY,
  "vendor_id" uuid,
  "start_time" timestamp,
  "end_time" timestamp,
  "is_booked" boolean,
  "buffer_time" interval
);

CREATE TABLE "vendor_availability" (
  "id" uuid PRIMARY KEY,
  "vendor_id" uuid,
  "day_of_week" day_of_week_enum,
  "date" date
);

CREATE TABLE "feedback" (
  "id" uuid PRIMARY KEY,
  "appointment_id" uuid,
  "rating" int,
  "comment" text
);

CREATE TABLE "services" (
  "id" uuid PRIMARY KEY,
  "vendor_id" uuid,
  "name" varchar,
  "description" text,
  "price" decimal(10,2),
  "duration" interval,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "payments" (
  "id" uuid PRIMARY KEY,
  "appointment_id" uuid,
  "customer_id" uuid,
  "vendor_id" uuid,
  "amount" decimal(10,2),
  "payment_method" payment_method_enum,
  "status" payment_status_enum,
  "payment_date" timestamp,
  "transaction_id" varchar
);

-- Step 2: Create foreign keys
ALTER TABLE "contacts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "vendors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "appointments" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "appointments" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "appointments" ADD FOREIGN KEY ("time_slot_id") REFERENCES "time_slots" ("id");

ALTER TABLE "time_slots" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "vendor_availability" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "feedback" ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");
