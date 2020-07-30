# Notre Dame

This project implements 3 main services, namely: *Warehouse Cloning*, *Warehouse API*, and *Datamart API*. The services split over multiple Docker containers whose names are referred to as in `docker-compose.yml`. The container `warehouse_cloning` downloads the data from [finnhub](https://finnhub.io/) and store it in `warehouse_mongo` for further use. Users can access the downloaded data locally from  `warehouse_mongo` via `warehouse_api`. The container `datamart_api` stores proprietary data e.g. financial scores in `datamart_mongo` and allow users to retrieve them systematically. Optionally, the containers `warehouse_mongo_express` and `datamart_mongo_express` provide web-based database management systems for quick administration and maintenance.

## Outline
- [Installation](#installation)
- [Configuations](#configuations)
    - [Warehouse Cloning](#warehouse-cloning-configuration) 
    - [Warehouse API](#warehouse-api-configuration) 
    - [Datamart API](#datamart-api-configuration) 
- [APIs](#apis)
    - [Warehouse API](#warehouse-api-reference) 
    - [Datamart API](#datamart-api-reference) 
- [Managing databases](#managing-databases)
    - [Warehouse Mongo Express](#warehouse-mongo-express)
    - [Datamart Mongo Express](#datamart-mongo-express)

## Installation
Should there be changes to the source files,  commit and push to *this repository* before proceed to the next step. Note that the data in the databases and the system logs are stored on the host machine (outside the containers), specifically at `~/data` and `~/logs`. Therefore, the data is not lost when a newer version is installed.  These locations can be changed in `docker-comppse.yml`

1. Install [docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/)
2. Clone (or pull) source files from *this repository*
3. Add [config.yaml](#configuations) to each microservice directory
4. In `./notredame`, run docker-compose:
```bash
$ sudo docker-compose up -d
```
5. To stop all the microservices,
 ```bash
$ sudo docker-compose down
```

## Configuations

### Warehouse cloning configuration
```yaml
source:
    host: "https://finnhub.io/api/v1/"      # endpoint host
    token: "token"                          # authentication token
    consumers: 20                           # number of concurrent request processors
    wait: 60                                # waiting time (seconds) when too_many_request encountered
    attempts: 10                            # attempts (times) when too_many_request encountered

target:
    database: warehouse_mongo               # database name
    host: "mongodb://warehouse_mongo:27017" # database host on docker network
    username: admin                         # username
    password: admin                         # password
    
exchanges:                                  # list of exchanges of interest
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

documents:                                # list of documents of interest
    - "profile"
    - "financials"
    - "candle"

logging:
    level: debug                          # log level {debug, error}
    stdout: true                          # log to stdout {true, false}
    dirname: "/logs"                      # log directory
    
```

### Warehouse API configuration
```yaml
source:
    database: warehouse_mongo             # database name
    host: mongodb://warehouse_mongo:27017 # database host on docker network
    username: admin                       # username
    password: admin                       # password

target:
    host: 0.0.0.0:1323

authen:
    usernames:                            # list of authorized users
        - dome                            # username
        - blank                           # username
        - yort                            # username
    secret: 
        - secret                          # secret for token generation
    expire: 2030-01-01                    # token expiration date (yyyy-mm-dd)

logging:
    level: debug                          # log level {debug, error}
    stdout: true                          # log to stdout {true, false}
    dirname: "/logs"                      # log directory
```

### Datamart API configuration
```yaml
source:
    database: datamart_mongo              # database name
    host: mongodb://datamart_mongo:27017  # database host on docker network
    username: admin                       # username
    password: admin                       # password

target:
    host: 0.0.0.0:1323                    # server binding port

authen:
    usernames:                            # list of authorized users
        - dome                            # username
        - blank                           # username
    secret: 
        - secret                          # secret for token generation
    expire: 2030-01-01                    # token expiration date (yyyy-mm-dd)

logging:
    level: debug                          # log level {debug, error}
    stdout: true                          # log to stdout {true, false}
    dirname: "/logs"                      # log directory
```

Note that the port **1323** on the Docker container is mapped to the port **1324** on the host machine as defined in `docker-compose.yml`.

## APIs
### Warehouse API reference

Use the following to retrieve an authentication token:
> GET /token?username={username}

where {username} is one of the [authorized users](#warehouse-api-configuration) 

The following endpoints require `username` and authentication `token` attached in the request **header**.

To retrieve a list of securities available in an exchange:
> GET /api/symbols?exchange={exchange}

To retrieve a company profile:
> GET /api/profile?exchange={exchange}&symbol={symbol}

To retrieve a financial statement:
> GET /api/financials/{statement}/{frequency}?exchange={exchange}&symbol={symbol}

where {statement} can be one of `bs, ic, cf` representing balance sheet, income statement, and cash flow statement, respectively, and {frequency} can be `annual, quarterly, ttm, ytd`.

To retrieve the historical prices (OHLC):
> GET /api/candle?exchange={exchange}&symbol={symbol}

To search for symbols:
> GET /api/search?symbol={symbol}&text={text}&limit={limit}

- If only {symbol} is provided, return securities whose symbol contains {symbol} as a substring. 
- If only {text} is provided, return securities whose symbol **or** description contains {text} as a substring. 
- If {symbol} and {text} are provided, return securities whose symbol contains {symbol} as a substring **and** description contains {text} as a substring. 
* If {limit} is provided, return a maximum of {limit} securities, ordered alphabetically by symbols.

### Datamart API reference

Use the following to retrieve an authentication token:
> GET /token?username={username}

where {username} is one of the [authorized users](#datamart-api-configuration) 

The following endpoints require `username` and authentication `token` attached in the request **header**.

To upload the data:
> PUT /api/replace?expert={expert}&tag={tag}
> PUT /api/update?expert={expert}&tag={tag}

where {expert} is the name of the human expert publishing the data, and {tag} is used for versioning the data. When data is uploaded, it will be tagged with {tag} and additionally `latest`.  The data must be a list of JSON objects attached to the request **body**, for example:
```JSON
    [
        {
            "exchange": "us",
            "symbol": "appl",
            "data": {
                "dates": [...],
                "scores": [...]
            }
        },
        {
            "exchange": "us",
            "symbol": "amzn",
            "data": {
                "dates": [...],
                "scores": [...]
            }
        },
    ]
```
The difference between `/api/replace` and `/api/update` arises from the different mechanisms between [replaceOne](https://docs.mongodb.com/manual/reference/method/db.collection.replaceOne/) and [updateOne](https://docs.mongodb.com/manual/reference/method/db.collection.updateOne/) of MongoDB.

To download the data for securities associated with the expert and tag:
> GET /api/find?expert={expert}&tag={tag}&exchange={exchange}&symbol={symbol}

* The parameter {expert} is required.
* The parameter {tag} is optional and set to `latest` if omitted.
* If {exchange} and {symbol} optional, apply as additional filters when provided.

## Managing Databases

### Warehouse Mongo Express
*Warehouse Cloning* and *Warehouse API* share the same database `warehouse_mongo` which (by default) is binded to port 27017 on the host machine. To access the web-based database management system, reach the host machine on port **8081**.

### Datamart Mongo Express
*Datamart API* use the database `datamart_mongo` which (by default) is binded to port 27018 on the host machine. To access the web-based database management system, reach the host machine on port **8082**.
