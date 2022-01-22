# Unsplashed - BFF

### Deploying the App
Click this button to deploy the app to the DigitalOcean App Platform. If you are not logged in, you will be prompted to log in with your DigitalOcean account.

[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/Esequiel378/unsplashed-bff)

### Config

In order to run the project, you will need set an env variable inside of
a ```.env``` file on the root directory

```js
UNSPLASH_ACCESS_KEY=your-access-key
```

### Usage

run the project with

```bash
go run main.go
```

Or, if you prefer docker

```
docker-compose up
```

then you can visit

```python
http://localhost/api/unsplash/photos
http://localhost/api/unsplash/search/photos?query=london
```

