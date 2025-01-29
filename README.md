<div align="center">
    <img alt="logo" data-is-relative="true" src="static/icon/android-chrome-512x512.png" width="200" height="200"/>
    <h1>ramonvermeulen.dev</h1>
    <a href="https://go.dev/doc/devel/release"><img alt="Go version" src="https://img.shields.io/github/go-mod/go-version/ramonvermeulen/ramonvermeulen.dev"></a>
    <a href="https://github.com/ramonvermeulen/ramonvermeulen.dev"><img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/ramonvermeulen/ramonvermeulen.dev"></a>
</div>

Repository containing the code for my personal website [ramonvermeulen.dev](https://ramonvermeulen.dev) which is 
currently still in early development. Since I am challenging myself this year to get more familiar with golang, I 
decided to build my personal website using golang.

## Directory Structure

```
ramonvermeulen.dev
│   main.go         // entry point of the application
│
└───internal        // main source code (go modules)
│
└───assets          // original css and js assets (not minified)
│
└───public          // public website assets such as minified css, js, images, icons and blogs
│                   // go file server used for dev, GCS bucket + cloudflare CDN for prod
│                   // publish new blog posts without redeployment of the application
│
└───templates       // HTML templates (go templates)
```

## Environment Variables

| Environment Variable | Description                                                                                                         | Required                      | Default Value    |
|----------------------|---------------------------------------------------------------------------------------------------------------------|-------------------------------|------------------|
| `ENV`                | Specifies the target environment. In production, the application uses a GCS bucket instead of the local filesystem. | No                            | `development`    |
| `POSTS_BASE_PATH`    | Base path for posts.                                                                                                | No                            | `./static/posts` |
| `GCS_POSTS_BUCKET`   | Name of the GCS bucket to use for posts. Required only if `ENV` is set to `production`.                             | Yes, if `ENV` == `production` |                  |

## Local Development

### Install dependencies
```bash
go mod tidy
npm install
```

### Run server
```bash
go run main.go
```

### Run watchers
```bash
npm run dev
```