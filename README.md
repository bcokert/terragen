# Terragen
A service for exploring, generating, experimenting, and tweaking generated terrain.

## Building - Service
The service is written in go, and provides a Rest api for all actions.
The index action is special, and is used exclusively to serve the website.

### Getting started
```
> brew install go
# Follow the instructions to setup go
```

### Updating Dependencies
```
> go get
```

### Running local dev server
```
> make run
# Serves up the backend on a local service - this service also served a website artifact, if one is available.
# However, you should use the web dev server if developing the website
```

### Build Production Artifact
The production artifact is built automatically when you run the deploy to production process.
The output is exaclty the same as the dev server, except it's compiled for a different architecture.

### Deploy to Production
```
> ./bin/deploy.sh path/to/pem/file hostaddress
```

The deploy will build a target specific version of the service, but depends on a previously built web artifact.
If it doesn't find it, it will prompt you accordingly.

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

### Build Production Artifact
```
> yarn run build
# Outputs artifact to build/static - this is where the backends index file will point to
```

### Deploy to Production
```
> ./bin/deploy.sh path/to/pem/file hostaddress
```

The deploy will build a target specific version of the service, but depends on a previously built web artifact.
If it doesn't find it, it will prompt you accordingly.
