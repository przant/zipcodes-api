DROP DATABASE IF EXISTS us_zipcodes;

CREATE DATABASE us_zipcodes;
USE us_zipcodes;

DROP TABLE IF EXISTS zipcodes;
CREATE TABLE zipcodes(
    id         INTEGER AUTO_INCREMENT NOT NULL,
    state      VARCHAR(25) NOT NULL,
    state_abbr CHAR(2) NOT NULL,
    zipcode    VARCHAR(7) NOT NULL,
    county     VARCHAR(50) NOT NULL,
    city       VARCHAR(35) NOT NULL,
    PRIMARY KEY (id)
);