CREATE TYPE "card_types" AS ENUM (
  'Visa',
  'Click',
  'Mastercard',
  'Humo'
);


CREATE TABLE "Organization.Тип оплаты.payment_method" (
  "guid" uuid PRIMARY KEY,
  "Наименование.name" varchar,
  "созданное время.created_time" date_time,
  "Тип.type" card_types NOT NULL 
);

