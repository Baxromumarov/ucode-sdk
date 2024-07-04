CREATE TYPE "region_type" AS ENUM (
  'city',
  'district'
);

CREATE TYPE "order_type" AS ENUM (
  'catalog',
  'master'
);

CREATE TYPE "payment_type" AS ENUM (
  'lick',
  'payme'
);

CREATE TYPE "client_status" AS ENUM (
  'order_accepted',
  'workshop_accepted',
  'ready_for_delivery',
  'courier_accepted',
  'delivered',
  'finished'
);

CREATE TYPE "master_status" AS ENUM (
  'in_process',
  'created',
  'approved'
);

CREATE TYPE "manufactory_status" AS ENUM (
  'in_process',
  'created',
  'approved'
);

CREATE TYPE "courier_status" AS ENUM (
  'in_process',
  'created',
  'approved'
);


CREATE TYPE "door_model_type" AS ENUM (
  'dila1',
  'dila2'
);

CREATE TYPE "opening_type" AS ENUM (
  'outward_right',
  'outward_left',
  'inside_left',
  'inside_right'
);

CREATE TYPE "door_step_type" AS ENUM (
  'autodoorstep',
  'with_doorstep',
  'without_doorstep'
);

CREATE TYPE "lock_type" AS ENUM (
  'magnetic',
  'standart'
);

CREATE TYPE "hinge_type" AS ENUM (
  'secret',
  'standart',
  'butterfly'
);

CREATE TYPE "texture_type" AS ENUM (
  'vertical',
  'horizontal'
);

CREATE TYPE "wood_type" AS ENUM (
  'veneered',
  'painted'
);

CREATE TYPE "color_type" AS ENUM (
  'ncs',
  'ral'
);

CREATE TYPE "sales_unit" AS ENUM (
  'dona',
  'metr',
  'litr',
  'kg',
  'kv'
);

CREATE TYPE "type" AS ENUM (
  'door',
  'cover',
  'bottom',
  'crown',
  'platband',
  'wood',
  'color',
  'door_handle',
  'lock',
  'hinge'
);

CREATE TYPE "furniture_position" AS ENUM (
  'bottom',
  'up'
);

CREATE TYPE "furniture_type" AS ENUM (
  'kitchen',
  'dressing_room'
);

CREATE TYPE "employee_type" AS ENUM (
  'master',
  'driver',
  'installer'
);

CREATE TYPE "object_type" AS ENUM (
  'door',
  'panel',
  'furniture'
);

CREATE TYPE "receiver" AS ENUM (
  'client',
  'employee'
);

CREATE TYPE "request_status" AS ENUM (
  'new',
  'in_process',
  'finished'
);

CREATE TYPE "payment_status" AS ENUM (
  'pending',
  'canceled',
  'done'
);


CREATE TYPE "door_items_type" AS ENUM{
  'on_the_one_side',
  'on_both_sides',
  'without'
}

CREATE TABLE "Справочник.Регион.region" (
  "guid" uuid PRIMARY KEY,
  "Имя.name" varchar,
  "Тип.type" region_type,
  "Регион.region_id" uuid
);

CREATE TABLE "Фабрика.manufactory" (
  "guid" uuid PRIMARY KEY,
  "Названия.name" varchar,
  "Номер Телефона.phone" varchar,
  "Регион.region_id" varchar,
  "Почта.email" varchar,
  "Адрес.address" varchar,
  "Локация.location" varchar,
  "Названия Компании.company_name" varchar
);

CREATE TABLE "Цвет.color" (
  "guid" uuid,
  "Названия.name" varchar,
  "Фотография.photo" varchar
);

CREATE TABLE "Параметр.parameter" (
  "guid" uuid,
  "Промежуток Для Пены.foam_gap" float,
  "Ширина Коробки.box_width" float,
  "Обклад Высота.cover_height" float,
  "Обклад Ширина.cover_width" float,
  "Обклад Обклад.cover_cover" float
);

CREATE TABLE "Основная Цена.master_price" (
  "guid" uuid,
  "Тип Сотрудника.employee_type" employee_type,
  "Тип Обекта.object_type" object_type,
  "Сумма От.price_from" float,
  "Сумма До.price_to" float,
  "Количества Продукта.product_count" int,
  "Сумма.price" float
);


CREATE TABLE "Нижняя Панель.bottom_bar" (
  "guid" uuid,
  "Названия.name" varchar,
  "Слуг.slug" varchar,
  "Иконка.icon" varchar,
  "Активен.is_active" bool,
  "Тип.type" object_type
);

CREATE TABLE "Пользователи.Сотрудники.employee" (
  "guid" uuid,
  "Имя.name" varchar,
  "Логин.login" varchar,
  "Пароль.password" varchar,
  "Номер Телефона.phone" varchar,
  "Роль.role_id" uuid,
  "Фабрика.manufactory_id" uuid,
  "Регион.region_id" uuid,
  "Почта.email" varchar,
  "Адрес.address" varchar
);

CREATE TABLE "Клиент.client" (
  "guid" uuid,
  "Имя.name" varchar,
  "Номер Телефона.phone" varchar,
  "Регион.region_id" varchar,
  "Адрес.address" varchar
);

CREATE TABLE "Заказы.Заказ.orders" (
  "guid" uuid,
  "Ид.id" varchar,
  "Тип Заказа.order_type" order_type,
  "Клиент.client_id" uuid,
  "Фабрика.manufactory_id" uuid,
  "Дата Создние.created_time" datetime,
  "Статус Клиента.client_status" client_status,
  "Статус Установшика.master_status" master_status,
  "Статус Фабрики.manufactory_status" manufactory_status,
  "Статус Курера.courier_status" courier_status,
  "Имя Клиента.client_name" varchar,
  "Номер Клиента.client_phone" varchar,
  "Адрес Клиента.client_address" varchar,
  "Коментария.comment" text,
  "Дата Установки.installation_date" datetime,
  "Дата Доставки.delivery_date" datetime,
  "Обшая Сумма.total_sum" float,
  "Оплаченая Сумма.paid_sum" float,
  "Статус Оплаты.paid_status" bool,
  "Тип Оплаты.payment_type" payment_type,
  "Количества Продуктов.product_count" int
);

CREATE TABLE "Заказ Продукт.order_product" (
  "guid" uuid,
  "Элемент.item_id" uuid,
  "Заказ.orders_id" uuid,
  "Количества.count" int,
  "Тип Измерения.sales_unit" sales_unit,
  "Сумма.price" float,
  "Общая Сумма.total_price" float
);

CREATE TABLE "Каталог.Категории.category" (
  "guid" uuid,
  "Названия.name" varchar,
  "Активен.is_active" bool,
  "Основной.is_main" bool,
  "Фотография.photo" varchar
);

CREATE TABLE "Готовые Продукции.product" (
  "guid" uuid,
  "Названия.name" varchar,
  "Категория.category_id" uuid,
  "Сумма.price" float,
  "Фотография.photo" uuid[],
  "Цвета.color_ids" uuid[],
  "Фабрика.manufactory_id" uuid,
  "Производитель.manufacturer" varchar,
  "Описания.description" text
);

CREATE TABLE "Корзинка.basket" (
  "guid" uuid,
  "Клиент.client_id" uuid,
  "Продукт.product_id" uuid,
  "Заказ.orders_id" uuid,
  "Цветь.color_id" uuid,
  "Количества.count" int
);

CREATE TABLE "Предметы.Модель Двери.door_model" (
  "guid" uuid,
  "Названия.name" uuid,
  "Фотография.photo" varchar,
  "Тип.type" door_model_type
);

CREATE TABLE "Панель.panel" (
  "guid" uuid,
  "Ширина.width" float,
  "Высота.height" float,
);

CREATE TABLE "Двери.door" (
  "guid" uuid,
  "Модель Двери.door_model_id" uuid,
  "Названия.name" varchar,
  "Название Комнаты.room_name" varchar,
  "Заказ.orders_id" uuid,
  "Категория.category_id" uuid,
  "Фотография.photo" varchar,
  "Тип Открытия.opening_type" opening_type,
  "Тип Порога.doorstep_type" door_step_type,
  "Тип Короны.door_crown_type" door_items_type,
  "Тип Сопожка.door_boot_type" door_items_type,
  "Тип Наличника.door_platband_type" door_items_type,
  "Элементы.item_ids" []uuid,
  "Общая Сумма.total_price" float,
  "Описания.description" text,
  "Количества.count" int,
  "Ширина.width" float,
  "Высота.height" float,
  "Толщина.thickness" float,
  "Ширина Коробки.box_width" float,
  "Высота Каробки.box_height" float,
  "Общая Ширина.clear_width" float,
  "Общая Высота.clear_height" float
);


CREATE TABLE "Предметь.item" (
  "guid" uuid,
  "Названия.name" varchar,
  "Модель.model" varchar,
  "Код Предмета.item_code" varchar,
  "Тип Измерения.sales_unit" sales_unit,
  "Фотография.photo" varchar,
  "Высота.height" float,
  "Толщина.thickness" float,
  "Ширина.width" float,
  "Тип.type" Type,
  "Тип Текстуры.texture_type" texture_type,
  "Тип Дерево.wood_type" wood_type,
  "Тип Цвета.color_type" color_type,
  "Тип Замка.lock_type" lock_type,
  "Тип Шарнира.hinge_type" hinge_type,
  "Сумма.price" float,
  "Фабрика.manufactory_id" uuid
);

CREATE TABLE "Мебель.furniture" (
  "guid" uuid,
  "Позиция.position" furniture_position,
  "Тип.type" furniture_type,
  "Ширина.width" float,
  "Высота.height" float,
  "Толщина.thickness" float,
  "Названия.name" varchar,
  "Фотография.photo" varchar,
  "Сумма.price" float
);

CREATE TABLE "Мебель Предметы.furniture_items" (
  "guid" uuid,
  "Мебель.furniture_id" uuid,
  "Предметь.item_id" uuid,
  "Количества.quantity" float,
  "Сумма.price" float,
  "Тип Измерения.sales_unit" sales_unit,
  "Общая Сумма.total_price" float
);


CREATE TABLE "Уведомления.Уведомление.notification" (
  "guid" uuid,
  "Титуль.title" varchar,
  "Текст.text" text,
  "Дата Создания.created_time" datetime,
  "Фотография.photo" varchar,
  "Получатель.receiver" receiver
);

CREATE TABLE "Отзывы.Отзыв.review" (
  "guid" uuid,
  "Рейтинг.rating" float,
  "Текст.text" text,
  "Сотрудник.employee_id" uuid
);

CREATE TABLE "Финансы.Транзакция.transaction" (
  "guid" uuid PRIMARY KEY,
  "Транзакция.transaction_id" varchar,
  "Клиент.client_id" uuid,
  "Заказ.orders_id" uuid,
  "Статус.status" payment_status,
  "Сумма.amount" float,
  "Примечание.note" varchar,
  "Дата Создания.created_time" datetime,
  "Время Выполнения.perform_time" datetime,
  "Время Отмены.cancel_time" datetime,
  "Состояние.state" int,
  "Причина.reason" int,
  "Врем Оплаты.time_payme" int,
  "Причина Ошибки.error_reason" text,
  "Тип Оплаты.payment_type" payment_type
);




ALTER TABLE "Регион.region" ADD FOREIGN KEY ("guid") REFERENCES "Сотрудник.employee" ("region_id");

ALTER TABLE "Цех.manufactory" ADD FOREIGN KEY ("guid") REFERENCES "Сотрудник.employee" ("manufactory_id");

ALTER TABLE "Клиент.client" ADD FOREIGN KEY ("guid") REFERENCES "Заказ.orders" ("client_id");

ALTER TABLE "Цех.manufactory" ADD FOREIGN KEY ("guid") REFERENCES "Заказ.orders" ("manufactory_id");

ALTER TABLE "Продукть.product" ADD FOREIGN KEY ("category_id") REFERENCES "Категория.category" ("guid");

ALTER TABLE "Продукть.product" ADD FOREIGN KEY ("manufactory_id") REFERENCES "Цех.manufactory" ("guid");

ALTER TABLE "Корзинка.basket" ADD FOREIGN KEY ("client_id") REFERENCES "Клиент.client" ("guid");

ALTER TABLE "Корзинка.basket" ADD FOREIGN KEY ("product_id") REFERENCES "Продукть.product" ("guid");

ALTER TABLE "Корзинка.basket" ADD FOREIGN KEY ("orders_id") REFERENCES "Заказ.orders" ("guid");

ALTER TABLE "Корзинка.basket" ADD FOREIGN KEY ("color_id") REFERENCES "Цветь.color" ("guid");

ALTER TABLE "Дверь.door" ADD FOREIGN KEY ("door_model_id") REFERENCES "Модель Дверя.door_model" ("guid");

ALTER TABLE "Заказ.orders" ADD FOREIGN KEY ("guid") REFERENCES "Дверь.door" ("orders_id");

ALTER TABLE "Категория.category" ADD FOREIGN KEY ("guid") REFERENCES "Двер.door" ("category_id");

ALTER TABLE "Предметь.item" ADD FOREIGN KEY ("manufactory_id") REFERENCES "Цех.manufactory" ("guid");

ALTER TABLE "Предметы Мебелья.furniture_items" ADD FOREIGN KEY ("furniture_id") REFERENCES "Мебель.furniture" ("guid");

ALTER TABLE "Предметы Мебелья.furniture_items" ADD FOREIGN KEY ("item_id") REFERENCES "Предметь.item" ("guid");

ALTER TABLE "Заказ Продукть.order_product" ADD FOREIGN KEY ("item_id") REFERENCES "Предметь.item" ("guid");

ALTER TABLE "Заказ Продукть.order_product" ADD FOREIGN KEY ("orders_id") REFERENCES "Заказ.orders" ("guid");

ALTER TABLE "Сотрудник.employee" ADD FOREIGN KEY ("guid") REFERENCES "Отзыв.review" ("employee_id");

ALTER TABLE "Запрос.request" ADD FOREIGN KEY ("client_id") REFERENCES "Клиент.client" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("client_id") REFERENCES "Клиент.client" ("guid");

ALTER TABLE "Транзакция.transaction" ADD FOREIGN KEY ("orders_id") REFERENCES "Заказ.orders" ("guid");

ALTER TABLE "Сотрудник.employee" ADD FOREIGN KEY ("guid") REFERENCES "ОТП.otp" ("employee_id");
