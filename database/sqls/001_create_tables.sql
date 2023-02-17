CREATE TABLE IF NOT EXISTS Students(
    id CHAR(30) NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(100)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS Tags(
    id CHAR(30) NOT NULL PRIMARY KEY,
    tagname VARCHAR(64) NOT NULL
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS Books(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(64) NOT NULL,
    title_kana VARCHAR(128),
    tagid VARCHAR(64) NOT NULL,
    ISBN BIGINT(15) NOT NULL,
    author VARCHAR(40),
    author_kana VARCHAR(80),
    publisher VARCHAR(60),
    item_caption VARCHAR(500),
    sales_date DATE,
    image_url VARCHAR(500)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS Rent_hist(
    id INT AUTO_INCREMENT PRIMARY KEY,
    bookid INT, 
    renterid VARCHAR(30),
    return_date DATE,
    rent_date DATE,
    returned boolean DEFAULT false,
    FOREIGN KEY fk_renter_id(renterid) REFERENCES Students(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY fk_book_id(bookid) REFERENCES Books(id) ON DELETE CASCADE ON UPDATE CASCADE
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
