


CREATE TABLE IF NOT EXISTS events (
    _id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(2000) NOT NULL,
    date DATE,
    location VARCHAR(500),
    max_participants INT
);


CREATE TABLE IF NOT EXISTS participations(
    firstname VARCHAR(50),
    event_id INT REFERENCES events(_id) ON DELETE CASCADE,
    lastname VARCHAR(50),
    email VARCHAR(50)
);