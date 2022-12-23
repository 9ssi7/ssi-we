# Docker Example

This example shows how to run SSI Web Events in a Docker container.

## Create .env file

```bash
cp .env.example .env
```

## Pull the image

```bash
docker pull ssibrahimbas/ssi-we
```

## Run the container

```bash
docker run -d -p 8080:8080 ssibrahimbas/ssi-we
```

## Open the application

Open your index.html file in your browser.
