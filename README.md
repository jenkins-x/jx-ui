Experimental atm!

This is a full feature UI for JenkinsX 3.
The aim is to initially mimic everything that octant does.

# Set up

## Development

- Build the go binary: `make backend`
- Run the binary in one terminal: `./bin/ui`
- In another terminal, navigate to the `web` directory and start the front end server:

  ```bash
  cd web
  npm install
  npm run start
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
