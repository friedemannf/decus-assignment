
# Decus Tech Assignment

**Please do not publicly fork this repository but rather make a private copy, thank you!**
  
This repository contains the boilerplate to a small microservice that is meant for you to show off your coding as well as general technical abilities.  
    
In here you will find two directories:  
 - `proto` contains different protobuf definitions:  
   - `service.proto` contains the protobuf service&message definitions for the ApikeyService that has to be implemented  
   - `errors.proto` contains general error code definitions that are meant to be implemented by different services, including the ApikeyService  
 - `assignment` contains a simple go application with the needed import structure for protobufs already set up.  
  
The goal of this technical assignment os to implement a small service that exposes a gRPC server interface for other services to interact with.  
The service revolves around several actions around apikeys. Decus' goal is to manage it's user's apikeys for custodians in a secure manner while still making them accessible for interacting with APIs.  
For the sake of this assignment, an apikey always consists of an "apikey" and a "secret". For both should be taken care to not expose them accidentally, but only when requested to. They should never  
appear in any logs, nor be stored in plaintext.   
  
The idea is to use this service as a central "repository" for apikeys. They can be added, listed and requested when needed.  
As some added security measure, every access to the stored apikeys should be logged in the form of an audit log, containing the time and requester of the apikey.  
  
There are four rpc methods to be implemented:  
- `AddApikey` can be used to add a new apikey to the system that is assigned to the specified userid  
- `ListApikeys` can be used to list all apikeys that are currently added for the given user  
- `GetApikey` can be used to get the plaintext details of the   
- `GetApikeyLogs` can be used to request a list of log entries  
  
For any kind of error that might happen during the rpc calls, refrain from using the builtin gRPC error handling but rather use the protobuf `errors` package to give errors more context and enable them to be shown and explained to the user.  
  
If there are any changes necessary to any of the already existing files, please do not refrain from changing or enhancing them in whatever way seems to be necessary to you.  
  
If you have any questions, please ask us. We're happy to help and discuss your ideas. 
  
### What we're expecting  
- An implementation of the ApikeyService server in Go, listening on port 50051  
- A Dockerfile to build the application and create a container  
- The service connects to an external SQL server of your choice for data storage  
- The credentials needed to connect to the SQL server can be passed to the service either using envvars or configuration files during runtime  
- The plain apikey & secret are stored in a secure way, this means not in plaintext. This could either by done in-service or by utilizing any external solution.  
- Some lightweight documentation how to start and interact with the Docker container  
  
### Extras  
There are a couple of bonus features that could be implemented but are not necessary to do so:  
- A caching layer using Redis, increasing the throughput and reducing the latency of some or all of the rpc methods. (Connecting to an external Redis server)  
- A Prometheus server that exports metrics about the service, such as throughput and call latency