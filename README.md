# oidc

> 没读懂，所以代码写的很拉跨

- 主要功能大概是：

    1. 我在op有个号

    2. 我在op登录获取登录用的token

       > POST localhost:9090/user/login
       >
       > 结果：
       >
       > {
       >
       >   "msg": "获取token成功",
       >
       >   "status": 200,
       >
       >   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIwMDAwMDAiLCJleHAiOiIxNjc4NTA4NzgzIiwibmJmIjoiMTY3ODUwNTE4MyJ9.1bc47bcd328a3c3f4c5b3c9d7e8111ba811dada8c52baf3b271cf4fff8dcf18a"
       >
       > }

    3. 在op申请client_id,client_secret

       > GET localhost:9090/oidc/GetClientId
       >
       > 结果：
       >
       > {
       >
       >   "client_id": "1c3c032d7f211e49543b98b21e6f92de"
       >
       > }
       >
       > GET localhost:9090/oidc/GetClientSecret
       >
       > 结果：
       >
       > {
       >
       >   "clientSecret": "b76b9e4e0836cb681e9715d0a761100e"
       >
       > }

    4. 在rp用client_id,client_secret,附上redirect_uri向op发送请求获得到了我在op号里的信息（好像目的不是这个）重定向到redirect_uri里展示出来

       >GET localhost:9090/oidc/GetCode?clientId=1c3c032d7f211e49543b98b21e6f92de&clientSecret=b76b9e4e0836cb681e9715d0a761100e&redirectUri=http://localhost:7070/redirect_uri
       >
       >结果：(一个html页面)
       >
       ><!DOCTYPE html>
       >   <html lang="en">
       >      <head>
       >      <meta charset="UTF-8">
       >         <title>
       >            RP-HOMEPAGE
       >         </title>
       ></head>
       >      <body>
       >         data:<br>
       >map[email:000000@qq.com money:0.00 phone:13700000000 photo:http://localhost:9091/user/photo/2.jpg username:000000]
       >      </body>
       ></html>

    5. 代码里尝试了一下gorm和logrus，分层没认真搞，写的很乱，自己也看不下去，实在抱歉