<template>
  <div>
    <Breadcrumbs :items="breadcrumbs"/>
    <v-data-table
      :items="groups"
      :headers="headers"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:no-data>
        {{$t('no-data')}}
      </template>

      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ $t('title') }}</v-toolbar-title>
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
import Breadcrumbs from '@/views/common/components/Breadcrumbs.vue';
import DictTable from '../../base/components/DictTable.vue';
import CreateGroupDialog from '../components/CreateGroupDialog.vue';
import DeleteGroupDialog from '../components/DeleteGroupDialog.vue';
import UpdateGroupDialog from '../components/UpdateGroupDialog.vue';

@Component({
  components: {
    CreateGroupDialog,
    DeleteGroupDialog,
    UpdateGroupDialog,
    Breadcrumbs,
  },
})
export default class TapGroupTable extends DictTable<Group> {
  readonly title = 'Группы'

  breadcrumbs: any[] = []

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('code').toString(), value: 'code', width: '15%' },
    { text: this.$t('description').toString(), value: 'description' },
    {
 text: this.$t('actions').toString(), value: 'actions', sortable: false, width: '100px',
},
  ]

  mounted() {
    this.breadcrumbs.push({
      text: this.$route.query.systemName,
    });
    this.breadcrumbs.push({
      text: this.$route.query.namespaceName,
    });
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
<i18n>
{
  "en": {
    "no-data": "No data",
    "title": "Groups",
    "code": "Code",
    "id": "Id",
    "description": "Description",
    "actions": "Actions"
  },
  "ru": {
    "no-data": " Нет данных",
    "title": "Группы",
    "code": "Код",
    "id": "Id",
    "description": "Описание",
    "actions": "Дейтсвия"
  }
}
</i18n>
