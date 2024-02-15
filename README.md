# omdb
Simple command-line tool to display Movie details from the Open Movie Database

# Usage

## Set API Key...

'''bash
export OMDB_KEY=XXXXXXXXX
```

Get one from [omdbapi](https://www.omdbapi.com/apikey.aspx)

## Search for a movie

```bash
./omdb -title "Evil Dead 2"
./omdb -title "The Dark Knight"
./omdb -title "Dead Alive" -year 1992
```

## Output extra debug info

```bash
./omdb -title "Evil Dead 2" -debug
```
