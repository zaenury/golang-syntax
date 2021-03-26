require('dotenv').config()

module.exports = {
  app: {
    port: process.env.PORT || 3030,
    host: process.env.HOST || 'localhost',
    tokenSecret: process.env.TOKEN_SECRET || 'INDWIRA00884d31c5d4766dc624e1225888',
    tokenExpired: 12 /** hours */,
    mainRoute: '/v1',
    redisExpired: 60 /** seconds */,
    openSslKeyPath: process.env.SSL_KEYPATH || null,
    openSslCertPath: process.env.SSL_CERTPATH || null,
    loggerFilePath: './logs/access.log',
    rateLimitSuspendTime: 5,
    rateLimitMaxHitPerIP: 500,
    pageLimit: 10,
  },
  db: {
    port: process.env.DB_PORT || 27017,
    host: process.env.DB_HOST,
    database: process.env.DB_NAME,
    username: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
  },
  redis: {
    port: process.env.RD_PORT || 6379,
    host: process.env.RD_HOST,
    password: process.env.RD_PASSWORD,
  },
  token: {
    audience: 'indowira.user',
    issuer: 'indowira.id-backend',
    subject: 'indowira-access-token',
  },
  header: {
    id: process.env.HEADER_CLIENT_ID || '1234567890',
    secret: process.env.HEADER_CLIENT_SECRET || 'client-secret-key',
  },
}
