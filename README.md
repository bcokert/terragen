# Terragen
A service for exploring, generating, experimenting, and tweaking generated terrain.

## Releasing

### Release a New Version
Versioning is done via git tags, as well as the version in package.json. Both are updated by the same mechanism.
```
> git add ... # all changes to be in the version
> git commit -m ...
> yarn version # and enter version number (semver)
> git push # pushes both the code and the version tag
```

### Deploy to Production
The front end and backend are deployed in the same step.
The backend artifact is AUTOMATICALLY rebuilt before a deploy. It's the same as the dev artifact, except built for a different architecture.
The frontend artifact is MANUALLY built. If the front end hasn't changed, then there's no reason to rebuild it, as it takes the longest.

To build the front end artifact:
```
> yarn run build
```

To deploy both artifacts (and build the service artifact):
```
> ./bin/deploy.sh path/to/pem/file hostaddress
```

## Developing
The service is written in go, and provides a Rest api.
The index (root) action is special, and is used exclusively to serve the website - it serves a tiny html file that loads the separate website artifact.
The website is otherwise completely separate from the backend.

### Getting Started
```
> brew install go # Follow the instructions to setup go
> brew install yarn # May need to install node first
```

### Updating Dependencies
```
> go get
> yarn install
```

### Running Local Dev Server
```
> make run
```
