const jwt = require('jsonwebtoken')
const config = require('../../configs/config')

module.exports = class JwtHelper {
  constructor() {}

  parseToken(token) {
    if (token.includes('Bearer ')) {
      return token.slice('Bearer '.length)
    }

    throw new Error('Invalid token format')
  }

  createJWT(data) {
    const options = config.token
    options.expiresIn = `${config.app.tokenExpired}h`

    return jwt.sign({ data }, config.app.tokenSecret, options)
  }

  verifyJwt(token) {
    const options = config.token
    options.expiresIn = `${config.app.tokenExpired}h`

    return jwt.verify(token, config.app.tokenSecret, options, (err, decoded) => {
      if (err) return err.message
      return decoded
    })
  }
}
