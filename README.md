# pdf-api
Convert `HTML` to `PDF` with a single endpoint.

## Install
Be sure to install [wkhtmltopdf](https://wkhtmltopdf.org/) before using `pdf-api`.

## Launch
```
$ PORT=3000 pdf-api
```

## Endpoints
### Generate
#### Arguments
```
(String) orientation: "[portrait|landscape]"
(String) url: ""
(String) html: ""
```
`url` or `html` are required

#### Request
```request
GET /generate
```
```
?orientation=[portrait|landscape] &url=https://google.fr &html=<html>Hello</html>
```

#### Response
##### Error
```json
{
  "msg": "error message"
}
```

##### Success
Your pdf file

#### Request
```request
POST /generate
```
```json
{
  "orientation": "[portrait|landscape]",
  "url": "https://google.fr",
  "html": "<html>Hello</html>"
}
```

#### Response
##### Error
```json
{
  "msg": "error message"
}
```

##### Success
Your pdf file

## License
[MIT](https://github.com/ItsJimi/pdf-api/blob/master/LICENSE)
