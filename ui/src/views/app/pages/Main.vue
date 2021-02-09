<template>
  <div v-if="isAuthorized">
    <TreeView/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TreeView from '@/views/tree-menu/components/TreeView.vue';

@Component({
  components: {
    TreeView,
  },
})
export default class Home extends Vue {
  get isAuthorized(): boolean {
    return this.$store.direct.getters.isAuthorized;
  }
}
</script>
