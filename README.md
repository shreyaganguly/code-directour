# code-directour

Ever felt the need to share a piece of useful code with your colleague? Or maybe share a part of a function of a big application your friend is writing? Sharing code seems a bit of tedious task with all its syntax highlighting and lack of proper code snippet sharing platforms.

**code-directour** is a web-application meant for easy maintenance and sharing of code snippets in various languages and across multiple platforms, server side is implemented in golang and client side implemetation uses jQuery, bootstrap. It helps you share code snippets with other code-directour users or any other users via link, mail or slack. Go ahead and start sharing!

## Installation

Assuming you have installed a recent version of
[Go](https://golang.org/doc/install), you can simply run

```
go get -u github.com/shreyaganguly/code-directour
```

This will download `code-directour` to `$GOPATH/src/github.com/shreyaganguly/code-directour`. From
  this directory run `go build` to create the `code-directour` binary.

## Usage

Start the server by executing `code-directour` binary. By default, server will listen to http://0.0.0.0:8080 for incoming requests.

```
Usage of ./code-directour:
  -b string
    	Host to start your code-directeur (default "0.0.0.0")
  -db string
    	File to store the db (default "directour.db")
  -p int
    	Port to start your code-directour (default 8080)
```
## Getting started

Start the server and sign up with a user name and password and start sharing with the world!!

### Profile Settings

#### Enable sharing via link
If you wish to share the code snippets via link. You **must** provide the `Link Endpoint` to see `Share By Link` option for your snippets

#### Enable sharing via mail
You **must** provide `SMTP Server`, `SMTP Port`, `Email Address`, `Email Password`, `Sender Email` to see `Share By Mail` option for your snippets. If you wish to use gmail as your smtp server (default) and your gmail ID is authorised with 2FA, make sure you create an [app password](https://myaccount.google.com/apppasswords).

#### Enable sharing via slack
You **must** provide `Slack Token` to see `Share By Slack` option for your snippets. Make sure you create a [slack bot token](https://api.slack.com/custom-integrations/bot-users).

## License

MIT, see the [LICENSE](https://raw.githubusercontent.com/shreyaganguly/code-directour/master/LICENSE.md) file.
