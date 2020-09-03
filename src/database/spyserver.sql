CREATE TABLE domain (
	domain STRING,
	servers_changed BOOL,
	ssl_grade STRING,
	previous_ssl_grade STRING,
	logo STRING,
	title STRING,
	is_down BOOL,
	updated STRING,
	PRIMARY KEY ("domain")
);

CREATE TABLE server (
	address STRING,
	domain STRING NOT NULL REFERENCES domain(domain) ON DELETE CASCADE,
	ssl_grade STRING,
	country STRING,
	owner STRING,
	PRIMARY KEY (address, domain)
);