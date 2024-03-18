CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

CREATE TABLE actors
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gender VARCHAR(50),
    birthdate DATE
);

CREATE TABLE movies
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT,
    release_date DATE,
    rating DECIMAL(2, 1) CHECK (rating >= 0 AND rating <= 10)
    actor int references actors (id) on delete cascade not null,
);

CREATE TABLE movies_actors
(
    movie_id INT REFERENCES movies(id) ON DELETE CASCADE,
    actor_id INT REFERENCES actors(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, actor_id)
);




