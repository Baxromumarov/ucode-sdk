CREATE TABLE "sdk-module.sdk_table" (
  "guid" uuid PRIMARY KEY,
  "balance" float
);

CREATE TABLE "sdk_product" (
  "guid" uuid PRIMARY KEY,
  "product_name" varchar,
  "price" integer,
  "sdk_table_id.balance" uuid
);

ALTER TABLE "sdk_product(to_table)" ADD FOREIGN KEY ("sdk_table_id") REFERENCES "sdk_table(from_table)" ("guid");
-- sdk_table one sdk_product many
ALTER TABLE "sdk_product" ADD FOREIGN KEY ("sdk_table_id") REFERENCES "sdk_table(from_table)" ("guid");
