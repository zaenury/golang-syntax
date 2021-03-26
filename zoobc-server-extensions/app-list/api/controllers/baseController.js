module.exports = class BaseController {
  constructor(service) {
    this.service = service
  }

  sendSuccess(res, resBody = {}) {
    resBody.meta.status = 200
    resBody.meta.success = true
    res.status(200).json(resBody)
  }

  sendCreated(res, resBody = {}) {
    resBody.meta.status = 201
    resBody.meta.success = true
    res.status(201).json(resBody)
  }

  sendBadRequest(res, resBody = {}) {
    resBody.meta.status = 400
    resBody.meta.success = false
    res.status(400).json(resBody)
  }

  sendNotFound(res, resBody = {}) {
    resBody.meta.status = 404
    resBody.meta.success = false
    res.status(404).json(resBody)
  }

  sendInvalidPayload(res, resBody = {}) {
    resBody.meta.status = 422
    resBody.meta.success = false
    res.status(422).json(resBody)
  }

  sendAlreadyExist(res, resBody = {}) {
    resBody.meta.status = 409
    resBody.meta.success = false
    res.status(409).json(resBody)
  }

  sendInternalServerError(res, resBody = {}) {
    resBody.meta.status = 500
    resBody.meta.success = false
    resBody.paginate = null
    res.status(500).json(resBody)
  }
}
