# web app sample in Go

## Steps

### Configure webpack

At first, install yarn packages

```bash
$ yarn init
$ yarn add -D @babel/cli @babel/core @babel/preset-env babel-loader webpack webpack-cli webpack-dev-server html-webpack-plugin vue-loader vue-template-compiler css-loader vue-style-loader sass-loader
$ yarn add babel-polyfill vue  node-sass axios
```

After that, create [webpack.config.js](webpack.config.js).

Additoinaly, put following setting into `package.json`.

```json
  "scripts": {
    "start": "webpack-dev-server",
    "build": "webpack --optimize-minimize"
  },
```

### Add web assets

Required following files:

- [src/css/main.scss](src/css/main.scss)
- [src/js/index.js](src/js/index.js)
- [src/js/app.vue](src/js/index.js)
- [static/index.html](static/index.html)

### Add go module

Initialize go module.

```sh
$ go mod init github.com/m-mizutani/web-app-go
```

After that, add [main.go](main.go).

## Development

### Run go server

```
$ go run .
```

[air](https://github.com/cosmtrek/air) command with [.air.conf](.air.conf) is recommended for hot reloading.

```
$ air -c .air.conf
```

### Start webpack development server

```
$ npm run start
```

### Open browser

Open http://localhost:8080 to confirm Web UI.

## Deploy

```
$ npm run build
$ go build -o sercer
$ cp -r server static /path/to/server
```
