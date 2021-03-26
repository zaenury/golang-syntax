const ResponseBody = require('./responseBody')
const BaseController = require('../controllers/baseController')

module.exports = class InternalErrors extends BaseController {
  constructor() {
    super()
  }

  send(res, error) {
    const response = new ResponseBody()

    return this.sendInternalServerError(
      res,
      response
        .setStatus(500)
        .setSuccess(false)
        .setPaginate(null)
        .setMessage(error.toString().replace(/"/gi, ''))
        .build()
    )
  }
}
