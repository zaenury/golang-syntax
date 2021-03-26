const { config } = require('../../configs')

module.exports = class BaseService {
  constructor(model) {
    this.model = model
  }

  static parseWhere(where) {
    if (!where) return {}
    return where
  }

  static parseOrder(order) {
    if (!order) return { _id: 'asc' }
    if (order[0] === '-') {
      return { [order.slice(1)]: 'desc' }
    }
    return { [order]: 'asc' }
  }

  async paginate({ page, limit = config.app.pageLimit, where, order }) {
    page = page !== undefined ? parseInt(page) : 1
    limit = limit !== undefined ? parseInt(limit) : parseInt(pageLimit)
    order = order !== undefined ? BaseService.parseOrder(order) : { _id: 'asc' }

    var findWhere = {}
    if (where) {
      const splitWhere = where.split(',')
      var NewWhere = []
      splitWhere.forEach(element => {
        NewWhere.push({ [element.split(':')[0]]: element.split(':')[1].toString() })
      })

      findWhere = NewWhere && NewWhere.length > 0 ? { $or: NewWhere } : { [NewWhere[0]]: NewWhere[1].toString() }
    }

    return new Promise((resolve, reject) => {
      this.model.countDocuments(findWhere, (err, total) => {
        if (err) return reject(err)

        this.model
          .find(findWhere)
          .select()
          .skip((page - 1) * limit)
          .limit(limit)
          .sort(order)
          .exec((err, data) => {
            if (err) return reject(err)
            return resolve({
              data,
              paginate: {
                page: parseInt(page),
                count: data.length,
                total,
              },
            })
          })
      })
    })
  }

  async paginateUser({ page, limit = config.app.pageLimit, where, order, user }) {
    page = page !== undefined ? parseInt(page) : 1
    limit = limit !== undefined ? parseInt(limit) : parseInt(pageLimit)
    order = order !== undefined ? BaseService.parseOrder(order) : { _id: 'asc' }

    var findWhere = { user }
    if (where) {
      const splitWhere = where.split(',')
      var NewWhere = []
      splitWhere.forEach(element => {
        NewWhere.push({ [element.split(':')[0]]: element.split(':')[1].toString() })
      })

      findWhere = NewWhere && NewWhere.length > 0 ? { $or: NewWhere } : { [NewWhere[0]]: NewWhere[1].toString() }
    }

    return new Promise((resolve, reject) => {
      this.model.countDocuments(findWhere, (err, total) => {
        if (err) return reject(err)

        this.model
          .find(findWhere)
          .select()
          .skip((page - 1) * limit)
          .limit(limit)
          .sort(order)
          .lean()
          .exec((err, data) => {
            if (err) return reject(err)
            return resolve({
              data,
              paginate: {
                page: parseInt(page),
                count: data.length,
                total,
              },
            })
          })
      })
    })
  }

  /**
   *
   * @param {where: <object>} params
   */
  async findOne(params) {
    const where = BaseService.parseWhere(params && params.where)

    return new Promise((resolve, reject) => {
      this.model
        .findOne(where)
        .select()
        .exec((err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
    })
  }

  async findOneUser({ where, user }) {
    return new Promise((resolve, reject) => {
      this.model
        .findOne({ ...where, user })
        .select()
        .exec((err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
    })
  }

  /**
   * @param {where: <object>, order: <string>} params
   */
  async findOnePopulate(params, populate) {
    const where = BaseService.parseWhere(params && params.where)
    const order = BaseService.parseOrder(params && params.order)

    return new Promise((resolve, reject) => {
      this.model
        .findOne(where)
        .populate(populate)
        .select()
        .sort(order)
        .exec((err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
    })
  }

  /**
   * @param {where: <object>, order: <string>} params
   */
  async findAll(params) {
    const where = BaseService.parseWhere(params && params.where)
    const order = BaseService.parseOrder(params && params.order)

    return new Promise((resolve, reject) => {
      this.model
        .find(where)
        .select()
        .sort(order)
        .lean()
        .exec((err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
    })
  }

  async create(payload) {
    try {
      return new Promise((resolve, reject) => {
        this.model.create(payload, (err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
      })
    } catch (error) {
      throw Error(error.message)
    }
  }

  async update(key, payload) {
    try {
      return new Promise((resolve, reject) => {
        this.model.findOneAndUpdate(key, { $set: payload }, { new: true, upsert: true }, (err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
      })
    } catch (error) {
      throw Error(error.message)
    }
  }

  async destroy(key) {
    try {
      return new Promise((resolve, reject) => {
        this.model.deleteMany(key, (err, res) => {
          if (err) return reject(err)
          return resolve(res)
        })
      })
    } catch (error) {
      throw Error(error.message)
    }
  }
}
