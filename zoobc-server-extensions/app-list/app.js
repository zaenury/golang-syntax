const fs = require('fs')
const http = require('http')
const https = require('https')
const express = require('express')

const { config } = require('./configs')
const app = express().set('port', config.app.port)

const server =
  config.app.openSslCertPath && config.app.openSslKeyPath
    ? https.createServer({
        key: fs.readFileSync(config.app.openSslKeyPath),
        cert: fs.readFileSync(config.app.openSslCertPath),
      })
    : http.createServer(app)

require('./extensions/cors')(app)
require('./extensions/compression')(app)
require('./extensions/log')(app)
require('./extensions/cluster')(server)
require('./extensions/mongoose')()

require('./api/routes')(app)

module.exports = app
