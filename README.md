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
