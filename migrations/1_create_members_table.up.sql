CREATE TABLE IF NOT EXISTS members
(
    id serial PRIMARY KEY,
    name varchar(150) NOT NULL
);

INSERT INTO members
VALUES
  ('1', 'DaiDV'), ('2', 'DinhLV');