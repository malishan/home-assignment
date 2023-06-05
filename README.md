# Assignment Service
Assignment service provides REST apis for fetch response from a dependent api.

- [Swagger](http://127.0.0.1:8080/swagger/index.html#/)

## Functional Requirement
### Health Api
* **Database Status**: ping database for health
### Assignment Api
* **Bored Api Response**: return 3 pairs of distinct key and assignment values
### Scheduler
* **Log Distinct Keya**: every 15 mins periodically print all distinct keys returned from boredapi

## Health Api
* api returns the status of database
* use the following curl to hit health api
```shell
curl --location 'http://localhost:8080/health/v1/status'
```

## Assignment Api
* api returns 3 unique key and assignment pair from boredapi.com
* if duplicate keys are returned from boredapi.com then assignment api also throws duplicate error to end client
* use the following curl to hit assignment api
```shell
curl --location 'http://localhost:8080/home/v1/activities'
```

## Scheduler Logs
* every 15 mins scheduler picks all distinct keys returned from boredapi.com and was saved in db
* logs are printed in file
* use the following log format to get those logs
```shell
ps aux | grep "PollActivityOperation"
```
* example as follows 
```json
{"level":"info","path":"Cron-PollActivity","responseBody":"[{\"count\":3,\"key\":\"8550768\"},{\"count\":2,\"key\":\"4266522\"},{\"count\":2,\"key\":\"6825484\"},{\"count\":2,\"key\":\"7091374\"},{\"count\":2,\"key\":\"3456114\"}]","time":"2023-06-05T02:15:00+05:30","message":"PollActivityOperation: success result"}
```


## Checks and Limitations
* gracefull shutdown of resources are handled


## How to run the application directly on local?

To run the application, you need to use the following instructions
1. Run **go build -o main** command to build the service 
2. Run **./main --port 8080 --env dev --base-config-path ./resources** command to start the service 
3. Note: the configurations are present in the env directory inside the base config path directory
4. Note: logs will be printed in the mentioned log path as per the config provided. Log rotation is also handled.


## How to run the application via docker on local?

To run the application, you need to use the following instructions
1. Facing problem staring the postgres dependency container



## What to do when you add a new API or change an existing one?

You will have to provide proper comments to that API as per the documentation mentioned [here](https://github.com/swaggo/swag#general-api-info). Once you have done that, you will have to run the following command to update the documentation.
```shell
make swagger
```
That's it, now just run the application and browse `http://localhost:8080/swagger/index.html` to view your changes.

## How to run tests?
* run the following make command
* test cover profile is generated which can be viewed in any browser
```shell
make test
```