<template>
  <div>
    <v-data-table
      :items="users"
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
          <CreateUserDialog/>
        </v-toolbar>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="view(item)">mdi-eye</v-icon>
        <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon small @click="remove(item)">mdi-delete</v-icon>
      </template>
    </v-data-table>
    <DeleteUserDialog
      :id="activeItemId"
      v-model="showDeleteDialog"
    />
    <UpdateUserDialog
      :id="activeItemId"
      v-model="showEditDialog"
    />
  </div>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';

import { DataTableHeader } from 'vuetify';
import { User } from '../services/user';
import DictTable from '../../base/components/DictTable.vue';
import CreateUserDialog from './CreateUserDialog.vue';
import DeleteUserDialog from './DeleteUserDialog.vue';
import UpdateUserDialog from './UpdateUserDialog.vue';

@Component({
  components: {
    CreateUserDialog,
    DeleteUserDialog,
    UpdateUserDialog,
  },
})
export default class TapUserTable extends DictTable<User> {
  protected title = 'Пользователи'

  get users(): readonly User[] {
    return this.$store.direct.getters.user.getEntities;
  }

  mounted() {
    this.$store.direct.dispatch.user.GetList();
  }

  protected view(item: User): void {
     this.$router.push({ name: 'User', params: { id: item.id.toString() } });
  }

  protected headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('externalId').toString(), value: 'externalId' },
    {
 text: this.$t('actions').toString(), value: 'actions', sortable: false, width: '100px',
},
  ];
}
</script>
<i18n>
{
  "en": {
    "no-data": "No data",
    "title": "Users",
    "id": "Id",
    "externalId": "External Id",
    "actions": "Actions"
  },
  "ru": {
    "no-data": "Нет данных",
    "title": "Пользователи",
    "id": "Id",
    "externalId": "Внешний идентификатор",
    "actions": "Дейтсвия"
  }
}
</i18n>
