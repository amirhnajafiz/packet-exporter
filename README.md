# Zap Operator

Zap is a blazing fast, structured, leveled logging in Go. Our service's logs are in JSON format.
In this project I create an operator to get these logs from OKD's pods ```stdout```, and publish
them over the following topics over a ```NATS``` cluster.

1. ```{service-name}.logs.debug```
2. ```{service-name}.logs.warning```
3. ```{service-name}.logs.info```
4. ```{service-name}.logs.error```
5. ```{service-name}.logs.unknown```
