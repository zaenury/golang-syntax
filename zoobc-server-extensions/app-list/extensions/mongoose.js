const saslprep = require('saslprep')
const mongoose = require('mongoose')

const config = require('../configs/config')

function connecting() {
  const uri = `mongodb://${config.db.host}:${config.db.port}/${config.db.database}`
  const options = {
    user: config.db.username,
    pass: config.db.password ? saslprep(config.db.password) : config.db.password,
    useNewUrlParser: true,
    useUnifiedTopology: true,
    useFindAndModify: false,
    useCreateIndex: true,
  }

  return mongoose.connect(uri, options, err => {
    if (err) {
      console.error(`MongoDB connection error - retrying in 5 sec\n${err}`)
      setTimeout(connecting, 5000)
    } else console.info('MongoDB connection succeess')
  })
}

module.exports = () => connecting()
