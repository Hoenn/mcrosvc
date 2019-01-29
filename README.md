# mcrosvc
An example system that uses some common microservice architecture utilites.

## Quick Summary
This is an example app with a simple api and schema:

udb - provides a gRPC api layer to validate and interact with a mysql database

web-api - provides a CRUD HTTP api layer that creates gRPC calls to userdb
