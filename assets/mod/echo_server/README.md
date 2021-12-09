# Echo Server

## Overview

It helps to start w/ a small-scale project when learning how sockets work at a low level. The echo server is a great way to introduce an apprentice to networking fundamentals

## Functional Requirements

A user should be able to interact w/ the echo server as follows:

* When a client sends a message to the server, the server sends a response back to the client containing the original message

* A client can send multiple messages to the server and get the echoed response back each time

* Multiple clients can send messages to server and get back their proper responses

## Implementation Requirements

* The server should establish a socket connection w/ the client using a low-level socket library. The goal of this exercise is to work w/ sockets directly

* The server should accept and return streams of data rather than raw strings

* The echo server should be covered by a robust suite of tests

## Dependencies

There are no dependencies for this project

## Prerequisites

It can be difficult to test around socket connections if an apprentice is not well-versed in good testing strategies. Thus, it is recommended that you build up a good knowledge of testing strategies (specifically, working w/ test doubles) before attempting this project

## Suggested Duration

It should take an apprentice one week to build the echo server

## Resources

## Evaluation

Here are a few example indicators that help you tell if an apprentice has successfully completed this project:

* Can the apprentice send a string to the server and get it back in a response?

* Can the apprentice gracefully start and stop the echo server w/o it throwing an error?
