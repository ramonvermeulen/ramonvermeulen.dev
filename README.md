# ramonvermeulen.dev
Repository containing the code for my personal website [ramonvermeulen.dev](https://ramonvermeulen.dev) which is 
currently still in early development. Since I am challenging myself this year to get more familiar with golang, I 
decided to build my personal website using golang.

## Environment Variables

| Environment Variable | Description                                                                                                         | Required                      | Default Value |
|----------------------|---------------------------------------------------------------------------------------------------------------------|-------------------------------|---------------|
| `ENV`                | Specifies the target environment. In production, the application uses a GCS bucket instead of the local filesystem. | No                            | `development` |
| `POSTS_BASE_PATH`    | Base path for posts.                                                                                                | No                            | `./posts`     |
| `GCS_POSTS_BUCKET`   | Name of the GCS bucket to use for posts. Required only if `ENV` is set to `production`.                             | Yes, if `ENV` == `production` |               |

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