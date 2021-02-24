<template>
  <div>
    <Breadcrumbs :items="breadcrumbs"/>
    <Table/>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

import Breadcrumbs from '@/views/common/components/Breadcrumbs.vue';
import Table from '../components/Table.vue';

@Component({
  components: {
    Table,
    Breadcrumbs,
  },
})
export default class GroupsPage extends Vue {
  breadcrumbs: any[] = []

  mounted() {
    this.breadcrumbs = [{
      text: this.$route.query.systemName,
    }, {
      text: this.$route.query.namespaceName,
    }];
  }
}
</script>

<style scoped lang="scss">

</style>
