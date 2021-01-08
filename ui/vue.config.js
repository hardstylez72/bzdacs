const host = 'http://localhost:3000';

module.exports = {
  devServer: {
    host: 'localhost',
    proxy: {
      '^/api': { changeOrigin: true, target: host },
    },
  },
  transpileDependencies: [
    'vuetify',
  ],
};
