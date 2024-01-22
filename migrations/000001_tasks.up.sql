-- Table: public.StockItem

-- DROP TABLE IF EXISTS public."StockItem";

CREATE TABLE IF NOT EXISTS public."StockItem"
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    "createdAt" timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" timestamp(3) without time zone NOT NULL,
    CONSTRAINT "StockItem_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."StockItem"
    OWNER to "user";