# management

http://127.0.0.1:8080/api/user/create

{
    "username":"aa",
    "password":"123",
    "age":18,
    "tel":"1355",
    "addr":"tty",
    "card":"dasdasda1"
}

http://127.0.0.1:8080/api/user/update

{
    "username":"sssssssss",
    "password":"123",
    "age":18,
    "tel":"1355",
    "addr":"tty",
    "card":"dasdasda1",
    "id":1
}

http://127.0.0.1:8080/api/user/delete

{
    "id":1
}

http://127.0.0.1:8080/api/user/list
{"page_no":0}


http://127.0.0.1:8080/api/user/batchUpdate
[{
    "username":"aa",
    "password":"666",
    "age":18,
    "tel":"1355",
    "addr":"tty",
    "id":1
},
{
    "username":"bb",
    "password":"666",
    "age":18,
    "tel":"1355",
    "addr":"tty",
    "id":2
},
{
    "username":"cc",
    "password":"666",
    "age":18,
    "tel":"1355",
    "addr":"tty",
    "id":3
}
]