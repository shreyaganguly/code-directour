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
  -s string
    	Host name of the SMTP Server (default "smtp.gmail.com")
  -sendermail string
    	Sender email (default "no-reply@code-directour.com")
  -sendername string
    	Sender name (default "Code Directour")
  -t int
    	SMTP port (default 587)
  -token string
    	Slack Token for code-directour bot (If not passed sharing code snippets through slack is disabled)
  -u string
    	Username for SMTP authentication (If not passed sharing code snippets through email is disabled)
  -w string
    	Password for SMTP authentication (If not passed sharing code snippets through email is disabled)
```
## Getting started

Start the server and sign up with a user name and password and start sharing with the world!!

### Authentication

#### Enable sharing via mail
If you wish to use gmail as your smtp server (default) and your gmail ID is authorised with 2FA, make sure you create an [app password](https://myaccount.google.com/apppasswords) and pass it with `-w` flag.

#### Enable sharing via slack
If you wish to share snippets via slack make sure you create a [slack bot token](https://api.slack.com/custom-integrations/bot-users) and pass it with `-token` flag


## License

MIT, see the [LICENSE](https://raw.githubusercontent.com/shreyaganguly/code-directour/master/LICENSE.md) file.
