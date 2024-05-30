CREATE TABLE "Справочник.Филиал.order_item" (
    "guid" uuid PRIMARY KEY,
    "Наименование.name" varchar,
    "расположение.location" varchar,
    "адрес.address" varchar,
    "созданное время.created_time" timestamp
);

CREATE TABLE "Адрес.addresses" (
    "guid" uuid PRIMARY KEY,
    "имя.name" varchar,
    "дом.home" varchar,

);


ALTER TABLE "Филиал.order_item" ADD FOREIGN KEY ("guid") REFERENCES "Адрес.addresses" ("oorder_item_id");
