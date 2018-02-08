DROP TABLE IF EXISTS video;
CREATE TABLE video
(
  id INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT,
  video_key VARCHAR(255) UNIQUE,
  title VARCHAR(255) NOT NULL,
  status TINYINT DEFAULT 1,
  duration INT UNSIGNED DEFAULT 0,
  url VARCHAR(255) NOT NULL,
  thumbnail_url VARCHAR(255),
  PRIMARY KEY(id)
);