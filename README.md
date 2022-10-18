# Problem
- From the email provided by censys
> Build a simple, maintainable key-value store by implementing two services that communicate with each other: a key-value store (KV service) and a client that tests the key-value store (test client). This project should ideally take 2 hours to complete and no more than 4 hours.
>
>The KV service should implement a basic JSON Rest API as the primary public interface. The interface must support the following operations:
>
>    Store a value at a given key.
>    Attempt to retrieve the value for a given key.
>    Delete a key.
>
>
>The test client should be a service that has at least two operations that exercise the correctness of the KV service interface. All endpoints will return whether or not they succeeded. The endpoints are:
>
>    test_deletion: This should instill confidence in the delete operation.
>    test_overwrite: This should instill confidence that the most recently set value for a key is the one that is served.
>
>
>Additional requirements:
>
>    The services should be executed independently within Docker containers.
>    It should be easy to test the KV store and the test client from the command line.
>    Please choose one of these programming languages for implementation: golang, python, java, scala.

## Running services
```make docker``` will build the docker image and start it running.

## Running tests
```curl http://localhost:10000/test``` will run a suite of tests that test insertion, deletion and overwriting data. It will make a simple api call to the test service that triggers the running of tests against the KV service.

The responses will be output to STDOUT, 1 JSON message per test as defined in test_service.go

Manual tests can also be run with curl
``` curl -X POST http://localhost:10000/kv -d '{"key": "name", "value": "cisco"}'```
``` curl -X PUT http://localhost:10000/kv -d '{"key": "name", "value": "john"}'```
``` curl -X GET 'http://localhost:10000/kv?key=name'```
``` curl -X DELETE http://localhost:10000/kv -d '{"key": "name"}'```

## Notes
There is a (purposeful) discrepancy while handling responses from the KV service. In the case of an error, we are returning a hand-rolled JSON payload. This is due to the fact that one of the failure cases is Marshaling our JSON responses, and that would taint the error returned to the client. In order to account for this and be consitant across all errors, I have made the decision to hand-roll a simple json error in case of errors.
