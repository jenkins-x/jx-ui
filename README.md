# README

**Experimental atm!**

This is a full feature UI for Jenkins X 3.
The aim is to initially mimic everything that octant does.

## Set up

## Development

- Build the go binary: 
    - Linux: `make backend`
    - Darwin: `make backend-darwin`
- Run the binary in one terminal: 
    - Linux: `./build/linux/jx-ui` 
    - Darwin: `./build/darwin/jx-ui` 
- In another terminal, navigate to the `web` directory and start the front end server:

  ```bash
  cd web
  npm install
  npm run dev
  ```

- This should open a browser tab at localhost:3000

### Testing

#### Backend

TBA

#### Frontend

We use the following libraries for testing

- To test simple typescript functions, we use vitest
- To test rendering components, we use testing-library/svelte
- To run End-to-End tests we use Playwright.

To run vitest, execute

```bash
npm run test
```

To get a coverage report, run

```bash
npm run coverage
```

To perform E2E (End-To-End) tests, run

```bash
npm run playwright
```

## Production

- Build backend and frontend: `make all`
- Run the binary: `./build/linux/jx-ui`
- Navigate to localhost:9200 (9200 is the default port)

## Docker build locally

- To build the docker image, run  `make build.docker.local`
- Run `docker run --net host -v $HOME/.kube:/root/.kube --rm --name jx-ui jenkins-x/jx-ui:latest`

## Features

- Pipeline views - Step by step view of the logs
- Raw logs
- Start pipelines from the UI
- Responsive UI

## Nice to have

- Tests!
- Audit logging
- Full parity with jx cli
- Graph view of the pipeline
- SSO

## Credits

Uses https://github.com/daison12006013/sveltekit-starter as the base template:

## Contributions

Always welcomed!
Todo - Contribution guide
