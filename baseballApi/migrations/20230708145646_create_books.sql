-- Active: 1681394654970@@127.0.0.1@3306@mydb
-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
   id INT PRIMARY KEY AUTO_INCREMENT,
   title VARCHAR(255),
   author VARCHAR(255),
   publisher VARCHAR(255),
   publication_year INT,
   genre VARCHAR(255)
);

CREATE TABLE users (
   id INT PRIMARY KEY AUTO_INCREMENT,
   name VARCHAR(255),
   address VARCHAR(255),
   contact_number VARCHAR(255)
);

CREATE TABLE loans (
   id INT PRIMARY KEY AUTO_INCREMENT,
   book_id INT,
   user_id INT,
   loan_date DATE,
   due_date DATE,
   return_date DATE,
   FOREIGN KEY (book_id) REFERENCES books(id),
   FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE loans;

DROP TABLE users;

DROP TABLE books;

-- +goose StatementEnd