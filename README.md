# Unsplashed - BFF

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

