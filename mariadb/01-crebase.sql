use music;

DROP TABLE IF EXISTS tracks;
DROP TABLE IF EXISTS records;
DROP TABLE IF EXISTS artists;

CREATE TABLE artists (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(100),
    PRIMARY KEY (id)
);

CREATE TABLE records (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    id_artist INT,
    year DECIMAL UNSIGNED,
    genre VARCHAR(255),
    support VARCHAR(30),
    nb_support DECIMAL UNSIGNED,
    label VARCHAR(255),
    PRIMARY KEY (id),
    CONSTRAINT fk_artist
        FOREIGN KEY (id_artist) REFERENCES artists (id)
        ON DELETE CASCADE
        ON UPDATE RESTRICT
);

CREATE TABLE tracks (
    id INT NOT NULL AUTO_INCREMENT,
    id_record INT NOT NULL,
    number DECIMAL UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    nb_support DECIMAL UNSIGNED,
    PRIMARY KEY (id),
    CONSTRAINT fk_record
        FOREIGN KEY (id_record) REFERENCES records (id)
        ON DELETE CASCADE
        ON UPDATE RESTRICT
);