<template>
  <div>
    <div>
      <v-data-table
        :items="items"
        :headers="headers"
        :loading="loading"
        hide-default-footer
        :items-per-page="-1"
        class="elevation-1"
      >
        <template v-slot:no-data>
          {{$t('no-data')}}
        </template>

        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title>{{$t('title')}}</v-toolbar-title>
            <v-divider class="mx-4" inset vertical/>
            <v-spacer />
            <CreateGroupDialog @groupCreated="loadItemsWrapper"/>
          </v-toolbar>
        </template>

        <template v-slot:item.actions="{ item }">
          <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
          <v-icon small @click="remove(item)">mdi-delete</v-icon>
          <v-icon small class="mr-2" @click="view(item)">mdi-eye</v-icon>
        </template>

      </v-data-table>
      <TablePagination v-bind:limit.sync="pageSize" :total="total" v-model="page"/>
    </div>
    <DeleteGroupDialog
      :id="activeItemId"
      v-model="showDeleteDialog"
      @groupDeleted="loadItemsWrapper"
    />
    <UpdateGroupDialog
      :id="activeItemId"
      v-model="showEditDialog"
      @groupUpdated="loadItemsWrapper"
    />
  </div>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';
import { Group } from '@/views/group/entity';
import { DataTableHeader } from 'vuetify';
import { ListResponse } from '@/views/common/helpers/types';
import ItemTable from '../../base/components/ItemTable.vue';
import CreateGroupDialog from './CreateDialog.vue';
import DeleteGroupDialog from './DeleteDialog.vue';
import UpdateGroupDialog from './UpdateDialog.vue';

@Component({
  components: {
    CreateGroupDialog,
    DeleteGroupDialog,
    UpdateGroupDialog,
  },
})
export default class TapGroupTable extends ItemTable<Group> {
  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('code').toString(), value: 'code', width: '15%' },
    { text: this.$t('description').toString(), value: 'description' },
    {
 text: this.$t('actions').toString(), value: 'actions', sortable: false, width: '100px',
},
  ]

  async loadItems(): Promise<ListResponse<Group>> {
    return this.$store.direct.dispatch.group.GetList({
      filter: {
        namespaceId: this.queryParams.namespaceId,
        pageSize: this.pageSize,
        page: this.page,
      },
    });
  }

  view(group: Group) {
    return this.$router.push({
      name: 'Group',
      params: { id: group.id.toString() },
      query: this.$route.query,
    });
  }
}
</script>

<style scoped lang="scss">

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
