module.exports = {
  '/dev-api': {
    target: 'http://127.0.0.1:8089',
    pathRewrite: {
      '^/dev-api': ''
    }
  }
}
