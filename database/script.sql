CREATE TABLE Books (
    ID BIGINT NOT NULL AUTO_INCREMENT,
    Title VARCHAR(255) NOT NULL,
    Price DECIMAL(10, 2) NOT NULL,
    ImageUrl VARCHAR(255),
    Author VARCHAR(255) NOT NULL,
    Genre VARCHAR(255) NOT NULL,
    YearOfPublication INT NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    
    PRIMARY KEY (ID)
);

ALTER TABLE Books ADD CONSTRAINT unique_title_author UNIQUE (Title, Author);

-- create sample of data
INSERT INTO Books (Title, Price, ImageUrl, Author, Genre, YearOfPublication, CreatedAt, UpdatedAt) VALUES
('The Great Gatsby', 10.00, 'https://img.freepik.com/free-vector/modern-annual-report-business-flyer-template-design_1017-25864.jpg?t=st=1719740351~exp=1719743951~hmac=6107f3de1cbb0b39806aba0af459568e954eb5850f47c659c4ac4204fdfd583e&w=1480', 'F. Scott Fitzgerald', 'Fiction', 1925, NOW(), NOW()),
('To Kill a Mockingbird', 15.00, 'https://img.freepik.com/free-vector/modern-annual-report-business-flyer-template-design_1017-25864.jpg?t=st=1719740351~exp=1719743951~hmac=6107f3de1cbb0b39806aba0af459568e954eb5850f47c659c4ac4204fdfd583e&w=1480', 'Harper Lee', 'Fiction', 1960, NOW(), NOW()),
('1984', 20.00, 'https://img.freepik.com/free-vector/modern-annual-report-business-flyer-template-design_1017-25864.jpg?t=st=1719740351~exp=1719743951~hmac=6107f3de1cbb0b39806aba0af459568e954eb5850f47c659c4ac4204fdfd583e&w=1480', 'George Orwell', 'Fiction', 1949, NOW(), NOW()),
('Pride and Prejudice', 25.00, 'https://img.freepik.com/free-vector/modern-annual-report-business-flyer-template-design_1017-25864.jpg?t=st=1719740351~exp=1719743951~hmac=6107f3de1cbb0b39806aba0af459568e954eb5850f47c659c4ac4204fdfd583e&w=1480', 'Jane Austen', 'Fiction', 1813, NOW(), NOW()),
('The Catcher in the Rye', 30.00, 'https://img.freepik.com/free-vector/modern-annual-report-business-flyer-template-design_1017-25864.jpg?t=st=1719740351~exp=1719743951~hmac=6107f3de1cbb0b39806aba0af459568e954eb5850f47c659c4ac4204fdfd583e&w=1480', 'J.D. Salinger', 'Fiction', 1951, NOW(), NOW()); 