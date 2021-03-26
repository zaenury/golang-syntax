// const moment = require('moment')
const axios = require('axios')

// const util = require('./util')
const { config } = require('../configs')

async function post(endpoint, payload) {
  return axios({
    method: 'post',
    url: `${config.midtrans.host}/${endpoint}`,
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
      Authorization: `Basic ${Buffer.from(config.midtrans.key).toString('base64')}`,
    },
    data: payload,
  })
    .then(response => response.data)
    .catch(error => error)
}

module.exports = api = { post }
