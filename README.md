# Unsplashed - BFF

### Config

In order to run the project, you will need set a few env variables inside of
a ```.env``` file on the root directory

```python
UNSPLASH_ACCESS_KEY=your-access-key
# Coma separated urls (* means all)
ALLOW_ORIGINS="*"
```

### Usage

Install go modules

```bash
go mod download
```

run the project with

```bash
go run main.go
```

then you can visit

```python
http://localhost:4000/api/unsplash/photos
http://localhost:4000/api/unsplash/search/photos?query=london
```
