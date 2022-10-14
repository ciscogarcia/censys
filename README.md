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
```make docker``` will build the docker images and start them running.

### Running test
```./test_service.sh``` will run a suite of tests that test insertion, deletion and overwriting data. It will make a simple api call to the test service that triggers the running of tests against the KV service.
