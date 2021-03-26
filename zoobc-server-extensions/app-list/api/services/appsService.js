const BaseService = require('./baseService')
const { Apps } = require('../models')

module.exports = class AppsService extends BaseService {
  constructor() {
    super(Apps)
  }

  static parseSelect(select) {
    if (!select) return ['Name', 'Api', 'Active', 'Order', 'Desciption']
    return select
  }

  async findAll(params) {
    const select = AppsService.parseSelect(params && params.select)
    const where = BaseService.parseWhere(params && params.where)
    const order = BaseService.parseOrder(params && params.order)

    return new Promise((resolve, reject) => {
      this.model
        .find(where)
        .select(select)
        .sort(order)
        .sort({ Order: 'asc' })
        .lean()
        .exec((err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
    })
  }
}
