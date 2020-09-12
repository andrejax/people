CREATE TABLE IF NOT EXISTS users_group (
	id int(64) NOT NULL AUTO_INCREMENT,
	name char(20),
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user (
	id int(64) NOT NULL AUTO_INCREMENT,
	name char(20) NOT NULL,
	email char(20) NOT NULL,
	password char(20) NOT NULL,
	group_id int(64),
	PRIMARY KEY (id),
	FOREIGN KEY (group_id) REFERENCES users_group(id) ON DELETE SET NULL
)