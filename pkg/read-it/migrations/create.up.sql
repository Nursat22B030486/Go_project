CREATE TABLE IF NOT EXISTS "Users" (
  "user_id" serial PRIMARY KEY,
  "full_name" varchar(80),
  "username" varchar(50) NOT NULL UNIQUE ,
  "password" varchar(30) NOT NULL UNIQUE ,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS "Articles" (
  "id" serial PRIMARY KEY,
  "title" varchar(255),
  "author_id" serial NOT NULL,
  "genre" varchar(50),
  "body" text,
  "created_at" timestamp    NOT NULL DEFAULT NOW(),
  "updated_at" timestamp    NOT NULL DEFAULT NOW(),
  FOREIGN KEY (author_id) references "Users"(user_id)
);


CREATE TABLE IF NOT EXISTS "User_saves" (
  "id" serial PRIMARY KEY,
  "user_id" serial NOT NULL ,
  "article_id" serial NOT NULL,
  foreign key (user_id) references "Users"(user_id),
  foreign key (article_id) references "Articles"(id)
);




