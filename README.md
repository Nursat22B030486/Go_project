# GO PROJECT  ->  Read It <BR>

****
#### Read It is a platform where you can read articles of another users and save it as a favorite to read it later. Also, you can add your own article, edit and remove it as well.  
---

## REST API
<br> 

 * `POST /articles` 
 * `GET  /articles/saves`
 * `GET  /articles/:genre`
 * `PUT  /article/:id`
 * `UPDATE /articles/:id`
 * `DELETE /articles/:id`

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
```


### Kordabay Nursat  22B030486