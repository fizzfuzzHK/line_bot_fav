CREATE TABLE users (
  id int PRIMARY KEY AUTO_INCREMENT,
  user_id varchar(255) NOT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE places (
  id int PRIMARY KEY AUTO_INCREMENT,
  place_id int NOT NULL,
  name varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE fav (
  id int PRIMARY KEY AUTO_INCREMENT,
  user_id varchar(255)  NOT NULL,
  place_id int NOT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX users_index_0 ON users (user_id);

CREATE INDEX places_index_1 ON places (place_id);

CREATE INDEX fav_index_2 ON fav (user_id);

CREATE INDEX fav_index_3 ON fav (place_id);

ALTER TABLE fav ADD FOREIGN KEY (user_id) REFERENCES users (user_id);

ALTER TABLE fav ADD FOREIGN KEY (place_id) REFERENCES places (place_id);