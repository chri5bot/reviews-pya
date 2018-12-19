-- Table: public.reviews
-- DROP TABLE public.reviews;
CREATE TABLE public.reviews(
    "id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "store_id" uuid NOT NULL,
    "purchase_id" integer NOT NUll,
    "score" integer NOT NUll,
    "opinion" character varying(300) NOT NULL,
    "created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone NOT NULL,
	"deleted_at" timestamp with time zone,
    CONSTRAINT "reviews_pkey" PRIMARY KEY ("id")
)
WITH (OIDS = FALSE)
TABLESPACE pg_default;