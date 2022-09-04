# Telegram pocketbook bot
## Description

This is a telegram bot that can save and keep you notes

## Purpose of project

The goals of development of this bot are:
- upgrade golang skills - working with filesystem, http, service architecture etc;
- host and run own applications on a cloud services (in this case - on _Digital Ocean_);
- learn CI (continuous integration) and autodeployment using GitHub Actions.

## Deployment

### Deploying on a server
To run this bot on server, you need to set up:
- Digital Ocean Droplet (server);
- Digital Ocean Container registry;

Next, you will need to fork repository and clone it:
```
git clone https://github.com/slbmax/telegram-pocketbook-bot
cd telegram-pocketbook-bot
```

Then you should specify GitHub secrets:
- `TELEGRAM_TOKEN` - token for telegram bot;
- `DIGITALOCEAN_ACCESS_TOKEN` - Digital Ocean access token;
- `HOST` - address of Digital Ocean droplet;
- `USERNAME` - username of host (_root_ by default)
- `PASSWORD` - password to connect to droplet server

To perform autodeployment, you can only push some changes and GitHub Actions will do it automatically.

### Local deployment

Clone repository:
```
git clone https://github.com/slbmax/telegram-pocketbook-bot
cd telegram-pocketbook-bot
```

If you have Go 1.18 installed, just run:
```
export TOKEN='your_telegram_bot'
go run main.go
```

If not, you can run docker container:
```
docker build .
docker run -e TOKEN=your_telegram_token docker_container_id
```

