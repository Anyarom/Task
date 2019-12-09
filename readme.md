# API Fetch Task

URL: http://127.0.0.1:8080

## Запрос к другим ресурсам

От клиента получаем запрос обратиться к другим ресурсам.

#### Request:

POST /task

BODY:

````json
{
    "method": "POST",
    "address": "http://google.com",
    "headers": [
      {
        "key": "key1",
        "value": "val1"
      },
      {
        "key": "key2",
        "value": "val2"
      }
    ],
    "body": "task1"
  }
````
  
#### Response  
200

````json
  {
      "request_id": 1,
      "response": {
          "status": 405,
          "headers": [
              {
                  "key": "Date",
                  "value": "Sun, 08 Dec 2019 17:27:32 GMT"
              },
              {
                  "key": "Content-Type",
                  "value": "text/html; charset=UTF-8"
              },
              {
                  "key": "Server",
                  "value": "gws"
              },
              {
                  "key": "Content-Length",
                  "value": "1589"
              },
              {
                  "key": "X-Xss-Protection",
                  "value": "0"
              },
              {
                  "key": "X-Frame-Options",
                  "value": "SAMEORIGIN"
              },
              {
                  "key": "Allow",
                  "value": "GET, HEAD"
              }
          ],
          "length": 1589
      }
  }
````
    
## Запрос на получение всех запросов к другим ресурсам

От клиента получаем запрос выдать все запросы

#### Request

GET /tasks

#### Response

200

````json
{
    "tasks": [
        {
            "request_id": 1,
            "request": {
                "method": "POST",
                "address": "http://google.com",
                "headers": [
                    {
                        "key": "key1",
                        "value": "val1"
                    },
                    {
                        "key": "key2",
                        "value": "val2"
                    }
                ],
                "body": "task1"
            }
        },
        {
            "request_id": 2,
            "request": {
                "method": "POST",
                "address": "http://google.com",
                "headers": [
                    {
                        "key": "key1",
                        "value": "val1"
                    },
                    {
                        "key": "key2",
                        "value": "val2"
                    }
                ],
                "body": "task1"
            }
        }
    ]
}
````

## Запрос на удаление запроса к другим ресурсам

От клиента получаем запрос на удаление запроса

#### Request

DELETE /task?id=1

id - id запроса

#### Response

200
