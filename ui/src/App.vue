<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <h1>{{appName}}</h1>
      <LanguageSelector  class="ml-12 language-selector pt-4" />
      <v-spacer></v-spacer>
      <span v-if="login" class="ms-5">{{$t('user')}}: {{login}}</span>
      <v-btn v-if="login" class="ms-5" @click="logout" >{{$t('logout')}}</v-btn>
    </v-app-bar>
    <v-snackbar v-if="showSnackbar" v-model="showSnackbar">{{snackbarMessage}}</v-snackbar>
      <v-main>
        <router-view />
      </v-main>
  </v-app>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import LanguageSelector from '@/views/app/components/LanguageSelector.vue';

@Component({
  components: {
    LanguageSelector,
  },
})
export default class App extends Vue {
  appName = 'BZDACS'

  get isAuthorized(): boolean {
    return this.$store.direct.getters.isAuthorized;
  }

  get login(): string {
    return this.$store.direct.getters.login;
  }

  logout(): Promise<void> {
    return this.$store.direct.dispatch.logout();
  }

  get snackbarMessage(): string {
    return this.$store.direct.getters.snackbarMessage;
  }

  get showSnackbar(): boolean {
    return this.$store.direct.state.showSnackbar;
  }

  set showSnackbar(val: boolean) {
    this.$store.direct.commit.setShowSnackbar(val);
  }

  created() {
    // eslint-disable-next-line no-undef
    window.addEventListener('req-status-403', () => {
      this.$store.direct.commit.showError('user does not have permissions to do this operation');
    });

    // eslint-disable-next-line no-undef
    window.addEventListener('req-status-401', () => {
      this.$store.direct.commit.showError('user is not authorized');
    });

    this.$store.direct.dispatch.userSession();
  }
}
</script>

<style scoped lang="css">
.v-main {
  font-family: "Lobster";
  width: 99%;
  margin-left: 0.5%;
  margin-top: 0.5%;
}

.language-selector {
  width: 100px;
}

@font-face {
  font-family: "Lobster";
  src: local("Lobster"),
  url(../public/fonts/lobster.ttf) format("truetype");
}
</style>

<i18n>
{
  "en": {
    "logout": "Logout",
    "user": "User"
  },
  "ru": {
    "logout": "Выйти",
    "user": "Пользователь"
  }
}
</i18n>
