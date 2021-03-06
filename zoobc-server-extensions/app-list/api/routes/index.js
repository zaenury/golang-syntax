const rateLimit = require('express-rate-limit')

const config = require('../../configs/config')
const limiter = rateLimit({
  /** 5 minutes after rate limit reach */
  windowMs: config.app.rateLimitSuspendTime * 60 * 1000,
  /** max hit per IP */
  max: config.app.rateLimitMaxHitPerIP,
})

module.exports = app => {
  app.use(`${config.app.mainRoute}/apps`, limiter, require('./apps'))
}
