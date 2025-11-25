<div align="center">
    <img alt="logo" data-is-relative="true" src="public/icon/android-chrome-512x512.png" width="200" height="200"/>
    <h1>ramonvermeulen.dev</h1>
    <a href="https://go.dev/doc/devel/release"><img alt="Go version" src="https://img.shields.io/github/go-mod/go-version/ramonvermeulen/ramonvermeulen.dev"></a>
    <a href="https://github.com/ramonvermeulen/ramonvermeulen.dev"><img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/ramonvermeulen/ramonvermeulen.dev"></a>
    <a href="https://github.com/ramonvermeulen/ramonvermeulen.dev/actions/workflows/cd_docker.yml"><img alt="CD" src="https://github.com/ramonvermeulen/ramonvermeulen.dev/actions/workflows/cd_docker.yml/badge.svg"></a>
</div>

Repository containing the source code for my personal website [**ramonvermeulen.dev**](https://ramonvermeulen.dev).
The website is built with Go and uses simple HTML templates in combination with [**TailwindCSS**](https://tailwindcss.com) for styling.
The blog is rendered using [**Goldmark**](https://github.com/yuin/goldmark) to convert blogs written in markdown files to HTML.

## Prerequisites
- [**Go 1.25+**](https://go.dev/dl/)
- [**npm**](https://nodejs.org/en/download/)
- [**Make**](https://www.gnu.org/software/make/)
- (Optional) [**Docker**](https://www.docker.com/get-started) - only needed if you want to build the production Docker image
- (Recommended) [**direnv**](https://direnv.net/) - for loading environment variables from `.env` file
- (Recommended) [**pre-commit**](https://pre-commit.com/) - for running git hooks

## Directory Structure

```
ramonvermeulen.dev
│   main.go         // entry point of the application
│
└───internal        // main source code
│
└───assets          // original css and js assets (not minified)
│
└───public          // public website assets such as minified css, js, images, icons and blogs
│                   // go file server used for dev, GCS bucket + cloudflare CDN for prod
│                   // publish new blog posts without redeployment of the application
│
└───templates       // HTML templates (go templates)
```

## Local Development

### Install dependencies
```bash
go mod tidy
npm install
```

### Run server & watchers
```bash
make dev
```

## Deployment
The website is deployed on [**Cloud Run**](https://docs.cloud.google.com/run/docs), and uses [**Google Cloud Storage**](https://docs.cloud.google.com/storage/docs) to serve static assets.
Everything is deployed using [**GitHub Actions**](https://docs.github.com/en/actions).

