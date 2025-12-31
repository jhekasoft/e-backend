# e-backend

`e-backend` is a backend for all the projects.

![cat](./modules/doc/data/public/android-chrome-192x192.png)

```
▗▄▄▄▖▗▄▄▖  ▗▄▖  ▗▄▄▖▗▖ ▗▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄ 
▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌▗▞▘▐▌   ▐▛▚▖▐▌▐▌  █
▐▛▀▀▘▐▛▀▚▖▐▛▀▜▌▐▌   ▐▛▚▖ ▐▛▀▀▘▐▌ ▝▜▌▐▌  █
▐▙▄▄▖▐▙▄▞▘▐▌ ▐▌▝▚▄▄▖▐▌ ▐▌▐▙▄▄▖▐▌  ▐▌▐▙▄▄▀
```

## Create database

```bash
sudo -iu postgres
createdb ebackend
```

## Prepare

```bash
cp .e-backend.example .e-backend
```

And then edit `.e-backend` file.

## Run HTTP-server

```bash
make run
```

## Building

Build binary:

```bash
make build
```

Clean:

```bash
make clean
```

Run binary:

```bash
./build/e-backend serve
```

## Run as service (POSIX systems with systemd)

```bash
sudo mkdir /opt/e-backend
sudo cp ./build/* /opt/e-backend -r
sudo cp /opt/e-backend/.e-backend.example /opt/e-backend/.e-backend
sudo cp ./systemd/e-backend.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now e-backend.service
```

## Module generation

```bash
go run -tags="all dev" main.go module create [name] -t crud
```

Where `name` is name of module is `lowerCamelCase`, `-t` is template name
(simple, crud).
