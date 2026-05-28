<div align="center">
  <img alt="logo" data-is-relative="true" src="static/android-chrome-512x512.png" width="200" height="200"/>
</div>

<h1 align="center">ramonvermeulen.dev</h1>

<p align="center">
  <a href="https://github.com/ramonvermeulen/ramonvermeulen.dev/actions/workflows/cd_hugo.yml">
    <img alt="CD" src="https://github.com/ramonvermeulen/ramonvermeulen.dev/actions/workflows/cd_hugo.yml/badge.svg">
  </a>
</p>

Repository containing the source code for my personal website [**ramonvermeulen.dev**](https://ramonvermeulen.dev).
The site is built with [**Hugo**](https://gohugo.io/) and [**Tailwind CSS**](https://tailwindcss.com).

## Prerequisites

- [**Hugo Extended 0.161+**](https://gohugo.io/installation/)
- [**Node.js 22+**](https://nodejs.org/en/download/)
- [**npm**](https://nodejs.org/en/download/)

## Directory Structure

```text
ramonvermeulen.dev
├── content/              # pages and blog content
├── data/                 # structured site data
├── static/               # site assets served as-is
├── themes/hugo-telegraph # local Hugo theme
├── hugo.toml             # site configuration
└── README.md
```

## Local Development

```bash
npm install
hugo server
```
