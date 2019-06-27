# Retarded Sponge Bot

<img src=https://upload.wikimedia.org/wikipedia/commons/thumb/8/82/Telegram_logo.svg/800px-Telegram_logo.svg.png width=100/>

Simple telegram bot that uses my other project [retarded bob generator](https://github.com/nyc4m/retarded-bob-generator).



## Build

```bash
$ go get -u //for dependencies
$ go build
```

## Run
To make it run you need to have an api token from telegram, and store it in a environment variable such as 

```bash
$ TOKEN="YOURTOKENHERE"
```

## Commands

| Command         | Description            |
|:---------------|:----------------------|
| `/retarded <sentence>` | send a retarded string |
| `/retardedPic <sentence>` | send a retarded picture (pic + sentence) |
