wat - Wolfram Alpha command line Tool
===

A simple command-line tool to query Wolfram Alpha's API

Installation
===
A simple go get will suffice: ```go get github.com/seen/wat```

Configuration
===
Wolfram Alpha requires an APP ID for querying it's API. You can
get one here: https://developer.wolframalpha.com/

`wat` loads the APP ID from the environment variable `WOLFRAM_API_ID`.
Set it with your APP ID from Wolfram Alpha and you're good to go!
