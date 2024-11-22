-- Step 1: Create necessary enum types
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
  "address" varchar,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE "vendors" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "biography" text,
  "profile_picture" varchar,
  "active" bool DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE "customers" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid
);

CREATE TABLE "appointments" (
  "id" uuid PRIMARY KEY,
  "customer_id" uuid NOT NULL,
  "vendor_id" uuid NOT NULL,
  "date" TIMESTAMP(3) NOT NULL,
  "time_slot_id" uuid NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
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
  "vendor_id" uuid NOT NULL,
  "day_of_week" varchar NOT NULL,
  "date" date
);

CREATE TABLE "feedback" (
  "id" uuid PRIMARY KEY,
  "appointment_id" uuid,
  "rating" int,
  "comment" text,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE "services" (
  "id" uuid PRIMARY KEY,
  "vendor_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "description" text,
  "price" decimal(10,2),
  "duration" interval,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
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
  "transaction_id" varchar,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
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
