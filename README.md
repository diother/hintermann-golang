### Project structure

```
/src
  /projects
    /2024_03_cool
      body.md
      meta.json
      /media
        image1.webp
        image2.webp

/dist
  /projects
    /cool-project
      index.html
      image1.webp
      image2.webp
```

### Static page structure

Found in `internal/views/static_pages.json`

```
[
  {
    "slug": "home",
    "path": "/",
    "lastMod": "2025-05-10"
  }
]
```

To render static pages, the parser looks for a Golang template with the name of the `slug` of the page

```
{{- define "home" -}}
  ...
{{- end -}}
```
