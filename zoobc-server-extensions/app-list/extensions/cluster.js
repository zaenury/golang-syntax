const os = require('os')
const cluster = require('cluster')

const config = require('../configs/config')

module.exports = server => {
  if (cluster.isMaster) {
    const cpus = os.cpus().length
    console.info(`Mode cluster. Forking for ${cpus} CPUs`)
    for (let i = 0; i < cpus; i++) {
      cluster.fork()
    }
  } else {
    server.listen(config.app.port, () =>
      console.info(`Starting API on Port ${config.app.port}. Handled by pid ${process.pid}`)
    )

    process.on('SIGINT', () => {
      server.close(err => {
        if (err) {
          console.error(`Force closing with error ${err}`)
          process.exit(1)
        } else {
          console.info(`Closing API on Port ${config.app.port}`)
          process.exit(0)
        }
      })
    })
  }
}
