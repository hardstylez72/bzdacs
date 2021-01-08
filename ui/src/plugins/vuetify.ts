import Vue from 'vue';
// eslint-disable-next-line @typescript-eslint/ban-ts-ignore
// @ts-ignore
import Vuetify from 'vuetify/lib/framework';
import colors from 'vuetify/lib/util/colors';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: colors.grey.darken1, // #E53935
        secondary: colors.amber.accent3, // #FFCDD2
        accent: colors.indigo.base, // #3F51B5
      },
    },
  },
});
