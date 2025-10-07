CREATE TABLE books(
    id_books SERIAL PRIMARY KEY,
    id_author INTEGER,
    name_books VARCHAR(255) NOT NULL,
    age INTEGER,
    FOREIGN KEY (id_author) REFERENCES author(id) ON DELETE CASCADE
);