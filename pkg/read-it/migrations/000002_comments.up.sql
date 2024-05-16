CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  article_id INTEGER NOT NULL REFERENCES Articles(id),
  user_id INTEGER NOT NULL REFERENCES Users(id),
  body TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);
