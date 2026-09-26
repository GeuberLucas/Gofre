-- ==========================================
-- 1. AUTH MICROSERVICE
-- ==========================================
BEGIN;

CREATE SCHEMA IF NOT EXISTS auth AUTHORIZATION postgres;

-- Tabela de utilizadores enxuta (apenas dados de credenciais/autenticação)
CREATE TABLE IF NOT EXISTS auth.users (
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY (
        INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1
    ),
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL UNIQUE,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS auth.reset_tokens (
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY (
        INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1
    ),
    user_id integer NOT NULL,
    hash_token character varying(255) COLLATE pg_catalog."default" NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    CONSTRAINT reset_tokens_pkey PRIMARY KEY (id),
    CONSTRAINT reset_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES auth.users (id) ON DELETE CASCADE
);

-- Índice na Foreign Key
CREATE INDEX idx_reset_tokens_user_id ON auth.reset_tokens(user_id);

COMMIT;

-- ==========================================
-- 2. PROFILES MICROSERVICE (Novo)
-- ==========================================
BEGIN;

CREATE SCHEMA IF NOT EXISTS profiles AUTHORIZATION postgres;

CREATE TABLE IF NOT EXISTS profiles.user_profiles (
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL UNIQUE,
    full_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    mobile_phone character varying(20) COLLATE pg_catalog."default",
    initial_balance bigint DEFAULT 0,
    CONSTRAINT user_profiles_user_id_fkey FOREIGN KEY (user_id) REFERENCES auth.users (id) ON DELETE CASCADE
);

COMMIT;

-- ==========================================
-- 3. TRANSACTIONS MICROSERVICE
-- ==========================================
BEGIN;

CREATE SCHEMA IF NOT EXISTS transactions AUTHORIZATION postgres;

-- Os ENUMs agora são criados explicitamente dentro do schema transactions
CREATE TYPE transactions.expense_category AS ENUM(
    'Mercado geral', 'Delivery', 'Restaurante e bares', 'Vestuário', 'Moradia', 'Utilidades', 'Decoração', 'Educação', 'Dependentes', 'Saúde', 'Entretenimento', 'Serviços', 'Impostos', 'Transporte', 'Presentes', 'Pets', 'Viagens', 'Doações', 'Apostas', 'Livre', 'Outros'
);
CREATE TYPE transactions.expense_type AS ENUM('Mensal', 'Variável', 'Fatura');
CREATE TYPE transactions.payment_method AS ENUM('pix', 'debito', 'credito', 'boleto', 'dinheiro', 'ted', 'cheque');
CREATE TYPE transactions.income_type AS ENUM('Trabalho', 'Extra', 'Investimento', 'Aposentadoria', 'Resgate', 'Outros');

CREATE TABLE IF NOT EXISTS transactions.expenses(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    description character varying(255) NOT NULL,
    target character varying(255),
    category transactions.expense_category NOT NULL,
    type transactions.expense_type NOT NULL,
    payment_method transactions.payment_method,
    payment_date timestamp with time zone NOT NULL,
    amount bigint NOT NULL,
    is_paid boolean NOT NULL DEFAULT False
);

-- Índice na Foreign Key
CREATE INDEX idx_expenses_user_id ON transactions.expenses(user_id);

CREATE TABLE IF NOT EXISTS transactions.revenue(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    description character varying(255) NOT NULL,
    origin character varying(255),
    type transactions.income_type NOT NULL,
    amount bigint NOT NULL,
    received_date timestamp with time zone NOT NULL,
    is_received boolean NOT NULL DEFAULT False
);

-- Índice na Foreign Key
CREATE INDEX idx_revenue_user_id ON transactions.revenue(user_id);

COMMIT;

-- ==========================================
-- 4. INVESTMENTS MICROSERVICE
-- ==========================================
BEGIN;

CREATE SCHEMA IF NOT EXISTS investments AUTHORIZATION postgres;

CREATE TABLE IF NOT EXISTS investments.asset(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name character varying(255) NOT NULL
);

INSERT INTO investments.asset(name)
VALUES
('Títulos privados'), ('Títulos públicos'), ('Ações'), ('ETFs'), ('FIIs'), ('Fundos'), ('Commodities'), ('Derivativos'), ('Criptomoeda'), ('Exterior'), ('Poupança'), ('Outros');

CREATE TABLE IF NOT EXISTS investments.portfolio(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    asset_id integer NOT NULL REFERENCES investments.asset(id) ON DELETE RESTRICT,
    deposit_date timestamp with time zone NOT NULL,
    broker character varying(255) NOT NULL,
    amount bigint NOT NULL,
    description character varying(255) NOT NULL,
    is_done boolean NOT NULL DEFAULT False
);

-- Índices nas Foreign Keys
CREATE INDEX idx_portfolio_user_id ON investments.portfolio(user_id);
CREATE INDEX idx_portfolio_asset_id ON investments.portfolio(asset_id);

COMMIT;