# ramonvermeulen.dev

Personal website built with [Hugo](https://gohugo.io/).

The site uses a local Hugo theme under `themes/hugo-telegraph` and is deployed
as static website on a GCP Bucket.

## Layout

```text
ramonvermeulen.dev/
├── content/              # pages and blog section content
├── data/                 # structured site data such as about/experience
├── static/               # site-specific static files
├── themes/
│   └── hugo-telegraph/   # local Hugo theme
├── hugo.toml             # site configuration
├── package.json          # Tailwind dependencies used by Hugo Pipes
└── README.md
```

## Requirements

- Hugo `0.161+`
- Node.js `23+` with `npm`

## Local development

Install the Tailwind dependencies once:

```bash
npm install
```

Then run the site:

```bash
hugo server
```

