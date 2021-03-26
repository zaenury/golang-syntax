const config = require('../../configs/config')

module.exports = class ResponseBody {
  constructor() {
    this.meta = { status: 200, success: true, message: 'Operations success' }
    this.data = {}
    this.paginate = { page: 0, total: 0, limit: config.app.pageLimit }
  }

  setStatus(status) {
    this.meta.status = status
    return this
  }

  setSuccess(success) {
    this.meta.success = success
    return this
  }

  setMessage(message) {
    this.meta.message = message
    return this
  }

  setData(data) {
    this.data = data
    return this
  }

  setPaginate(paginate) {
    this.paginate = paginate
    return this
  }

  setPage(page) {
    this.paginate.page = page
    return this
  }

  setTotal(total) {
    this.paginate.total = total
    return this
  }

  setLimit(limit) {
    this.paginate.limit = limit
    return this
  }

  build() {
    return {
      meta: this.meta,
      data: this.data,
      paginate: this.paginate,
    }
  }
}
