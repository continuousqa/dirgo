# DIRGO - Directory Buster
Similar to "dirb," Dirgo is a Golang project to simulate scanning a host for valid directories.

By supplying a dictionary, the scanner goes to the http(s)://[host] and appends /[word from dictionary], for each
desired status code (200 OK, 200, 400, 404, 500) that is returned, the path is printed out.

# Running
To run, execute and then input the host (including the http:// or https://)
Next, enter the status code you want to hunt for (i.e. 200)

# Future Goals
I'd love to get concurrency figured out for this one.  It would be great to supply a large dictionary and have this run
concurrent attempts.