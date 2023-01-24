# CVEdash Backend API

The backend REST API server is written in Go.

External dependencies:
- [Chi](https://github.com/go-chi/chi) for the API server
- [GORM](https://github.com/go-gorm/gorm) with the [GORM SQLite Driver](https://github.com/go-gorm/sqlite) for database management

## Config

**Copy or rename `default-config.json` to `config.json` and edit it to your needs.** Please be aware of [NVD's API best practices](https://nvd.nist.gov/developers/start-here), especially concerning the time interval between updates (`minutesBetweenNVDUpdates`)**

## Usage
To start the API server, simply run the following command inside the `api` (this) directory.
```cmd
go run cmd/cvedash/main.go
```

On the first startup, the server immediately starts pulling all CVE data from NVD's API and stores it in a new SQLite database file named according to the config.

⚠️ The dataset currently contains over `200,000` entries, which is approximately 1 GB of JSON data being downloaded. Due to API rate limiting, this process will take a while, so get a cup of coffee in the meantime (each request returns up to 2000 entries, so there will be a little over 100 requests with a sleep of 6 seconds (1 second if you have an API key) between each request).

Since much of this data is irrelevant and removed, the database file will only be around 100 MB in size.

After all data is imported, the API server starts at http://localhost:8077/. While running, the database will be updated with new changes from NVD's API every once in a while.

## Disclaimer
This product uses the NVD API but is not endorsed or certified by the NVD.
