# go-web-scraper

A modular and extensible HTML parsing component written in Go, following clean architecture principles. This parser ingests HTML files (e.g., from pre-rendered web pages) and extracts structured job listings or other data.

It is designed to be both locally testable and cloud-deployable (e.g., on AWS Lambda), with the ability to fetch, parse, and store job listings or other structured content from web sources.

This component pairs with [node-web-fetcher](https://github.com/PhilNel/node-web-fetcher).

📚 **Documentation:** [https://philnel.github.io/docs-web-scraper](https://philnel.github.io/docs-web-scraper)

## 🛠 Installation

Make sure you have Go installed (version 1.21+ recommended).

This project vendors the dependencies so no explicit installation is required.

## 📦 Requirements
This project depends on a Node.js-based renderer to fetch the HTML for parsing.

Make sure the `node-web-fetcher project` is available and has run to produce rendered.html.

## 🧪 Usage
Run the full flow:

```make
make parse
```

## 🔧 Dependencies
Runtime:
- PuerkitoBio/goquery — jQuery-style HTML parsing
- aws/aws-sdk-go-v2 — AWS components
- sirupsen/logrus — Structured logging
- jessevdk/go-flags — Loads config from environment variables

Dev only:
- sirupsen/godotenv — Loads local environment variables
- golangci-lint — For linting the project source code
- stretchr/testify — For unit testing