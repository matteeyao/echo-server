# Structured Data

## Overview

In this feature, an apprentice will expand upon an existing HTTP server to support a variety of data types. Their server will be able to return structured data in formats such as HTML, JSON, and XML, each with the appropriate MIME type.

This feature corresponds to the acceptance tests in `02_structured_data` in the [HTTP Server Spec](../../../http_server_spec/README.md).

## Functional Requirements

The HTTP server should have the following behavior:

* It should be able to return structured data in a variety of formats including HTML, JSON, and XML.

* The HTTP response should have the appropriate content type headers to specify what type of data is being returned.

## Implementation Requirements

* The apprentice must leverage an existing HTTP server (e.g. use the routes, requests, and responses that already exist).

* The HTTP server should be covered by a robust suite of unit tests.

* The HTTP server should pass all of the tests covered in `02_structured_data` in the [HTTP Server Spec](../../../http_server_spec/README.md).

## Dependencies

The [Basic HTTP Server](https://github.com/8thlight/apprenticeship_syllabus/blob/48437f37ecfce041928afebc004d859b7a992911/shared_resources/projects/http_server/01_beginner/basic_http_server.md) is a dependency of this feature.

## Prerequisites

It can be difficult to test around socket connections if an apprentice is not well-versed in good testing strategies. Thus, it is recommended that you build up a good knowledge of testing strategies (specifically, working with test doubles) before attempting this project.

## Suggested Duration

It should take an apprentice one week to implement this feature.

## Resources

* [HTML Tutorial](https://www.w3schools.com/html/)

* [Introducing JSON](https://www.json.org/json-en.html)

* [MIME Types (IANA Media Types)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)

* [Schema.org](https://schema.org/)

* [XML Tutorial](https://www.w3schools.com/xml/)

## Evaluation

Here are a few example indicators that help you tell if an apprentice has successfully completed this feature:

* Can the apprentice send a properly formatted HTTP request to the server via the browser and get back the appropriate response?

* Can the apprentice send a properly formatted HTTP request to the server via cURL and get back the appropriate response?

* Can the apprentice send a properly formatted HTTP request to the server via a GUI app like Postman or Paw and get back the appropriate response?
