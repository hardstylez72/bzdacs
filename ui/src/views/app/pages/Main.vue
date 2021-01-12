<template>
  <div v-if="isAuthorized">
    <MainTabs />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import MainTabs from '@/views/app/components/Tabs.vue';

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
