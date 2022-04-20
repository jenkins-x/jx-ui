Experimental atm!

This is a full feature UI for JenkinsX 3.
The aim is to initially mimic everything that octant does.

# Set up

## Development

- Install a chrome plugin to disable CORS (not at all recommended for production)
- Build the go binary: `make backend`
- Run the binary: `./bin/ui`
- In another terminal, navigate to the `web` directory and start the front end server:
  ```
  cd web
  npm run start
  ```
- This should open a browser tab at localhost:3000

## Production

- Build production optimized frontend build: `make frontend`
- Build the go binary: `make backend`
- Run the binary: `./bin/ui`
- Navigate to localhost:8080 (8080 is the default port)

# Features

- Pipeline views - Step by step view of the logs
- Raw logs
- Start pipelines from the UI
- Responsive UI

# Nice to have

- Tests!
- Audit logging
- Full parity with jx cli
- Graph view of the pipeline
- SSO

# Credits

Uses https://github.com/daison12006013/sveltekit-starter as the base template:

# Contributions

Always welcomed!
Todo - Contribution guide
