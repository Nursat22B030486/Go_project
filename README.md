# GO PROJECT  ->  Read It <BR>

****
#### Read It is a platform where you can read articles of another users, leave a comment and save it as a favorite to read it later. Also, you can add your own article, edit and remove it as well.  
---

## REST API
<br> 

#### Helath check
   * `GET /healthcheck`
#### User Registration
   * `POST /register`
#### User Activation
   * `PUT /activate`
#### User Login
 * `POST /login`
#### Endpoints which related with articles (Entity #1)
 * `POST /articles` 
 * `GET  /articles`
 * `GET  /articles/:id`
 * `PUT  /article/:id`
 * `DELETE /articles/:id`

#### Endpoints which related with comments (Entity #2)
 * `POST /articles/:article_id/comments` 
 * `GET  /articles/:article_id/comments`
 * `GET  /articles/:article_id/comments/:comment_id`
 * `PUT  /articles/:article_id/comments/:comment_id`
 * `DELETE /articles/:article_id/comments/:comment_id`


## Database Structure

``` sql
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table Users{
  user_id serial [primary key]
  full_name varchar(80)
  username varchar(50) [not null] 
  password varchar(30) [not null]
  created_at timestamp 
}

Table Articles {
  id serial [primary key]
  title varchar(255)
  author_id serial [not null, ref: > Users.user_id]
  genre varchar(50)
  body text 
  created_at timestamp
  updated_at timestamp
}

// many-to-many
Table User_saves {
  id serial [primary key]
  user_id serial [not null, ref: > Users.user_id]
  article_id serial [not null, ref: > Articles.id]
}

Table Comments {
  id serial [primary key]
  article_id integer [not null, ref: > Articles.id]
  user_id  integer [not null, ref: > Users.user_id]
  body text [not null]
  created_at timestamp 
}
```


### Kordabay Nursat  22B030486
