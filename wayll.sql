CREATE TYPE "gender" AS ENUM (
  'male',
  'female'
);

CREATE TYPE "investor_status" AS ENUM (
  'new',
  'accepted',
  'rejected',
  'not_verified'
);

CREATE TYPE "project_status" AS ENUM (
  'investment',
  'collect_money',
  'closed'
);

CREATE TYPE "chart_of_account_category" AS ENUM (
  'income',
  'expense',
  'asset',
  'commitment',
  'capital'
);

CREATE TYPE "account_status" AS ENUM (
  'opened',
  'closed',
  'suspended',
  'awaiting_confirmation'
);

CREATE TYPE "owner_type" AS ENUM (
  'individual',
  'entity',
  'financial_account'
);

CREATE TYPE "account_type" AS ENUM (
  'cash',
  'bank_account',
  'card',
  'transit',
  'project',
  'agreement',
  'financial_account',
  'invoice',
  'equipment',
  'deposit',
  'investment'
);

CREATE TYPE "order_status" AS ENUM (
  'canceled',
  'rejected',
  'pending',
  'filled'
);

CREATE TYPE "order_side" AS ENUM (
  'buy',
  'sell',
  'purchase'
);

CREATE TYPE "transaction_status" AS ENUM (
  'pending',
  'confirmed',
  'finished'
);

CREATE TYPE "payment_type" AS ENUM (
  'bank',
  'card',
  'cash'
);

CREATE TYPE "operation" AS ENUM (
  'topup',
  'withdraw',
  'dividend',
  'buy',
  'sell'
);

CREATE TYPE "dokument_status" AS ENUM (
  'not_signed',
  'signed'
);

CREATE TYPE "multiselect" AS ENUM (
  'select',
  'multiselect',
  'input'
);

CREATE TYPE "project_type" AS ENUM (
  'public',
  'private'
);

CREATE TYPE "dividend_type" AS ENUM (
  'akt'
);

CREATE TABLE "Справочник.Валюта.currency" (
  "guid" uuid,
  "Имя.name" varchar,
  "Короткое имя.short_name" varchar,
  "Код.code" varchar,
  "Фото.photo" photo
);

CREATE TABLE "Chart.chart_of_account" (
  "guid" uuid,
  "Код.code" varchar,
  "Имя.name" varchar,
  "Категория.category" chart_of_account_category
);

CREATE TABLE "Пакет.Пакет.projects" (
  "guid" UUID,
  "Имя.name" VARCHAR,
  "Общий бюджет.total_budget" float,
  "Собранная сумма.raised_amount" float,
  "Примерный доход.approx_income" float,
  "Местоположение.location" text,
  "Описание.description" text,
  "Выплата дивидендов.dividend_payment" float,
  "Дата начала.start_date" date,
  "Дата окончания.end_date" date,
  "Статус.status" project_status,
  "Изображение.image" text,
  "Оплаченная сумма.paid_amount" float,
  "Количество лотов.number_of_lots" int,
  "Сумма лота.lot_amount" float,
  "Минимальная сумма.minimal_sum" float,
  "Тип.type" project_type,
  "Тип дивидендов.dividend_type" dividend_type
);

CREATE TABLE "Лоты.lot" (
  "guid" uuid,
  "ID проекта.projects_id" uuid,
  "Сумма лота.lot_sum" float,
  "Имя.name" varchar,
  "Доля.portion" float
);

CREATE TABLE "Документы.documents" (
  "guid" uuid,
  "Имя.name" varchar,
  "Файл.file" text,
  "ID проекта.projects_id" uuid
);

CREATE TABLE "Инвесторы.Инвесторы.investor" (
  "guid" uuid,
  "Имя.name" varchar,
  "Псевдоним.nickname" varchar,
  "Телефон.phone" varchar,
  "Электронная почта.email" varchar,
  "Фамилия.surname" varchar,
  "Отчество.last_name" varchar,
  "Место рождения.birth_place" varchar,
  "Дата рождения.birth_date" date,
  "Пол.gender" gender,
  "ПИНФЛ.pinfl" varchar,
  "Номер паспорта.passport_number" varchar,
  "Срок действия.passport_expire_date" date,
  "Город.city" varchar,
  "Дом.home" varchar,
  "Квартира.apartment" varchar,
  "Улица.street" varchar,
  "Почтовый индекс.postcode" varchar,
  "ID роли.role_id" uuid,
  "FCM токен.fcm_token" text,
  "Изображение.image" text,
  "Статус.status" investor_status,
  "KYC пройден.kyc_passed" switch,
  "Телефон в Telegram.telegram_phone" varchar,
  "Тип платформы.platform_type" varchar,
  "Тип проекта.project_type" project_type
);

CREATE TABLE "Документы.dokuments" (
  "guid" uuid,
  "ID инвестора.investor_id" uuid,
  "Файл.file" text,
  "Статус.status" dokument_status,
  "ID проекта.projects_id" uuid
);

CREATE TABLE "Инвесторские лоты.investor_lot" (
  "guid" uuid,
  "ID инвестора.investor_id" uuid,
  "ID лота.lot_id" uuid,
  "Сумма.amount" float,
  "Сумма лота.lot_amount" float,
  "Процент доли.share_percent" float
);

CREATE TABLE "Финанс.Счет.account" (
  "guid" uuid,
  "Имя.name" varchar,
  "Номер счета.account_number" varchar,
  "ID плана счетов.chart_of_account_id" uuid,
  "Описание.description" text,
  "ID валюты.currency_id" uuid,
  "Источник данных.data_source" varchar,
  "Статус.status" account_status,
  "Тип владельца.owner_type" owner_type,
  "Тип.type" account_type,
  "ID инвестора.investor_id" uuid,
  "Холдинговая компания.holding_company" varchar,
  "Дата открытия.opened_date" timestamp,
  "Дата закрытия.closed_date" timestamp,
  "Дата обновления.updated_date" timestamp,
  "Баланс.balance" float,
  "Кредитный лимит.credit_limit" float,
  "Основной счет.is_main" switch
);

CREATE TABLE "Ордер.orders" (
  "guid" uuid,
  "ID инвестора.investor_id" uuid,
  "ID проекта.projects_id" uuid,
  "Цена.price" float,
  "Количество.quantity" float,
  "Общая цена.total_price" float,
  "Прибыль и убыток.profit_and_lost" float,
  "Дата и время.date_time" timestamp,
  "Тип.type" order_side,
  "Статус.status" order_status,
  "ID счета.account_id" uuid,
  "Дата создания.created_date" timestamp,
  "Дата обновления.updated_date" timestamp,
  "Дата удаления.deleted_date" timestamp 
);

CREATE TABLE "Транзакция.transaction" (
  "guid" uuid,
  "ID счета.account_id" uuid,
  "Сумма.amount" float,
  "ID валюты.currency_id" uuid,
  "Описание.description" text,
  "Статус.status" transaction_status,
  "Дата начисления.accrual_date" timestamp,
  "Дата создания.created_date" timestamp,
  "Валютный курс.currency_rate" float,
  "ID проекта.projects_id" uuid,
  "ID инвестора.investor_id" uuid,
  "Тип оплаты.payment_type" payment_type,
  "Операция.operation" operation,
  "ID ордера.order_id" uuid,
  "Дата платежа.payment_date" date,
  "Количество.quantity" float
);

CREATE TABLE "Order Match.order_match" (
  "ID.id" uuid,
  "ID основного ордера.primary_order_id" uuid,
  "ID вторичного ордера.secondary_order_id" uuid
);

CREATE TABLE "Организация.Сотрудник.Employee" (
  "guid" uuid,
  "Имя.name" varchar,
  "Логин.login" varchar,
  "Пароль.parol" varchar,
  "ID типа клиента.client_type_id" uuid,
  "ID роли.role_id" uuid
);

CREATE TABLE "Юр лицо.legal_entity" (
  "guid" uuid,
  "Имя.name" varchar,
  "Название банка.bank_name" varchar,
  "ИНН.inn" varchar,
  "ОКЕД.oked" varchar,
  "МФО.mfo" varchar,
  "Корреспондентский счет.correspondent_account" varchar,
  "Номер счета.account_number" varchar,
  "ID валюты.currency_id" uuid,
  "ID проекта.projects_id" uuid
);

CREATE TABLE "Маркетинг.Новости.news" (
  "guid" uuid,
  "ID проекта.projects_id" uuid,
  "Содержание.content" text,
  "Содержание.toc" text,
  "Дата выпуска.date_of_issue" timestamp,
  "Баннер.banner" varchar,
  "Фото.photo" photo,
  "Тег.tag" switch,
  "Уведомление.notification" switch
);

CREATE TABLE "Уведомление.notification" (
  "guid" uuid,
  "Заголовок.title" varchar,
  "Содержание.toc" text,
  "Дата выпуска.date_of_issue" timestamp,
  "Фото.photo" photo,
  "Баннер.banner" varchar,
  "Уведомление.notification" switch
);

CREATE TABLE "Опросник.questionnaire" (
  "guid" uuid,
  "Имя.name" varchar,
  "ID опросника.questionnaire_id" uuid,
  "Мультивыбор.multiselect" multiselect
);

CREATE TABLE "FAQ.faq" (
  "guid" uuid,
  "Имя.name" varchar,
  "ID FAQ.faq_id" uuid
);

ALTER TABLE "Документы.documents" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");

ALTER TABLE "Лоты.lot" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");

ALTER TABLE "Инвесторские лоты.investor_lot" ADD FOREIGN KEY ("investor_id") REFERENCES "Инвесторы.investor" ("guid");

ALTER TABLE "Инвесторские лоты.investor_lot" ADD FOREIGN KEY ("lot_id") REFERENCES "Лоты.lot" ("guid");

ALTER TABLE "Счет.account" ADD FOREIGN KEY ("chart_of_account_id") REFERENCES "Chart.chart_of_account" ("guid");

ALTER TABLE "Счет.account" ADD FOREIGN KEY ("currency_id") REFERENCES "Валюта.currency" ("guid");

ALTER TABLE "Счет.account" ADD FOREIGN KEY ("investor_id") REFERENCES "Инвесторы.investor" ("guid");

ALTER TABLE "Ордер.orders" ADD FOREIGN KEY ("investor_id") REFERENCES "Инвесторы.investor" ("guid");

ALTER TABLE "Ордер.orders" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");

ALTER TABLE "Ордер.orders" ADD FOREIGN KEY ("account_id") REFERENCES "Счет.account" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("account_id") REFERENCES "Счет.account" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("currency_id") REFERENCES "Валюта.currency" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("investor_id") REFERENCES "Инвесторы.investor" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("orders_id") REFERENCES "Ордер.orders" ("guid");

ALTER TABLE "Юр лицо.legal_entity" ADD FOREIGN KEY ("currency_id") REFERENCES "Валюта.currency" ("guid");

ALTER TABLE "Юр лицо.legal_entity" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");

ALTER TABLE "Документы.dokuments" ADD FOREIGN KEY ("investor_id") REFERENCES "Инвесторы.investor" ("guid");

ALTER TABLE "Документы.dokuments" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");

ALTER TABLE "Новости.news" ADD FOREIGN KEY ("projects_id") REFERENCES "Пакет.projects" ("guid");







