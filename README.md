# iseeyou

Simple experiment with a client server approach.

Basic idea is:

### Client

* The client written in Python is running on a Raspberry PI
* It should handle a 3G/4G modem connection to work on a remote place with no WiFi
* It should then send data read by a PIR (movement information)

### Server

* Should expose a REST endpoint
* Receive the data from a POST request and store the information into an SQL database
