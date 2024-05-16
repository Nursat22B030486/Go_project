CREATE ROLE root WITH LOGIN PASSWORD 'password';

select * from users;

select * from comments;
select * from tokens;

select * from articles;


-- Give faith@example.com the 'movies:write' permission
INSERT INTO users_permissions
    VALUES (
    (SELECT id FROM users WHERE email = 'n_kordabay@kbtu.kz'),
    (SELECT id FROM permissions WHERE code = 'articles:write')
    );

-- List all activated users and their permissions.
SELECT email,  array_agg(permissions.code) as permissions
FROM permissions
INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
INNER JOIN users ON users_permissions.user_id = users.id
GROUP BY email;

delete from articles where id=1;

insert into articles (title, author_id, genre, body)
values ('First', 3, 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
       ('Second', 4, 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcjw'),
       ('Third', 5, 'comedy', 'dfkwhuhbeieuhdvjejdcvjebcj'),
       ('Forth', 6, 'comedy', 'dfkwhuhbeieuhdvjejdcvjebc'),
       ('Fivth', 7, 'comedy', 'dfkwhuhbeieuhdvjejdcvjec'),
       ('Sixth', 4, 'comedy', 'dfkwhuhbeieuhdvjejdcvjebjw'),
       ('Seventh', 5, 'comedy', 'dfkwhuhbeieuhdvjejdebcjw'),
       ('Eighth', 6, 'comedy', 'dfkwhuhbeieuhvjebcjw');

insert into articles (title, author_id, genre, body)
    values ('Tenth', 1, 'comedy', 'fekdjfdnwekn');


select version();