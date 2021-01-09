<template>
  <div v-if="isAuthorized">
    <MainTabs />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import MainTabs from '@/views/base/pages/Main.vue';
import LoginPage from '@/views/AdminPage.vue';

@Component({
  components: {
    MainTabs,
  },
})
export default class Home extends Vue {
  get isAuthorized(): boolean {
    return this.$store.direct.getters.isAuthorized;
  }
}
</script>
