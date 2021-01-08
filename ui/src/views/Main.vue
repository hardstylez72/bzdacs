<template>
  <div v-if="isAuthorized">
    <v-snackbar v-if="showSnackbar" v-model="showSnackbar">{{snackbarMessage}}</v-snackbar>
    <MainTabs/>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import MainTabs from '@/views/base/pages/Main.vue';
import LoginPage from '@/views/AdminPage.vue';

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

  created() {
    this.$store.direct.dispatch.userSession().finally(() => {
      console.log('2');
    });
  }
}
</script>
