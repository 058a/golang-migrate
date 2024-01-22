-- Table: public.StockItem

-- DROP TABLE IF EXISTS public."StockItem";

CREATE TABLE IF NOT EXISTS public."stock_item"
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    "created_at" timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp(3) without time zone NOT NULL,
    CONSTRAINT "stock_item_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."stock_item"
    OWNER to "user";