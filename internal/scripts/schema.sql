CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE phone_numbers (
    id SERIAL PRIMARY KEY,
    contact_id INT REFERENCES contacts(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    number VARCHAR(100) NOT NULL
);
