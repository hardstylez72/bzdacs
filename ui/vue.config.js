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
  chainWebpack: (config) => {
    config.module
      .rule('i18n')
      .resourceQuery(/blockType=i18n/)
      .type('javascript/auto')
      .use('i18n')
      .loader('@kazupon/vue-i18n-loader');
  },
  // configureWebpack: {
  //   module: {
  //     rules: [
  //       {
  //         resourceQuery: /blockType=i18n/,
  //         type: 'javascript/auto',
  //         loader: '@intlify/vue-i18n-loader',
  //       },
  //       // ...
  //     ],
  //   },
  // },
};
