module.exports = {
  '/dev-api': {
    target: 'http://127.0.0.1:1323',
    pathRewrite: {
      '^/dev-api': ''
    }
  }
}
