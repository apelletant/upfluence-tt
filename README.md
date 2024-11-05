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
     - Server
        The server is the API of the project, it exposed one route "analyze", his job is to receive user request, pass the requiered data to the core, and return the analyzed data.
     - Client
        The client, is a "interface" to the upfluence stream. The client will listen to upfluence data stream fo a given amount of time (sent in the request) and send received data to the Core using channel.
    
### How to use
A Makefile is available a the root of the repository, multiple command are available:
    - build (build the program)
    - run (build and run the program)
    - test (run unit tests)
    - run-docker (and run-docker-linux, run the program on a docker container)
    - build-docker ((and build-docker-linux) build the docker container)

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
  "total_posts": 20
  "minimum_timestamp": X,
  "maximum_timestamp": Y,
  "avg_likes": 50
}
```

### Going further


<style>
    red {
        color: red;
    }
</style>