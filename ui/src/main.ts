import Vue from 'vue';
import VueI18n from 'vue-i18n';
import App from './views/app/pages/App.vue';
import router from './views/app/router';
import store from './views/base/store';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
Vue.use(VueI18n);

const i18n = new VueI18n({
  sync: true,
  locale: 'en',
});

new Vue({
  router,
  store: store.original,
  vuetify,
  i18n,
  render: (h) => h(App),
}).$mount('#app');
