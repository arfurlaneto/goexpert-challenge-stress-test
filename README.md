# What is this?

A stress tester written in Golang.

It is also one of the final challenges of https://goexpert.fullcycle.com.br/pos-goexpert/.

# How to use it?

Build the docker image:

```bash
docker build -t goexpert-stresser .
```

Run with docker:

```bash
docker run --rm goexpert-stresser --url=https://www.google.com/ --requests=20 --concurrency=5
```

## Arguments

|Value|Shorthand|Description|
|---|---|---|
|url|u|URL that will be stress tested.|
|requests|r|Total number of requests that will be made.|
|concurrency|c|Number of threads that will be making concurrent requests.|


## Windows Tip

If you are on Windows, you can use 172.17.0.1 to request to something running on your host.