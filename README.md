<div align="center">
<h1> upfluence-tt</h1>
<div style="background-color:blue; width:30%; margin:auto" >
  <img src="https://www.upfluence.com/wp-content/uploads/2024/02/upfluence-mono-white.svg" alt="Markdownify" width="100"/>
</div>
<br />

<a href="#what's-Upfluence-tt">What's upfluence-tt</a> • <a href="#how-to-use">How to use</a> • <a href="#going-further">Going further</a>
</div>

### What's Upfluence-tt
Upfluence-tt is an API use to analyze social network posts' engagement. It analyze a given data for a given amount of time 
#### Architecture
Project is develop using [hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)). This architecture enhance code readability, allow us to mock part of the code to make unit test easier and easily change part of the code. 

The API can analyze a lot of data, to prevent high usage of memory or use more memory than available nothing is stored in memory, only the results.

The code base is divided in three main package:
- Core:
  The core contains all the code related to the actual "work", it analyze data received and compute the expected result.
- Server:
  The server is the API of the project, it exposed one route "analyze", his job is to receive user request, pass the requiered data to the core, and return the analyzed data.
- Client:
  The client, is a "interface" to the upfluence stream. The client will listen to upfluence data stream fo a given amount of time (sent in the request) and send received data to the Core using channel.
    
### How to use
A Makefile is available a the root of the repository, multiple command are available:
- build (build the program)
- run (build and run the program)
- test (run unit tests)
- run-docker ((and run-docker-linux not tested) run the program on a docker container)
- build-docker ((and build-docker-linux not tested) build the docker container)

If you want to run the program manually and localy, simply run thoose commands:
```bash
$ go build ./cmd/upfluencett
$ ./upfluencett -upfluence-url="<url of upfluence stream>" -server-port="<the wanted port for the server>"
```
Or
```bash
$ go run ./...  -upfluence-url="<url of upfluence stream>" -server-port="<the wanted port for the server>"
```

Only one request is available: "/analysis"
The query take 2 <red>mandatory</red> parameters:
1. Dimension
    The data to be analyzed. Multiple dimensions are available:
    - likes
    - comments
    - favorites
    - retweets
2. Duration
    The duration of the analyzis (5s, 10m, 24h are all valid input)

If your API is running localy you can run for exemple:
`localhost:8080/analysis?duration=5s&dimension=likes`
The API while analyze "likes" for 5 seconds

Here is an exemple of a response
```json
{
  "total_posts": 20,
  "minimum_timestamp": 1440494430,
  "maximum_timestamp": 1638774882,
  "avg_likes": 50
}
```

### Going further
What should happend if the input stream is stoped during analyzis. Currently I did not handle any scenario, but we could imagine implementing a retry system or an early response.

I've devided to use pointer of int to serialize data from the stream, because in Go, int default value is 0, which mean even if the serialized value is not in the input, its value will be 0. The way I did my result calculation would have included thoose value in the result. 
Instead of using int pointer, we could imagine implementing a decision maxtrix depending of the dimension needed, which mean we could define a data structure to serialize input, and use the decision matrix

More testing need to be done, currently only the serializing and computing are tested, API response data and timing is not tested

#### Trade off
I've decided to go for a "generic" way of serializing upfluence's data stream, using map[string]interface{} instead of defined type, this choice make the code a bit less clear but make it easier to add new input data type.

The server does not store input data, everything is computed directly after receiving a message, which mean the memory consumption of the program is really little. If the user need to gather data for multiple weeks or months, storing every message means that the server would need a lot of memory, by not doing that and computing data every time we receive a message, only the strcutur containing the result is persistant over the course of the program.
