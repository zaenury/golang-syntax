const cors = require('cors')
const helmeet = require('helmet')
const bodyParser = require('body-parser')

const bodyParserOpt = { limit: '30mb', extended: true }

module.exports = app => {
  app.use(cors())
  app.use(bodyParser.json(bodyParserOpt))
  app.use(bodyParser.urlencoded(bodyParserOpt))
  app.use(helmeet())
}
