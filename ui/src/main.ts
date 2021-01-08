import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './views/base/router';
import store from './views/base/store';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;

new Vue({
  router,
  store: store.original,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
