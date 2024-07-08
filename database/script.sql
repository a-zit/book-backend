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

