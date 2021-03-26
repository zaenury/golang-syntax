const router = require('express').Router()
const { AppsController } = require('../controllers')

const apps = new AppsController()

router.get('/', (req, res) => {
  apps.list(req, res)
})

router.post('/', (req, res) => {
  apps.create(req, res)
})

module.exports = router
