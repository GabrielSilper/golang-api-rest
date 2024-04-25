CREATE TABLE artists (
	id serial primary key,
  name varchar not null,
  nationality varchar(30) not null
);

insert into
  artists (name, nationality)
values
  ('Jimi Hendrix', 'American'),
	('Adele', 'British'),
	('Tupac Shakur', 'American'),
	('Wolfgang Amadeus Mozart', 'Austrian');