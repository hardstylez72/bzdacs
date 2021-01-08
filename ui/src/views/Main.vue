<template>
  <div v-if="!loading">
    <MainTabs v-if="isAuthorized"/>
    <LoginPage v-else/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
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

  loading = true

  created() {
      this.$store.direct.dispatch.login().finally(() => {
        this.loading = false;
      });
  }
}
</script>
