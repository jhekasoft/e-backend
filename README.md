# e-backend

```
░█▀▀░█▀▄░█▀█░█▀▀░█░█░█▀▀░█▀█░█▀▄
░█▀▀░█▀▄░█▀█░█░░░█▀▄░█▀▀░█░█░█░█
░▀▀▀░▀▀░░▀░▀░▀▀▀░▀░▀░▀▀▀░▀░▀░▀▀░
```

`e-backend` is a backend for all the projects.

![GitHub Release](https://img.shields.io/github/v/release/jhekasoft/e-backend?display_name=tag&logo=go&label=e-backend&color=%230279BA)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jhekasoft/e-backend)
![GitHub License](https://img.shields.io/github/license/jhekasoft/e-backend)

## Get started

### Installation

Install the latest version of `e-backend-cli`:

```bash
go install github.com/jhekasoft/e-backend-cli@latest
```

### Generate a new project

```bash
e-backend-cli app create my-ebackend-app
```

Then go to the project folder and copy the example of configuration file:

```bash
cd my-ebackend-app
cp config.example.yml config.yml
```

Then edit the `config.yml` file and try to run the project:

```bash
make run
```

You should see the following output:

```
░█▀▀░█▀▄░█▀█░█▀▀░█░█░█▀▀░█▀█░█▀▄
░█▀▀░█▀▄░█▀█░█░░░█▀▄░█▀▀░█░█░█░█
░▀▀▀░▀▀░░▀░▀░▀▀▀░▀░▀░▀▀▀░▀░▀░▀▀my-ebackend-app
Version: 1.0.0 | Build Time: 2026-07-02T03:30:29+03:00

Using config file: /home/user/src/my-ebackend-app/config.yaml
Run module Doc
Run module Health
⇨ http server started on [::]:1988
```

Check the health of the application:

```bash
curl http://localhost:1988/health
```

You should see the following output:

```json
{"OK":true,"Version":"1.0.0","BuildTime":"2026-07-02T03:35:57+03:00"}
```

### Generate a new module

Simple module:

```bash
e-backend-cli module create mySimpleModule
```

Run the project and check the module. In output you should see the following:

```bash
...
Run module Doc
Run module Health
Run module Mysimplemodule
⇨ http server started on [::]:1988
```

Check HTTP endpoint of the module:

```bash
curl http://localhost:1988/mySimpleModule
```

Output will be:

```json
{"result":"Hello from mySimpleModule module!"}
```

Create module with CRUD, that uses GORM and PostgreSQL:

```bash
e-backend-cli module create myCrudModule -t crud
```

And add PostgreSQL configuration to the `config.yml` file:

```yaml
DB:
  Enabled: true
  DSN: "host=localhost port=5432 user=user dbname=ebackend sslmode=disable"
```

Then run the project and open REST API documentation in the browser: [http://localhost:1988/doc/restapi/#tag--Mycrudmodule](http://localhost:1988/doc/restapi/#tag--Mycrudmodule). And you can check the CRUD endpoints of the module.

### Additional information about e-backend-cli
See [e-backend-cli](https://github.com/jhekasoft/e-backend-cli).

## Example of e-backend project

See [e-backend-boilerplate](https://github.com/jhekasoft/e-backend-boilerplate).

