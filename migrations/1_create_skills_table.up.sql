CREATE TABLE IF NOT EXISTS skills
(
    id serial PRIMARY KEY,
    category varchar(100) NOT NULL,
    name varchar(150) NOT NULL,
    exp int NOT NULL,
    member_id int NOT NULL
);

INSERT INTO skills
VALUES
  ('1', 'IT', 'PHP', '7', '1'),
  ('2', 'IT', 'JAVA', '10', '2'),
  ('3', 'Design', 'DD', '2', '1');