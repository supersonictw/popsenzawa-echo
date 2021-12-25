# Step to step guide on setting up go server

Origin: <https://git.io/Jy3gM>

Not sure if this is the best format for submitting a minimal setup guide, if @supersonictw thinks it's not appropriate then I will close it anyway

### Required docker images:

Popcat server require redis for cache and mysql for stats aggregation so you will these image other than the popcat image.

```
docker pull supersonictw/popcat-echo
docker pull redis
docker pull mysql:5.7
```

```
docker run -p 6379:6379  -d redis:latest
docker run -p 3306:3306 -e  MYSQL_ROOT_PASSWORD=<Your password> -d mysql:5.7
```

## Setup MYSQL

Use the initialize.sql script to setup your MYSQL database schema. You may need to connect to your database before this.

## Execute image

Setup docker environment variable file according to [.env.sample](https://github.com/supersonictw/popcat-echo/blob/main/.env.sample)

I find this DSN format working under my own setup ( under Linux environment you will need to change docker.for.mac.localhost to the docker IP )

```
MYSQL_DSN=<mysql user>:<password>@tcp(docker.for.mac.localhost:3306)/<your database name>?charset=utf8
```

Execute popcat server

```
docker run -p 8013:8013 --env-file <your environment variable file> supersonictw/popcat-echo
```

You should be able to access popcat server at localhost:8013 !
