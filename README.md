# Gauge ⚖️
Measure the weight of languages in your GitHub profile. This tiny tool gets the amount of code written on every language in your public repos and compiles them as a table. 

## Usage 

To run it with go from the project's directory, simply do `go run . <github username>`, after building it just do

```
gauge <github username>
```

You can set an environment variable `GITHUB_TOKEN` (or in a .env file) with a [personal token](https://github.com/settings/tokens) to access private repositories.
