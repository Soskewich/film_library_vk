CREATE TABLE actor 
(
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "patronymic" VARCHAR(255),
    "birthday" date NOT NULL,
    "gender" VARCHAR(7) NOT NULL
);

CREATE TABLE film
(
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(150),
    "description" VARCHAR(1000) NOT NULL,
    "dateExit" DATE NOT NULL,
    "rating" INT
);

CREATE TABLE film_actor
(
    id        BIGSERIAL PRIMARY KEY,
    actor_id  INTEGER NOT NULL REFERENCES actor,
    film_id   INTEGER NOT NULL REFERENCES film,
    UNIQUE (actor_id, film_id)
);