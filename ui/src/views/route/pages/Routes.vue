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

export default class RouteTable extends Vue {
  breadcrumbs: any[] = []

  mounted() {
    this.breadcrumbs.push({
      text: this.$route.query.systemName,
    });
    this.breadcrumbs.push({
      text: this.$route.query.namespaceName,
    });
  }
}
</script>

<style scoped lang="scss">

</style>
