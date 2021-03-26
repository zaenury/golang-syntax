const JwtHelper = require('./jwtHelper')
const ResponseBody = require('./responseBody')

module.exports = function (roles = []) {
  return async function (req, res, next) {
    const { authorization } = req.headers
    const jwtHelper = new JwtHelper()
    const responseBody = new ResponseBody()

    try {
      if (authorization === undefined || authorization === '' || !authorization)
        return res
          .status(403)
          .json(
            responseBody
              .setStatus(403)
              .setSuccess(false)
              .setMessage('Authorization header not provided or empty')
              .build()
          )

      const token = jwtHelper.parseToken(authorization)
      const result = await jwtHelper.verifyJwt(token)

      if (result === 'invalid signature')
        return res
          .status(403)
          .json(responseBody.setStatus(403).setSuccess(false).setMessage('Invalid format token').build())

      if (result === 'jwt expired')
        return res
          .status(403)
          .json(responseBody.setStatus(403).setSuccess(false).setMessage('Session login has been expired').build())

      if (roles.length < 1)
        return res
          .status(403)
          .json(
            responseBody.setStatus(403).setSuccess(false).setMessage('You dont have authority to access this').build()
          )

      if (roles.indexOf(result && result.data && result.data.role) < 0)
        return res
          .status(403)
          .json(
            responseBody.setStatus(403).setSuccess(false).setMessage('You dont have authority to access this').build()
          )

      req.auth = result.data
      next()
    } catch (error) {
      throw Error(error.message)
    }
  }
}
