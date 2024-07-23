CREATE TYPE "invoice_status" AS ENUM (
  'In_process',
  'Signed',
  'Refused',
  'Cancelled'
);

CREATE TYPE "counteragent_type" AS ENUM (
  'Buyer',
  'Seller'
);

CREATE TYPE "invoice_type" AS ENUM (
  'Standart',
  'Additional',
  'Expence',
  'Free',
  'Edited'
);

CREATE TYPE "contract_type" AS ENUM (
  'Act',
  'Contract',
  'Proxy',
  'Document'
);

CREATE TABLE "Организация.Контрагент.counteragent" (
  "guid" uuid,
  "Тип.type" counteragent_type,
  "Имя.name" varchar,
  "ИНН.inn" varchar,
  "Адресс.address" varchar,
  "Рег.Код плательщика НДС.vat_reg_code" varchar,
  "Расчётный счёт.payment_account" varchar,
  "ОКЭД.oked" varchar,
  "МФО.mfo" varchar,
  "Банк.bank" varchar,
  "Директор.director" varchar,
  "Бухгалтер.accountant" varchar
);

CREATE TABLE "Финанс.Договор.contract" (
  "guid" uuid,
  "Номер.number" varchar,
  "Дата.date" date,
  "contract_type" contract_type,
  "Статус.status" invoice_status,
  "counteragent_id" uuid
);

CREATE TABLE "ТТН.ttn" (
  "guid" uuid,
  "Номер ТТН.number" varchar,
  "Дата.date" date,
  "contract_id" uuid,
  "counteragent_id" uuid
);

CREATE TABLE "Счет Фактура.invoice" (
  "guid" uuid,
  "Статус.status" invoice_status,
  "Тип.type" invoice_type,
  "Номер счет фактуры.number" varchar,
  "Дата.date" date,
  "provider_id" uuid,
  "buyer_id" uuid,
  "contract_id" uuid
);

CREATE TABLE "Справочник.Единицы измерения.unit" (
  "guid" uuid,
  "Наименование.name" varchar,
  "Код.code" varchar
);

CREATE TABLE "Каталог.Товары.items" (
  "guid" uuid,
  "Наименоание.name" varchar,
  "Идентификатор и наименование ИКПУ.external_id" varchar,
  "unit_id" uuid,
  "invoice_id" uuid
);

ALTER TABLE "Договор.contract" ADD FOREIGN KEY ("counteragent_id") REFERENCES "Организация.Контрагент.counteragent" ("guid");

ALTER TABLE "ТТН.ttn" ADD FOREIGN KEY ("contract_id") REFERENCES "Договор.contract" ("guid");

ALTER TABLE "ТТН.ttn" ADD FOREIGN KEY ("counteragent_id") REFERENCES "Организация.Контрагент.counteragent" ("guid");

ALTER TABLE "Счет Фактура.invoice" ADD FOREIGN KEY ("provider_id") REFERENCES "Организация.Контрагент.counteragent" ("guid");

ALTER TABLE "Счет Фактура.invoice" ADD FOREIGN KEY ("buyer_id") REFERENCES "Организация.Контрагент.counteragent" ("guid");

ALTER TABLE "Счет Фактура.invoice" ADD FOREIGN KEY ("contract_id") REFERENCES "Договор.contract" ("guid");

ALTER TABLE "Товары.items" ADD FOREIGN KEY ("unit_id") REFERENCES "Единицы измерения.unit" ("guid");

ALTER TABLE "Товары.items" ADD FOREIGN KEY ("invoice_id") REFERENCES "Счет Фактура.invoice" ("guid");
