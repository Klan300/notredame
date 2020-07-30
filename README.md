# notredame
This is Document about Warehouse_api / Warehouse_cloning / Datamart_api. There was contain by docker file. Warehouse Cloning will clone data from finnhub and store into warehouse mongo and connect by warehouse api. Datamart api is api that collect scores and find scores from datamart mongo.
## Topic
- [**Project Setup**](##Setup)
- [**Config**](##Config)
- [**Api**](##Api)


## Setup
- **At first** If you have to fix something, You should to fix and push to github repository first 
- **Next** After you fix or change some thing and push it to github so you have to follow this step   

1. connect to server
```bash
$ssh ubuntu@18.141.209.89
``` 
2. Pull or Clone from repository
```bash
## git clone
$git clone 
## or git pull from repo
$git pull
```
3.  Add [**config.yaml**](##Config) to all folder 

4. install [**docker**](https://docs.docker.com/engine/install/) and [**docker-compose**](https://docs.docker.com/compose/install/)

5. run docker compose
```bash
$cd notredame
$docker-compose up -d
```
***
## Config
- [**Datamart Api**](###DatamartApi_Config)
- [**Warehouse Api**](###WarehouseApi_Config)
- [**Warehouse Cloning**](###WarehouseCloning_Config)

### DatamartApi_Config
**source**
> source database connection

```yaml
source:
    database: datamart_mongo #database name
    host: mongodb://datamart_mongo:27017  #database 
    username: username #username db
    password: password #password db
```
**target**
> target port to connect
```yaml
target:
    host: 0.0.0.0:1323 #port number
```
**authen**
> authentication for create token and user that can connect totken

```yaml
authen:
    usernames:
       - name #list of user
       - name #list of user
    secret:
        - secret #secret for token

    expire: yyyy-mm-dd #expire date format: yyyy-mm-dd
```
**log**
> log position and log level

```yaml
logging:
    level: level #level to log (debug,error)
    stdout: true #true,false log or not
    dirname: /dir/to/log #directory to log
```
***Example***

```yaml
source:
    database: datamart_mongo
    host: mongodb://datamart_mongo:27017
    username: username
    password: password

target:
    host: 0.0.0.0:1323

authen:
    usernames:
        - dome
        - blank
    secret:
        - secret

    expire: 2020-01-01

logging:
    level: debug
    stdout: true
    dirname: "/logs"
```
***
### WarehouseApi_Config

**source**
> source database connection

```yaml
source:
    database: warehouse_mongo #database name
    host: mongodb://warehouse_mongo:27017 #database port
    username: username #username db
    password: password #password db
```
**target**
> target port to connect
```yaml
target:
    host: 0.0.0.0:1323 #port number
```
**authen**
> authentication for create token and user that can connect totken

```yaml
authen:
    usernames:
       - name #list of user
       - name #list of user
    secret:
        - secret #secret for token

    expire: yyyy-mm-dd #expire date format: yyyy-mm-dd
```
**log**
> log position and log level

```yaml
logging:
    level: level #level to log (debug,error)
    stdout: true #true,false log or not
    dirname: /dir/to/log #directory to log
```

***Example***

```yaml
source:
    database: warehouse_mongo 
    host: mongodb://warehouse_mongo:27017 
    username: username 
    password: password 

target:
    host: 0.0.0.0:1323

authen:
    usernames:
        - dome
        - blank
    secret:
        - secret

    expire: 2020-01-01

logging:
    level: debug
    stdout: true
    dirname: "/logs"
```

***
### WarehouseCloning_Config

**source**
> source of api to connect

```yaml
source:
    host: "https://url/path/path" #path to connection
    token: "token" #token for authentication to api
    consumers: 20 #integer for number of consumer
    wait: 60 #time for wait when it return http res 429 
    attempts: 10 #limit when it found 429
```
**target**
> target database to collect data
```yaml
target:
    database: warehouse_mongo #database name
    host: "mongodb://warehouse_mongo:27017" #database port
    username: username #username db
    password: password #password db
```
**Exchange**
> exchange that you want to collect data

```yaml
exchanges: 
   - exchange1
   - exchange2
   - exchange3 #list of exchange
```
**documents**

>document that you want to get the data

```yaml
documents:
    - "profile" #CompanyProfile
    - "financials" #financial statement
    - "candle"  #Candle by daily
    # list of document
```
**log**
> log position and log level

```yaml
logging:
    level: level #level to log (debug,error)
    stdout: true #true,false log or not
    dirname: /dir/to/log #directory to log
```

***Example***

```yaml
source:
    host: "https://finnhub.io/api/v1/"
    token: "token"
    consumers: 20
    wait: 60
    attempts: 10

target:
    database: warehouse_mongo
    host: "mongodb://warehouse_mongo:27017"
    username: admin
    password: admin
    
exchanges: 
    - "US"
    - "BK"
    - "L"
    - "CN"
    - "T"
    - "HK"
    - "VN"
    - "AX"
    - "SS"
    - "SZ"
    - "SG"

documents:
    - "profile"
    - "financials"

logging:
    level: debug
    stdout: true
    Dirname: "/logs" 
```
***
## Api
- [**Token**](###Token)
- [**Datamart api**](###Datamart)
- [**Warehouse api**](###Warehouse)

### **Token**
- the token will use all of path that have **/api/** 
#### how to get token
- token for warehouse api use on port **1323**
```
http://18.141.209.89:1323/token?username={username}
```
- token for datamart api use on port **1324**
```
http://18.141.209.89:1324/token?username={username}
```
***

### **Datamart**
This is api for collect scores of stock data it have
- [**replace**](####Replace) for replace data in score
- [**update**](####Update) for update data in score
- [**find**](####Find) for find score data in database

#### Replace
> It was **PUT** method so you have to request and sent data in body with token
```
http://18.141.209.89:1324/api/replace?expert={expertname}&tag={tag version}
```
- you have to send with **Token**

- body must be **List** of **JSON** 

- in body should have 
    ```JSON
    [
        {
            "Exchange": "us",
            "Symbol": "appl",
            "Data": {
                "date": [],
                "scores": []
            }
        }
    ]
    ```

#### Update
> It was **PUT** method so you have to request and sent data in body with token
```
http://18.141.209.89:1324/api/update?expert={expertname}&tag={tag version}
```
- you have to send with **Token**

- body must be **List** of **JSON** 

- in body should have 
    ```JSON
    [
        {
            "Exchange": "us",
            "Symbol": "appl",
            "Data": {
                "date": [],
                "scores": []
            }
        }
    ]
    ```

#### Find
> This is **GET** method so you have to request in Correct path and query param you will recieve correct data

```
http://18.141.209.89:1324/api/find?tag={tag version}&expert={expertname}&exchange={exchange}&symbol={symbol}
```
- **expert** is fix query to find
- **tag** is not fix to use but if you not set tag vesion it will send **"lastest"** version
- **exchange & symbol** not fix to send if you not set it will return all data that match

***

### **Warehouse**
This is Api for collect data that collect from **finnhub** by **Warehouse cloning**
- [**Profile**](####Profile) **GET** company profile
- [**Financial**](####Financial) **GET** financial statement
- [**Candle**](####Candle) **GET** candle daily
- [**Search**](####Search) **GET** Company symbol and name by text or by symbol
- [**Symbol**](####Symbol) **GET** list of symbol from Exchange

#### Profile

```
http://18.141.209.89:1323/api/profile?symbol={symbol};&exchange={exchange}
```
- **symbol & exchange** are query params

#### Financial
```
http://18.141.209.89:1323/api/financials/{statement}/{frequency}?symbol={symbol}&exchange={exchange}
```
- **statement** have ```bs,ic,cf```
- **frequency** have ```annual, quarterly, ttm, ytd```
- **symbol & exchange** are query params
#### Candle
```
http://18.141.209.89:1323/api/candle?symbol={symbol}&exchange={exchange}
```
- **symbol & exchange** are query params
- **Data** was price in day by day but start from last 10 years
#### Search
```
http://18.141.209.89:1323/api/search?symbol={symbol}&text={text}&limit={limit}
```
- **limit** it can limit number of data when was search. If you not use it will send all data that match
- **text & symbol** if you use only text it will find name of company or symbol base on text however if you use with symbol or use only symbol it will find only symbol in data 
#### Symbol
```
http://18.141.209.89:1323/api/symbols?exchange={exchange}
```
- it use only **exchange** for find
- It will return with **list** of **symbol** 
***