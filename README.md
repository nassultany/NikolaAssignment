## Challenge 1

Challenge 1 reads and sorts an input json file. The file is assumed to be
an array of json objects containing id, first and last name, and address values.
They are sorted by last name and then by address in descending order, and output
to an output file. To run the program:

    go run challenge1.go -input test1.json -output test1out.json

## Challenge 2

Challenge 2 performs redirection from an old endpoint (/api/testing) to a new one
(/api/[module]), based on the value of module in the query. The new endpoint simply
returns the query parameters for verification. To run, do:

    go run challenge2.go -port 8080

and then open up your browser and test it out. For example, the url

    http://localhost:8080/api/testing?key=dGVzdAo&module=tapper&client=james

will redirect to

    http://localhost:8080/api/tapper?client=james&key=dGVzdAo&redirectFrom=%2Fapi%2Ftesting%3Fkey%3DdGVzdAo%26module%3Dtapper%26client%3Djames

You can also use curl:

    curl -L localhost:9098/api/testing -G -d 'key=dGVzdAo' -d 'client=jack' -d 'module=tapper'

to see headers:

    curl -L -v localhost:9098/api/testing -G -d 'key=dGVzdAo' -d 'client=jack' -d 'module=tapper'
