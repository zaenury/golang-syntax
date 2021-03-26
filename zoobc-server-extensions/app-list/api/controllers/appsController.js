const BaseController = require('./baseController')
const { util } = require('../../utils')
const { AppsService } = require('../services')
const { ResponseBody, InternalErrors } = require('../helpers')

module.exports = class AppsController extends BaseController {
  constructor() {
    super(new AppsService())
  }

  async list(req, res) {
    const resBody = new ResponseBody()
    const intErrors = new InternalErrors()
    const { where, order } = req.query

    try {
      const result = await this.service.findAll({ where, order })
      return this.sendSuccess(res, resBody.setData(result).setMessage('Success').build())
    } catch (error) {
      intErrors.send(res, error.message)
    }
  }

  async create(req, res) {
    const resBody = new ResponseBody()
    const intErrors = new InternalErrors()
    const { Name, Api, Active, Order, Desciption } = req.body

    try {
      if (!Name || !Api || !Active)
        return this.sendInvalidPayload(res, resBody.setMessage('Invalid format input').build())

      const data = await this.service.findOne({ where: { Api } })
      if (!util.isObjEmpty(data)) return this.sendAlreadyExist(res, resBody.setMessage('API already exist').build())

      const payload = { Name, Api, Active, Order, Desciption }
      const result = await this.service.create(payload)
      return this.sendSuccess(res, resBody.setData(result).setMessage('Success created data').build())
    } catch (error) {
      intErrors.send(res, error.message)
    }
  }
}
