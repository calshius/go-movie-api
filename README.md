# Movie API

This api retrieves data frm OMDB and TMDB and finds wether a movie has made a profit or a loss and calculates a percentage based on the budget and return revenue.

[![Build Status](https://travis-ci.org/calshius/go-movie-api.svg?branch=master)](https://travis-ci.org/calshius/go-movie-api)

##  Using the API

To start this api you need an OMDB api key and a TMDB api key.

Here's how you get the relevant keys:

[OMDB](http://www.omdbapi.com/)

[TMDB](https://developers.themoviedb.org/3/movies/get-movie-details)

Once you have retrieved the API keys you can then start the app by cloning down this repo and moving into the scripts folder and running:

```bash
./api-run.sh <OMDB_APIKEY> <TMDB_APIKEY>
```

and this will start the application:

```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET   /movie                    --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

Now that the api has started you can query it like so:

Profitable movie:

```bash
curl -s -X GET http://localhost:8080/movie?name=avengers | python -m json.tool
{
    "Percentage": 690,
    "Profit": 1299557910,
    "budget": 220000000,
    "original_title": "The Avengers",
    "release_date": "2012-04-25",
    "revenue": 1519557910
}

```

Lossly movie:

```bash
curl -s -X GET http://localhost:8080/movie?name=the+room | python -m json.tool
{
    "Loses": 5998200,
    "Percentage": 333333,
    "budget": 6000000,
    "original_title": "The Room",
    "release_date": "2003-06-27",
    "revenue": 1800
}

```