# Monorepo Backend Services
Go Commerce uses a monorepo service and uses the Golang language.

**Table of Contents**

- [Monorepo Backend Services](#monorepo-backend-services)
- [Project Explain](#explain)
- [Reference Project](#reference)
- [Installation](#installation)
- [Prequisites](#prequisites)
- [Installation](#installation)
- [Run Project](#how-to-run)
- [Port HTTP](#port-http)
- [Port RPC](#port-rpc-services)
- [Next Development](#next)

## Explain
1. Use gRPC as gateway And service
2. Use buf for generate proto
3. Use dbq as sql library because speed is important
4. Build query builder for dynamic query options
5. Structure code is base on clean architecture
6. Still not use unit test because i'm still learning write unit test on golang
7. Run project must 1 on 1 terminal with file run.sh because still learning use to compose or kube for run monorepo golang project

## Reference
1. https://miro.medium.com/max/1186/1*GTzwD1BH-jxlK9X7olDZOw.png
2. https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5
3. https://lbhackney-it.github.io/API-Playbook/clean_architecture/

## Prequisites
1. Golang v1.18.4 or above
2. MySQL DBMS
3. Using Sonarlint extension for quality code

## Installation

1. copy environment by services `cp .env.example .env`
2. sync go modules `go mod tidy`
3. install migration for go https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
4. sh migration-mysql.sh servicename


## How to run
1. sh run-service.sh servicename (ex : sh run-service.sh api)
2. sh run-service.sh api (gateway)
3. sh run-service.sh product (service)

## Port http
| Service Name | Port     |
| :----------- | :------- |
| api          | 8001     |

## Port RPC Services

| Service Name | Port RPC | DBMS       |
| :----------- | :------- | :--------- |
| user         | 7001     | mysql      |
| product      | 7002     | mysql      |
| transaction  | 7003     | mysql      |

## Next
1. Create feature user and authentication
2. Create feature transaction
3. Want research how implement a good flash sale on ecommerce example 1million request in same time
4. Add logger for debug
5. Research and implement monitoring for golang project