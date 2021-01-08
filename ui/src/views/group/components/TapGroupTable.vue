<template>
  <div>
    <v-data-table
      :items="groups"
      :headers="headers"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:no-data>
        Нет данных
      </template>

      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ title }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer/>
          <CreateGroupDialog/>
        </v-toolbar>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="view(item)">mdi-eye</v-icon>
        <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon small @click="remove(item)">mdi-delete</v-icon>
      </template>
    </v-data-table>
    <DeleteGroupDialog
      :id="activeItemId"
      v-model="showDeleteDialog"
    />
    <UpdateGroupDialog
      :id="activeItemId"
      v-model="showEditDialog"
    />
  </div>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';
import { Group } from '@/views/group/services/group';
import { DataTableHeader } from 'vuetify';
import DictTable from '../../base/components/DictTable.vue';
import CreateGroupDialog from './CreateGroupDialog.vue';
import DeleteGroupDialog from './DeleteGroupDialog.vue';
import UpdateGroupDialog from './UpdateGroupDialog.vue';

@Component({
  components: {
    CreateGroupDialog,
    DeleteGroupDialog,
    UpdateGroupDialog,
  },
})
export default class TapGroupTable extends DictTable<Group> {
  readonly title = 'Группы'

  readonly headers: DataTableHeader[] = [
    { text: 'ID', value: 'id', width: '50px' },
    { text: 'Код', value: 'code', width: '15%' },
    { text: 'Описание', value: 'description' },
    {
 text: 'Actions', value: 'actions', sortable: false, width: '100px',
},
  ]

  mounted() {
    this.$store.direct.dispatch.group.GetList();
  }

 get groups(): readonly Group[] {
    return this.$store.direct.getters.group.getEntities;
  }

  view(group: Group) {
    return this.$router.push({ name: 'Group', params: { id: group.id.toString() } });
  }
}
</script>

<style scoped lang="scss">
.group-list-item {
  cursor: pointer;
}
.group-list-item:hover {
  color: blue;
}

.routes-tab-container {
  display: flex;
  height: 1000px;
  justify-content: space-between;
}
.create-route-btn {
  margin: 10px;

}
</style>
