#!/bin/bash

# Environment variables
OMDB_APIKEY=$1
TMDB_APIKEY=$2

# Functions are
function start_api() {
    export OMDB_APIKEY=$OMDB_APIKEY
    export TMDB_APIKEY=$TMDB_APIKEY
    if [[ -z $OMDB_APIKEY || -z $TMDB_APIKEY ]] 
    then
        echo "One of the API keys is missing"
        echo "OMDB apikey: ${OMDB_APIKEY}"
        echo "TMDB apikey: ${TMDB_APIKEY}"
    else
        # exec ../output/go-movie-api
        ../cmd/go-movie-api
    fi
}

# Main run area
start_api