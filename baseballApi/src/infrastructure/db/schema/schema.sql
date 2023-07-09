CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(255),
    publisher VARCHAR(255),
    publication_year INT,
    genre VARCHAR(255),
    isbn VARCHAR(255)
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    contact_number VARCHAR(255)
);

CREATE TABLE loans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    book_id INT,
    user_id INT,
    loan_date DATE,
    due_date DATE,
    return_date DATE,
    CONSTRAINT fk_books_loans FOREIGN KEY (book_id) REFERENCES books (id),
    CONSTRAINT fk_users_loans FOREIGN KEY (user_id) REFERENCES users (id)
);