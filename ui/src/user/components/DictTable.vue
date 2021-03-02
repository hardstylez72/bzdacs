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
            <CreateDialog @userCreated="loadItemsWrapper"/>
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
    <DeleteDialog
      :id="activeItemId"
      v-model="showDeleteDialog"
      @userDeleted="loadItemsWrapper"
    />
    <UpdateDialog
      :id="activeItemId"
      v-model="showEditDialog"
      @userUpdated="loadItemsWrapper"
    />
  </div>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';
import { DataTableHeader } from 'vuetify';
import { ListResponse } from '@/common/helpers/types';
import { User } from '@/user/entity';
import ItemTable from '../../base/components/ItemTable.vue';
import CreateDialog from './CreateDialog.vue';
import DeleteDialog from './DeleteDialog.vue';
import UpdateDialog from './UpdateDialog.vue';

@Component({
  components: {
    CreateDialog,
    DeleteDialog,
    UpdateDialog,
  },
})
export default class UserTable extends ItemTable<User> {
  protected headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('externalId').toString(), value: 'externalId' },
    {
      text: this.$t('actions').toString(), value: 'actions', sortable: false, width: '100px',
    },
  ];

  namespaceId = Number(this.$route.query.namespaceId)

  async loadItems(): Promise<ListResponse<User>> {
    return this.$store.direct.dispatch.user.GetList({
      filter: {
        namespaceId: this.namespaceId,
        pageSize: this.pageSize,
        page: this.page,
      },
    });
  }

  view(user: User) {
    return this.$router.push({
      name: 'User',
      params: { id: user.id.toString() },
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
    "externalId": "externalId",
    "description": "Description",
    "actions": "Actions"
  },
  "ru": {
    "no-data": " Нет данных",
    "title": "Группы",
    "code": "Код",
    "id": "Id",
    "externalId": "Идентификатор",
    "description": "Описание",
    "actions": "Дейтсвия"
  }
}
</i18n>
