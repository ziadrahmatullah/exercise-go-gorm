CREATE DATABASE pet_database;

CREATE TABLE owners(
    owner_id bigserial,
    name varchar NOT NULL,
    address varchar,
    created_at Date NOT NULL,
    updated_at Date NOT Null,
    deleted_at Date,
    PRIMARY KEY (owner_id)
);

CREATE TABLE pet_categories(
    pet_category_id bigserial,
    name varchar NOT NULL,
    created_at Date NOT NULL,
    updated_at Date NOT Null,
    deleted_at Date,
    PRIMARY KEY (pet_category_id)
);

CREATE TABLE pets(
    pet_id bigserial,
    name varchar NOT NULL,
    date_of_birth date,
    is_aggressive boolean NOT NULL,
    pet_category_id bigint NOT NULL,
    created_at Date NOT NULL,
    updated_at Date NOT Null,
    deleted_at Date,
    PRIMARY KEY (pet_id), 
    FOREIGN KEY (pet_category_id) REFERENCES pet_categories(pet_category_id)
);

CREATE TABLE owner_pets(
    owner_pet_id bigserial,
    owner_id bigint NOT NULL,
    pet_id bigint NOT NULL,
    start_date date NOT NULL,
    end_date date,
    created_at Date NOT NULL,
    updated_at Date NOT Null,
    deleted_at Date,
    PRIMARY KEY (owner_pet_id),
    FOREIGN KEY (owner_id) REFERENCES owners(owner_id),
    FOREIGN KEY (pet_id) REFERENCES pets(pet_id)
);