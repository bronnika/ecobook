CREATE TABLE "user" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "phone" varchar,
  "password" varchar,
  "rating" integer,
  "create_date" timestamp with time zone default current_timestamp
);

CREATE TABLE "otp" (
  "id" serial PRIMARY KEY,
  "otp_code" integer,
  "phone" varchar,
  "exp_date" timestamp with time zone,
  "is_used" boolean,
  "create_date" timestamp with time zone default current_timestamp
);

CREATE TABLE "product_category" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "code" varchar,
  "photo" varchar
);

CREATE TABLE "product" (
  "id" serial PRIMARY KEY,
  "category_id" integer,
  "name" varchar,
  "photo" varchar,
  "description" varchar,
  "user_id" integer,
  "sale_type_id" integer,
  "price" numeric,
  "is_active" boolean,
  "create_date" timestamp with time zone default current_timestamp
);

CREATE TABLE "conversation" (
  "id" serial PRIMARY KEY,
  "product_id" integer,
  "user_seller_id" integer,
  "user_buyer_id" integer,
  "create_date" timestamp with time zone default current_timestamp
);

CREATE TABLE "message" (
  "id" serial PRIMARY KEY,
  "user_id" integer,
  "conversation_id" integer,
  "message" text,
  "create_date" timestamp with time zone default current_timestamp
);

CREATE TABLE "sale_type" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "code" varchar
);

CREATE TABLE "point_category" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "code" varchar,
  "photo" varchar
);

CREATE TABLE "utilize_point" (
  "id" serial PRIMARY KEY,
  "category_id" integer,
  "name" varchar,
  "description" varchar,
  "address" varchar,
  "phone" varchar,
  "latitude" varchar,
  "longitude" varchar,
  "photo" varchar
);

CREATE TABLE "news_type" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "code" varchar
);

CREATE TABLE "news" (
  "id" serial PRIMARY KEY,
  "title" varchar,
  "text" text,
  "photo" varchar,
  "event_date" timestamp with time zone default current_timestamp,
  "news_type_id" integer,
  "create_date" timestamp with time zone default current_timestamp
);

CREATE TABLE "participant" (
  "news_id" integer,
  "user_id" integer
);

ALTER TABLE "conversation" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("conversation_id") REFERENCES "conversation" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "product_category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("sale_type_id") REFERENCES "sale_type" ("id");

ALTER TABLE "utilize_point" ADD FOREIGN KEY ("category_id") REFERENCES "point_category" ("id");

ALTER TABLE "news" ADD FOREIGN KEY ("news_type_id") REFERENCES "news_type" ("id");

ALTER TABLE "participant" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "participant" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
