<template>
  <div>
    <Breadcrumbs :items="breadcrumbs"/>
    <DictTable/>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

import Breadcrumbs from '@/views/common/components/Breadcrumbs.vue';
import DictTable from '../components/DictTable.vue';

@Component({
  components: {
    DictTable,
    Breadcrumbs,
  },
})
export default class TapUserTable extends Vue {
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
