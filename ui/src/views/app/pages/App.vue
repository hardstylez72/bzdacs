<template>
  <v-app :key="$i18n.locale" style="display: flex">
    <Nav class="top-nav-bar"/>
    <TreeView v-if="isAuthorized" class="left-side-menu"/>
    <v-snackbar v-if="showSnackbar" v-model="showSnackbar">{{snackbarMessage}}</v-snackbar>
      <v-main>
        <v-container :class="isAuthorized ? 'app-body-container' : ''" fluid>
          <router-view />
        </v-container>
      </v-main>
  </v-app>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import Nav from '@/views/app/components/Nav.vue';
import TreeView from '@/views/tree-menu/components/TreeView.vue';

@Component({
  components: {
    Nav,
    TreeView,
  },
})
export default class App extends Vue {
  get isAuthorized(): boolean {
    return this.$store.direct.getters.isAuthorized;
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
      this.$store.direct.commit.showError(this.$t('403').toString());
    });

    // eslint-disable-next-line no-undef
    window.addEventListener('req-status-401', () => {
      this.$store.direct.commit.showError(this.$t('401').toString());
    });

    this.$store.direct.dispatch.userSession().catch(() => {
      console.log('login');
      this.$router.push({ name: 'Login' });
    });
  }
}
</script>

<style scoped lang="css">

.app-body-container {
  margin-left: 325px;
  width: auto;
}

.left-side-menu {
  z-index: 1;
  margin-top: 70px;
  width: 320px;
  height: 100%;
  position: fixed;
}

@font-face {
  font-family: "Lobster";
  src: local("Lobster"),
  url(../../../../public/fonts/lobster.ttf) format("truetype");
}
</style>

<i18n>
{
  "en": {
    "logout": "Logout",
    "user": "User",
    "401": "user is not authorized",
    "403": "user does not have permissions to do this operation"
  },
  "ru": {
    "logout": "Выйти",
    "user": "Пользователь",
    "401": "Пользователь не авторизован",
    "403": "У вас нет прав на выполнение этого действия"
  }
}
</i18n>
