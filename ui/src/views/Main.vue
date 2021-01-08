<template>
  <div v-if="!loading">
    <v-snackbar v-if="showSnackbar" v-model="showSnackbar">{{snackbarMessage}}</v-snackbar>
    <MainTabs v-if="isAuthorized"/>
    <LoginPage v-else/>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import MainTabs from '@/views/base/pages/Main.vue';
import LoginPage from '@/views/LoginPage.vue';

@Component({
  components: {
    MainTabs,
    LoginPage,
  },
})
export default class Home extends Vue {
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

  loading = true

  created() {
      this.$store.direct.dispatch.login().finally(() => {
        this.loading = false;
      });
  }
}
</script>
