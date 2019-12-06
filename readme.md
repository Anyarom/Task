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
        "Id": 4037200794235010051,
        "Status": 405,
        "Headers": [
            {
                "Key": "Content-Length",
                "Value": "1589"
            },
            {
                "Key": "X-Xss-Protection",
                "Value": "0"
            },
            {
                "Key": "X-Frame-Options",
                "Value": "SAMEORIGIN"
            },
            {
                "Key": "Allow",
                "Value": "GET, HEAD"
            },
            {
                "Key": "Date",
                "Value": "Sun, 01 Dec 2019 20:46:17 GMT"
            },
            {
                "Key": "Content-Type",
                "Value": "text/html; charset=UTF-8"
            },
            {
                "Key": "Server",
                "Value": "gws"
            }
        ],
        "Length": 1589
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
      "ReqTasks": [
          {
              "Method": "POST",
              "Address": "http://google.com",
              "Headers": [
                  {
                      "Key": "key1",
                      "Value": "val1"
                  },
                  {
                      "Key": "key2",
                      "Value": "val2-"
                  }
              ],
              "Body": "task1"
          },
          {
              "Method": "POST",
              "Address": "http://google.com",
              "Headers": [
                  {
                      "Key": "key1",
                      "Value": "val1"
                  },
                  {
                      "Key": "key2",
                      "Value": "val2"
                  }
              ],
              "Body": "task1"
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
