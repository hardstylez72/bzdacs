<template>
  <div>
    <h2>Группа <span style="color: teal">{{getGroup.code}}</span></h2>
    <RoutesBelongGroupTable :group-id="groupId" />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/views/group/entity';
import RoutesBelongGroupTable from '../components/RoutesBelongGroupTable.vue';
import AddRoutesButton from '../components/AddRoutesDialog.vue';

@Component({
  components: {
    RoutesBelongGroupTable,
    AddRoutesButton,
  },
})
export default class RoutesTab extends Vue {
  group: Group = {
    id: -1,
    description: 'не известный',
    code: 'Не определена',
    namespaceId: Number(this.$route.query.namespaceId),
    createdAt: '',
    updatedAt: '',
  }

  get getGroup(): Group {
    return this.group;
  }

  async created() {
    this.group = await this.$store.direct.dispatch.group.GetById({ namespaceId: this.group.namespaceId, id: this.groupId });
  }

  groupId = Number(this.$route.params.id);
}
</script>

<style scoped lang="scss">

</style>
