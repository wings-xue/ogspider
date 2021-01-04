-- Table: public.job

-- DROP TABLE public.job;

CREATE TABLE public.job
(
    id integer NOT NULL DEFAULT nextval('job_id_seq'::regclass),
    url character varying(1000) COLLATE pg_catalog."default" NOT NULL,
    host character varying(300) COLLATE pg_catalog."default" NOT NULL,
    download character varying(300) COLLATE pg_catalog."default",
    datas json NOT NULL,
    uuid character varying COLLATE pg_catalog."default",
    status character varying COLLATE pg_catalog."default",
    retry integer,
    log character varying COLLATE pg_catalog."default",
    cookie character varying(300) COLLATE pg_catalog."default",
    insert_date timestamp with time zone,
    update_date timestamp with time zone,
    fresh_life bigint,
    alive_num bigint,
    CONSTRAINT job_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.job
    OWNER to postgres;
-- Index: uuid_index

-- DROP INDEX public.uuid_index;

CREATE UNIQUE INDEX uuid_index
    ON public.job USING btree
    (uuid COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;