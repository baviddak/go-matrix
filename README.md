# Go Matrix Adder

This is a sandbox project to experiment with sending HTTP requests from a python
script to a Go server. The flow of the CI pipeline is to send a pair of 
matrices as a JSON payload to the Go server and verify that the response is the 
total of all the ints in the matrices.