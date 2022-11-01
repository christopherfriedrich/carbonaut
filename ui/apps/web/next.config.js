const withTM = require('next-transpile-modules')([
  '@carbonaut-cloud/ui',
  '@carbonaut-cloud/api',
]);

module.exports = withTM({
  reactStrictMode: true,
});
