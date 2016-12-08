# Terragen
A service for exploring, generating, experimenting, and tweaking generated terrain.

## Building - Service

## Building - Web
The web server is built separately from the backend service. The backend service will serve up a simple index file that loads the front end.

### Getting started
```
> brew install yarn
```

### Updating Dependencies
```
> yarn install
```

### Running local dev web server
```
> yarn run dev
# Serves up the frontend on a dev index file, address displayed in stdout
```

### Build production artifact
```
> yarn run build
# Outputs artifact to build/static - this is where the backends index file will point to
```

### Build and Deploy production artifact
Currently, ssh into the production server and run
```
> git pull
> yarn install
> yarn run build
> make deploy-web-local
```
The running service will then use the resulting output
