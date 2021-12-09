# Basic HTTP Server

## Overview

In this feature, an apprentice will build an HTTP server which includes routes, requests, and responses. The routes must be customizable with a URL, a verb, and an action to take when the route is called. This work will form the basis of many other features.

This feature corresponds to the acceptance tests in `01_getting_started` in the [HTTP Server Spec](../../../http_server_spec/README.md).

## Functional Requirements

A user should be able to interact with the HTTP server as follows:

* When a client sends a properly formatted request to the server, the server sends an appropriate response back to the client.

* A client can send different HTTP requests to the server and get the appropriate response back each time.

* Different clients can send messages to server and get back their proper responses.

* The server should be able to handle 200, 300, and 400-level responses. Not every response code needs to be complete, but there should be a few representative response codes implemented for each level.

## Implementation Requirements

* The server should establish a socket connection with the client using a low-level socket library. The goal of this exercise is to work with sockets directly.

* The server should accept and return streams of data rather than raw strings.

* Although not strictly speaking necessary, the HTTP server is a good time to introduce statically typed languages like Java, C#, or Swift.

* The HTTP server should be covered by a robust suite of unit tests.

* The HTTP server should pass all of the tests covered in `01_getting_started` in the [HTTP Server Spec](../../../http_server_spec/README.md).

## Dependencies

Although the echo server is not a dependency of the HTTP server, many apprentices find it useful to implement this smaller project before attempting the more complicated HTTP server.

## Suggested Duration

It should take an apprentice two weeks to implement this feature.

## Resources

* [WizardZines: HTTP](https://wizardzines.com/zines/http/)

* [Boundaries: Gary Bernhardt](https://www.destroyallsoftware.com/talks/boundaries)

* [HTTP in a Nutshell](https://medium.freecodecamp.org/restful-services-part-i-http-in-a-nutshell-aab3bfedd131)

* [HTTP Made Really Easy](https://www.jmarshall.com/easy/http/)

* [HTTP Server Spec](https://github.com/8thlight/http_server_spec)

* [What is a Socket?](https://docs.oracle.com/javase/tutorial/networking/sockets/definition.html)

## Evaluation

Here are a few example indicators that help you tell if an apprentice has successfully completed this feature:

* Can the apprentice send a properly formatted HTTP request to the server via the browser and get back the appropriate response?

* Can the apprentice send a properly formatted HTTP request to the server via cURL and get back the appropriate response?

* Can the apprentice send a properly formatted HTTP request to the server via a GUI app like Postman or Paw and get back the appropriate response?
