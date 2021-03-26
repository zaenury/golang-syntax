// const moment = require('moment')
// const CryptoJS = require('crypto-js')

// const config = require('../configs/config')

// const keySize = 256
// const iterations = 100
// const secretKey = config.app.tokenSecret

function randomPassword(length = 8) {
  // var chars = 'abcdefghijklmnopqrstuvwxyz!@#$%^&*()-+<>ABCDEFGHIJKLMNOP1234567890'
  var chars = 'abcdefghijklmnopqrstuvwxyz!@$%*()-+<>ABCDEFGHIJKLMNOP1234567890'
  var pass = ''

  for (var x = 0; x < length; x++) {
    var i = Math.floor(Math.random() * chars.length)
    pass += chars.charAt(i)
  }

  return pass
}

// function hmacEncrypt(message) {
//   const encrypted = CryptoJS.HmacSHA256(message, secretKey)
//   return encrypted.toString(CryptoJS.enc.Base64)
// }

// function encrypt(payload) {
//   if (!payload) {
//     return null
//   }

//   if (isObject(payload)) {
//     payload = JSON.stringify(payload)
//   }

//   let salt = CryptoJS.lib.WordArray.random(128 / 8)
//   let iv = CryptoJS.lib.WordArray.random(128 / 8)
//   let key = CryptoJS.PBKDF2(secretKey, salt, {
//     keySize: keySize / 32,
//     iterations: iterations,
//   })
//   let encrypted = CryptoJS.AES.encrypt(payload, key, {
//     iv: iv,
//     padding: CryptoJS.pad.Pkcs7,
//     mode: CryptoJS.mode.CBC,
//   })

//   return salt.toString() + iv.toString() + encrypted.toString()
// }

// function decrypt(payload) {
//   if (!payload) {
//     return null
//   }

//   let salt = CryptoJS.enc.Hex.parse(payload.substr(0, 32))
//   let iv = CryptoJS.enc.Hex.parse(payload.substr(32, 32))
//   let encrypted = payload.substring(64)
//   let key = CryptoJS.PBKDF2(secretKey, salt, {
//     keySize: keySize / 32,
//     iterations: iterations,
//   })
//   let result = CryptoJS.AES.decrypt(encrypted, key, {
//     iv: iv,
//     padding: CryptoJS.pad.Pkcs7,
//     mode: CryptoJS.mode.CBC,
//   }).toString(CryptoJS.enc.Utf8)

//   return result ? (isObject(result) ? JSON.parse(result) : result) : null
// }

function isList(val) {
  return val != null && typeof val != 'function' && typeof val.length == 'number'
}

function isFunction(val) {
  return val && {}.toString.call(val) === '[object Function]'
}

function isObject(val) {
  return val && {}.toString.call(val) === '[object Object]'
}

function isArray(val) {
  return Array.isArray(val)
}

function basePath(str) {
  var n = str.lastIndexOf('/')
  return str.substr(0, n)
}

function pathName() {
  return window.location.pathname
}

function pathNames(val) {
  const path = val ? val : pathName()
  return path
    .split('/')
    .filter(f => f !== '')
    .map(m => `/${m}`)
}

function parseObject(val) {
  if (!val) {
    return null
  }

  return JSON.parse(val)
}

function dashSpace(str) {
  str = str.replace(/\s/g, '-')
  return uCase(str)
}

function uCase(str) {
  return str.replace(/\w\S*/g, txt => {
    const exceptTxt = txt === 'DKI' || txt === 'DI' ? txt : txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase()
    return exceptTxt
  })
}

function isObjEmpty(obj) {
  for (var key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key)) return false
    // if (obj.hasOwnProperty(key)) return false
  }
  return true
}

function isEmailValid(email) {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)
}

function isPhoneValid(phone) {
  return phone && phone.substring(0, 3) === '+62'
}

function isDateValid(date) {
  return moment(date, 'YYYY-MM-DD', true).isValid()
}

function isTimeValid(time) {
  return moment(time, 'HH:mm', true).isValid()
}

function formatDate(date) {
  if (date != null && moment(date, 'YYYY-MM-DD', true).isValid()) return moment(date).format('YYYY-MM-DD')

  return null
}

module.exports = util = {
  randomPassword,
  // hmacEncrypt,
  // encrypt,
  // decrypt,
  isList,
  isFunction,
  isObject,
  isArray,
  isObjEmpty,
  basePath,
  pathName,
  pathNames,
  parseObject,
  dashSpace,
  uCase,
  isEmailValid,
  isPhoneValid,
  isDateValid,
  isTimeValid,
  formatDate,
}
