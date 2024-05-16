
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone not null default now(),
  name text not null,
  email citext unique not null,
  password_hash bytea not null,
  activated bool not null,
  version integer not null default 1
);


CREATE TABLE IF NOT EXISTS tokens (
    hash bytea PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    expiry timestamp(0) with time zone NOT NULL,
    scope text NOT NULL
);



CREATE TABLE IF NOT EXISTS articles (
  id serial PRIMARY KEY,
  title varchar(255) not null,
  author_id integer not null,
  genre varchar(100) not null,
  body text UNIQUE NOT NULL ,
  created_at timestamp    NOT NULL DEFAULT NOW()
  , updated_at timestamp    NOT NULL DEFAULT NOW()
  , FOREIGN KEY (author_id) references users(id)
);



CREATE TABLE IF NOT EXISTS user_saves (
  id serial PRIMARY KEY,
  user_id serial NOT NULL ,
  article_id serial NOT NULL,
  foreign key (user_id) references users(id),
  foreign key (article_id) references articles(id)
);


CREATE TABLE IF NOT EXISTS permissions (
    id bigserial PRIMARY KEY,
    code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);

-- Add the two permissions to the table.
INSERT INTO permissions (code)
    VALUES
    ('articles:read'),
    ('articles:write');

--
-- insert into articles (title, genre, body)
-- values ('First', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Second', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Third', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Forth', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Fivth', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Sixth', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Seventh', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
--        ('Eighth', 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw');
--
-- insert into articles (title, genre, body)
--     values ('Tenth', 'comedy', 'fekdjfdnwekn');




