# notredame
This is Document about Warehouse_api / Warehouse_cloning / Datamart_api. There was contain by docker file. Warehouse Cloning will clone data from finnhub and store into warehouse mongo and connect by warehouse api. Datamart api is api that collect scores and find scores from datamart mongo.
## Project setup
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

## Config
- [**Datamart Api**](###DatamartApi_Config)
- [**Warehouse Api**](###WarehouseApi_Config)
- [**Warehouse Cloning**](###WarehouseCloning_Config)

### DatamartApi_Config
### WarehouseApi_Config
### WarehouseCloning_Config



## Api