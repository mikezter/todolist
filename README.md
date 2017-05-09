# Todolist

[![](https://goreportcard.com/badge/github.com/gammons/todolist)](https://goreportcard.com/report/github.com/gammons/todolist)
[![Build Status](https://travis-ci.org/gammons/todolist.svg?branch=master)](https://travis-ci.org/gammons/todolist)
[![Coverage Status](https://coveralls.io/repos/github/mikezter/todolist/badge.svg?branch=edit-todos-tests)](https://coveralls.io/github/mikezter/todolist?branch=edit-todos-tests)

Todolist is a simple and very fast task manager for the command line.  It is based on the [Getting Things Done][gtd] methodology.

[gtd]: http://lifehacker.com/productivity-101-a-primer-to-the-getting-things-done-1551880955

## Documentation

See [The main Todolist website][tdl] for the current documentation.

[tdl]: http://todolist.site

Quick start:

		todo a[dd] remember the milk due monday

		todo l[ist]

		todo c[omplete] 14

		todo ar[chive] 14

## Is it good?

Yes.  Yes it is.

## Web interface

Quick Start using [Docker](https://github.com/docker/docker.git):

Build the docker image:

		$ git clone https://github.com/gammons/todolist.git
		$ cd todolist
		$ docker build -t todolist .

If you don't have an existing todo-list file, yet, create one in your
home directory:

		$ docker run -v ~/.todos.json:/.todos.json todolist init

Start the webserver and expose port 7890 on localhost:

		$ docker run -d -v ~/.todos.json:/.todos.json -p 127.0.0.1:7890:7890 todolist

Happy GTD!

## Author

Please send complaints, complements, rants, etc to [Grant Ammons][ga]

## License

Todolist is open source, and uses the [MIT license](https://github.com/gammons/todolist/blob/master/LICENSE.md).

[ga]: https://twitter.com/gammons
