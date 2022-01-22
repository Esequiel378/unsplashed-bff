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

then you can visit

```python
http://localhost:4000/api/unsplash/photos
http://localhost:4000/api/unsplash/search/photos?query=london
```
