update 'users'
set 'username'     = '柳亦钦',
    'password'     = 'f7e0babe54bd24b5a77e868e3fdbab4809c9b83b12cbc21c5b630ef5a2bddae1',
    'phone'        = '16688886668',
    'email'        = 'www.liuyiqin@mail.com.cn',
    'money'        = 0.00,
    'photo'        = '',
    'shopping_car' = '',
    'address'= ''
where user_id = '1';
select user_id,
       username,
       password,
       phone,
       email,
       money,
       photo,
       shopping_car,
       address
from users
where user_id = '3';
