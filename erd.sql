CREATE TYPE "location_type" AS ENUM (
  'country',
  'city',
  'region',
  'district'
);

CREATE TYPE "gender" AS ENUM (
  'male',
  'female',
  'other'
);

CREATE TYPE "document_type" AS ENUM (
  'ticket'
);

CREATE TYPE "flight_type" AS ENUM (
  'direct',
  'transfer'
);

CREATE TYPE "room_type" AS ENUM (
  'president',
  'luxe',
  'standard'
);

CREATE TABLE "Справочник.Расположение.location" (
  "guid" uuid,
  "Имя.name" varchar,
  "Тип.type" location_type
);

CREATE TABLE "Авиакомпания.airline" (
  "guid" uuid,
  "Имя.name" varchar,
  "Icon.icon" varchar
);

CREATE TABLE "Дополнительная услуга.extra_service" (
  "guid" uuid,
  "Имя.name" varchar,
  "Icon.icon" varchar
);

CREATE TABLE "Билет.ticket" (
  "guid" uuid,
  "От.from_location" uuid,
  "До.to_location" uuid,
  "Цена.price" float,
  "Продолжительность.flight_duration" varchar,
  "from_airline" uuid,
  "to_airline" uuid
);

CREATE TABLE "Способ оплаты.payment_method" (
  "guid" uuid,
  "Имя.name" varchar
);

CREATE TABLE "Заказ.Заказ.order" (
  "guid" uuid,
  "Полное имя.buyer_full_name" varchar,
  "Электронная почта.email" varchar,
  "Номер телефона.phone_number" varchar,
  "Фамилияsurname" varchar,
  "Имя.name" varchar,
  "Дата рождения.birth_date" datetime,
  "Пол.gender" gender,
  "Страна выдачи.country_of_issue" varchar,
  "Тип документа.document_type" document_type,
  "Номер документа.document_number" varchar,
  "Дата выдачи документа.document_issued_date" datetime,
  "Срок действия документа.document_vaidity_period" varchar,
  "Почта.email_passenger" varchar,
  "Номер телефона.phone_number_passenger" varchar,
  "Цена.price" float,
  "Количество пассажиров.passenger_count" float,
  "ticket_id" uuid,
  "hotel_id" uuid,
  "hotel_room_id" uuid,
  "Общая стоимость.total_price" float,
  "payment_method_id" uuid,
  "Тип рейса.flight_type" flight_type
);

CREATE TABLE "Организация.Отель.hotel" (
  "guid" uuid,
  "Расположение.location" varchar,
  "Рейтинг.rating" float,
  "Review.review" float,
  "Имя.name" varchar,
  "Адрес.address" varchar,
  "Описание.description" text
);

CREATE TABLE "Гостиничный номер.hotel_room" (
  "guid" uuid,
  "Тип номера.room_type" room_type,
  "Цена.price" float,
  "hotel_id" uuid,
  "Описание.description" text
);

CREATE TABLE "Категория обзора.review_category" (
  "guid" uuid,
  "Имя.name" varchar
);

CREATE TABLE "Обзор.review" (
  "guid" uuid,
  "hotel_id" uuid,
  "review_category_id" uuid,
  "Рейтинг.rating" float
);

CREATE TABLE "Фотографии.photos" (
  "guid" uuid,
  "hotel_id" uuid,
  "hotel_room_id" uuid
);

CREATE TABLE "Клиент.Клиент.client" (
  "guid" uuid,
  "Полное имя.full_name" varchar,
  "Электронная почта.email" varchar,
  "Номер телефона.phone_number" varchar
);

ALTER TABLE "Билет.ticket" ADD FOREIGN KEY ("from_location") REFERENCES "Расположение.location" ("guid");

ALTER TABLE "Билет.ticket" ADD FOREIGN KEY ("to_location") REFERENCES "Расположение.location" ("guid");

ALTER TABLE "Билет.ticket" ADD FOREIGN KEY ("from_airline") REFERENCES "Авиакомпания.airline" ("guid");

ALTER TABLE "Билет.ticket" ADD FOREIGN KEY ("to_airline") REFERENCES "Авиакомпания.airline" ("guid");

ALTER TABLE "Заказ.order" ADD FOREIGN KEY ("ticket_id") REFERENCES "Билет.ticket" ("guid");

ALTER TABLE "Заказ.order" ADD FOREIGN KEY ("hotel_id") REFERENCES "Отель.hotel" ("guid");

ALTER TABLE "Заказ.order" ADD FOREIGN KEY ("hotel_room_id") REFERENCES "Гостиничный номер.hotel_room" ("guid");

ALTER TABLE "Заказ.order" ADD FOREIGN KEY ("payment_method_id") REFERENCES "payment_method" ("guid");

ALTER TABLE "Гостиничный номер.hotel_room" ADD FOREIGN KEY ("hotel_id") REFERENCES "Отель.hotel" ("guid");

ALTER TABLE "Обзор.review" ADD FOREIGN KEY ("hotel_id") REFERENCES "Отель.hotel" ("guid");

ALTER TABLE "Обзор.review" ADD FOREIGN KEY ("review_category_id") REFERENCES "Категория обзора.review_category" ("guid");

ALTER TABLE "Фотографии.photos" ADD FOREIGN KEY ("hotel_id") REFERENCES "Отель.hotel" ("guid");

ALTER TABLE "Фотографии.photos" ADD FOREIGN KEY ("hotel_room_id") REFERENCES "Гостиничный номер.hotel_room" ("guid");
