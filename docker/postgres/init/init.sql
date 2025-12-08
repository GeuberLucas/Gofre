CREATE DATABASE IF NOT EXISTS "Gofre"
WITH
    OWNER = postgres ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8' LOCALE_PROVIDER = 'libc' TABLESPACE = pg_default CONNECTION
LIMIT
    = -1 IS_TEMPLATE = False;

Begin;
-- auth micro service
CREATE SCHEMA IF NOT EXISTS auth AUTHORIZATION postgres;

CREATE TABLE
    IF NOT EXISTS auth.reset_tokens (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY (
            INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1
        ),
        user_id integer NOT NULL,
        hash_token character varying(255) COLLATE pg_catalog."default" NOT NULL,
        expires_at timestamp
        with
            time zone NOT NULL,
            CONSTRAINT reset_tokens_pkey PRIMARY KEY (id),
            CONSTRAINT reset_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES auth.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
    );
CREATE TABLE
    IF NOT EXISTS auth.users (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY (
            INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1
        ),
        name character varying(255) COLLATE pg_catalog."default" NOT NULL,
        username character varying(255) COLLATE pg_catalog."default" NOT NULL,
        last_name character varying(255) COLLATE pg_catalog."default",
        cell_phone character varying(255) COLLATE pg_catalog."default",
        email character varying(255) COLLATE pg_catalog."default" NOT NULL,
        password character varying(255) COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp
        with
            time zone NOT NULL,
            updated_at timestamp
        with
            time zone NOT NULL,
            CONSTRAINT users_pkey PRIMARY KEY (id)
    );
commit;
Begin;
-- transaction micro service
create schema IF NOT EXISTS transactions;
create type expense_category as ENUM(
'Mercado geral',
'Delivery',
'Restaurante e bares',
'Vestuário',
'Moradia',
'Utilidades',
'Decoração',
'Educação',
'Dependentes',
'Saúde',
'Entretenimento',
'Serviços',
'Impostos',
'Transporte',
'Presentes',
'Pets',
'Viagens',
'Doações',
'Apostas',
'Livre',
'Outros'
);
create type expense_type as enum(
'Mensal',
'Variável',
'Fatura'
);
create type payment_method as enum(
    'pix',
    'debito',
    'credito',
    'boleto',
    'dinheiro',
    'ted',
    'cheque'
);
create type income_type as enum(
    'Trabalho',
    'Extra',
    'Investimento',
    'Aposentadoria',
    'Resgate',
    'Outros'
);

create table IF NOT EXISTS transactions.expenses(
    id serial PRIMARY KEY,
    user_id integer not null,
    description varchar(255) not null,
    target varchar(255),
    category expense_category not null,
    type expense_type not null,
    payment_method payment_method,
    payment_date timestamp with time zone not null,
    amount integer not null,
    is_paid boolean not null default False

);
create table IF NOT EXISTS transactions.revenue(
     id serial PRIMARY KEY,
    user_id integer not null,
    description varchar(255) not null,
    origin varchar(255),
    type income_type not null,
    amount integer not null,
    received_date timestamp with time zone not null,
    is_recieved boolean not null default False
);
commit;
Begin;
-- Investments micro service
create schema IF NOT EXISTS investments;
create table if not exists investments.asset(
    id serial PRIMARY KEY,
    name varchar(255)
);

insert into investments.asset(
    name
)
values
('Títulos privados'),
('Títulos públicos'),
('Ações'),
('ETFs'),
('FIIs'),
('Fundos'),
('Commodities'),
('Derivativos'),
('Criptomoeda'),
('Exterior'),
('Poupança'),
('Outros');



create table if not EXISTS investments.portfolio(
    id serial PRIMARY KEY,
    user_id integer not null,
    asset_id interger not null,
    deposit_date timestamp with time zone not null,
    broker varchar(255) not null,
    FOREIGN KEY (asset_id) REFERENCES investments.asset (id)
);
commit;
BEGIN;
alter table investments.portfolio add column if not exists  amount integer not null;
alter table investments.portfolio add column if not exists  description varchar(255) not null;
alter table investments.portfolio add column if not exists  is_done boolean not null default False;
commit;


begin;
create schema IF NOT EXISTS reports;
CREATE TABLE reports.revenue (
  Month int,
  Year int,
  Planned decimal(10,2),
  Actual decimal(10,2),
  Pending decimal(10,2),
  User_Id int
  UNIQUE ("Month", "Year", "User_Id")
);


CREATE TABLE reports."Aggregated" (
  "Month" int,
  "Year" int,
  "Revenue" decimal(10,2),
  "Expense" decimal(10,2),
  "Investments" decimal(10,2),
  "Montly_without_credit" decimal(10,2),
  "Montly_with_credit" decimal(10,2),
  "Variable_without_credit" decimal(10,2),
  "Variable_with_credit" decimal(10,2),
  "Invoice" decimal(10,2),
  "Result" decimal(10,2),
  "User_Id" int

  UNIQUE ("Month", "Year", "User_Id")
);


CREATE TABLE reports."Investments" (
  "Month" int,
  "Year" int,
  "Planned" decimal(10,2),
  "Actual" decimal(10,2),
  "Pending" decimal(10,2),
  "User_Id" int
  UNIQUE ("Month", "Year", "User_Id")
);



CREATE TABLE reports."Expense" (
  "Month" int,
  "Year" int,
  "Planned" decimal(10,2),
  "Actual" decimal(10,2),
  "Pending" decimal(10,2),
  "Invoice" decimal(10,2),
  "Variable" decimal(10,2),
  "Monthly" decimal(10,2),
  "User_Id" int
  UNIQUE ("Month", "Year", "User_Id")
);

commit;

BEGIN;
CREATE TABLE public."event_track" (
  "event_ID" uuid,
  "aggregate_id" int,
  "consumer_group" varchar(50),
  "processed_at" timestamp
);

CREATE INDEX "pk" ON  "event_track" ("event_ID");
COMMIT;