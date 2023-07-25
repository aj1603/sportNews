CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    "language" VARCHAR
);

CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    surname VARCHAR,
    "password" VARCHAR,
    email VARCHAR UNIQUE,
    phone_num INTEGER UNIQUE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "password" VARCHAR,
    phone_num INTEGER UNIQUE
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    img_url VARCHAR
);

CREATE TABLE subcategories (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    img_url VARCHAR,
    categories_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id)
);

CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    context VARCHAR,
    "type" INTEGER,
    categories_id INTEGER,
    subcategories_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id),

    CONSTRAINT subcategories_id
        FOREIGN KEY (subcategories_id)
            REFERENCES subcategories (id)
);

CREATE TABLE news_img (
    id SERIAL PRIMARY KEY,
    img_url VARCHAR,
    news_id INTEGER,

    CONSTRAINT news_id
        FOREIGN KEY (news_id)
            REFERENCES news (id)
);

CREATE TABLE athlete (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    surname VARCHAR,
    father_name INTEGER,
    "description" VARCHAR,
    is_turkmen INTEGER,
    categories_id INTEGER,
    subcategories_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id),

    CONSTRAINT subcategories_id
        FOREIGN KEY (subcategories_id)
            REFERENCES subcategories (id)
);

CREATE TABLE athlete_img (
    id SERIAL PRIMARY KEY,
    img_url VARCHAR,
    athlete_id INTEGER,

    CONSTRAINT athlete_id
        FOREIGN KEY (athlete_id)
            REFERENCES athlete (id)
);

CREATE TABLE categories_tran (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    categories_id INTEGER,
    languages_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id),

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id)
);