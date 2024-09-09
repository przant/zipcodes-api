DROP DATABASE IF EXISTS us_zipcodes;

CREATE DATABASE us_zipcodes;
USE us_zipcodes;

DROP TABLE IF EXISTS zipcodes;
CREATE TABLE zipcodes(
    id         INTEGER AUTO_INCREMENT NOT NULL,
    state      VARCHAR(15) NOT NULL,
    state_abbr CHAR(2) NOT NULL,
    zipcode    INTEGER NOT NULL,
    county     VARCHAR(25) NOT NULL,
    city       VARCHAR(25) NOT NULL,
    PRIMARY KEY (id)
);