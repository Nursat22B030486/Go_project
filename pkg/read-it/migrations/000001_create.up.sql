CREATE TABLE IF NOT EXISTS users (
  user_id serial PRIMARY KEY,
  full_name varchar(80),
  username varchar(50) NOT NULL UNIQUE ,
  password varchar(30) NOT NULL UNIQUE ,
  created_at timestamp NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS articles (
  id serial PRIMARY KEY,
  title varchar(255),
--   "author_id" serial ,
  genre varchar(100),
  body text,
  created_at timestamp    NOT NULL DEFAULT NOW(),
  updated_at timestamp    NOT NULL DEFAULT NOW()
--   FOREIGN KEY (author_id) references "Users"(user_id)
);

INSERT INTO articles(title, genre, body)
    VALUES  ('title 1', 'education', 'body no.1'),
            ('title 2', 'fantasy', 'body no.2'),
            ('title 3', 'weather', 'body no.3'),
            ('title 4', 'fantasy', 'body no.4'),
            ('title 5', 'education', 'body no.5'),
            ('title 6', 'technology', 'body no.6'),
            ('title 7', 'comedy', 'body no.7'),
            ('title 8', 'crime', 'body no.8'),
            ('title 9', 'drama', 'body no.9');

DELETE FROM articles WHERE id = 1;

select * from articles;

CREATE TABLE IF NOT EXISTS user_saves (
  id serial PRIMARY KEY,
  user_id serial NOT NULL ,
  article_id serial NOT NULL,
  foreign key (user_id) references users(user_id),
  foreign key (article_id) references articles(id)
);




