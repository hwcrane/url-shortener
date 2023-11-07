DROP TABLE IF EXISTS urls;
CREATE TABLE urls(
                       key      VARCHAR(128) NOT NULL,
                       url     VARCHAR(255) NOT NULL,
                       PRIMARY KEY (`key`)
);