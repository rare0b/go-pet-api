CREATE TYPE pet_status AS ENUM ('available', 'pending', 'sold');

CREATE TABLE categories
(
    category_id   INT,
    category_name TEXT,
    CONSTRAINT pk_categories PRIMARY KEY (category_id)
);

CREATE TABLE pets
(
    pet_id      INT,
    category_id INT,
    pet_name    TEXT   NOT NULL,
    photo_urls  TEXT[] NOT NULL,
    status      pet_status,
    CONSTRAINT pk_pets PRIMARY KEY (pet_id),
    FOREIGN KEY (category_id) REFERENCES categories (category_id)
);

CREATE TABLE tags
(
    tag_id   INT,
    tag_name TEXT,
    CONSTRAINT pk_tags PRIMARY KEY (tag_id)
);

CREATE TABLE pet_tags
(
    pet_id INT NOT NULL,
    tag_id INT NOT NULL,
    PRIMARY KEY (pet_id, tag_id),
    FOREIGN KEY (pet_id) REFERENCES pets (pet_id),
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id)
);
