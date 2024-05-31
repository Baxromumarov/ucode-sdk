CREATE TYPE "gender" AS ENUM (
  'Male',
  'Female',
  'Other'
);

CREATE TABLE "Организация.Клиент.customer" (
  "guid" uuid,
  "Наименование.name" varchar,
  "Телефон.phone_number" phone,
  "Гендер.gender" gender,
  "car_id" uuid
);

CREATE TABLE "Мошина.car" (
  "guid" uuid,
  "Наименование.name" varchar,
  "Номер.car_number" varchar,
  "Модель.model" varchar
);

ALTER TABLE "Мошина.car" ADD FOREIGN KEY ("guid") REFERENCES "Организация.Клиент.customer" ("car_id");
