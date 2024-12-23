-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
SET timezone = 'UTC';

-- create user schema
CREATE SCHEMA IF NOT EXISTS "user" ;

-- create user table in user schema
CREATE TABLE "user"."user"
(
    user_id                 SERIAL PRIMARY KEY,
    uuid               UUID NOT NULL UNIQUE ,
    first_name              VARCHAR(100),
    last_name               VARCHAR(100),
    avatar                  VARCHAR(100),
    created_at              timestamptz NOT NULL DEFAULT now(),
    updated_at              timestamptz NOT NULL DEFAULT now(),
    active                  BOOLEAN DEFAULT TRUE

);

CREATE INDEX idx_uuid ON "user"."user"(uuid);

-- create wallet table in user schema
CREATE TABLE "user".wallet
(
    wallet_id                 SERIAL PRIMARY KEY,
    user_fk                INT REFERENCES "user"."user"(user_id),
    balance                  DECIMAL(10,2) DEFAULT 0,
    created_at              timestamptz NOT NULL DEFAULT now(),
    updated_at              timestamptz NOT NULL DEFAULT now(),
    active                  BOOLEAN DEFAULT TRUE
);


-- create auth schema
CREATE SCHEMA IF NOT EXISTS  "auth";

-- create auth table in auth schema
CREATE TABLE "auth".auth
(
    auth_id                 SERIAL PRIMARY KEY,
    email                   VARCHAR(100) UNIQUE,
    user_fk                INT REFERENCES "user"."user"(user_id),
    password                VARCHAR(60) NOT NULL,
    created_at              timestamptz NOT NULL DEFAULT now(),
    updated_at              timestamptz NOT NULL DEFAULT now(),
    active                  BOOLEAN DEFAULT TRUE
);

CREATE INDEX idx_email ON "auth".auth(email);

-- create slot schema
CREATE SCHEMA IF NOT EXISTS  "slot";

CREATE TABLE "slot".spin_result_type
(
    spin_result_type_id            SERIAL PRIMARY KEY,
    key                  VARCHAR(100) UNIQUE,
    name                   VARCHAR(100) UNIQUE,
    created_at              timestamptz NOT NULL DEFAULT now(),
    updated_at              timestamptz NOT NULL DEFAULT now(),
    active                  BOOLEAN DEFAULT TRUE
);

INSERT INTO "slot".spin_result_type(key, name) VALUES
('x10', 'Profit x10'),
('x2', 'Profit x2'),
('loss', 'Loss');



-- create slot table in slot schema
CREATE TABLE "slot".spin
(
    slot_game_id                 SERIAL PRIMARY KEY,
    user_fk                   INT REFERENCES "user"."user"(user_id),
    bet                DECIMAL(10,2)  NOT NULL ,
    payout               DECIMAL(10,2)  NOT NULL ,
    symbols                     INT[] NOT NULL,
    winning                    BOOLEAN NOT NULL ,
    spin_result_type_fk    INT REFERENCES "slot".spin_result_type(spin_result_type_id),
    user_balance               DECIMAL(10,2)  NOT NULL ,
    created_at              timestamptz NOT NULL DEFAULT now()
);


-- +migrate Down
-- SQL in section 'Down' is executed when this migration is applied

DROP TABLE "auth"."auth";
DROP SCHEMA "auth";
DROP TABLE "user".wallet;
DROP TABLE "user"."user";
DROP SCHEMA "user";
DROP SCHEMA "slot".spin;
DROP SCHEMA "slot".spin_result_type;
DROP SCHEMA "slot";


