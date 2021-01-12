<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="items"
      sort-by="calories"
      class="elevation-1"
      show-select
    >
      <template v-slot:no-data>
       {{$t('no-data')}}
      </template>

      <template v-slot:top>
        <slot name="top" />
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="view(item)">
          mdi-eye
        </v-icon>
      </template>

    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Component } from 'vue-property-decorator';
import SelectableTable from '@/views/base/components/SelectableTable.vue';
import { Group } from '@/views/group/services/group';
import { DataTableHeader } from 'vuetify';

@Component
export default class GroupsBelongUserSelectableTable extends SelectableTable<Group> {
  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('method').toString(), value: 'method', width: '100px' },
    { text: this.$t('code').toString(), value: 'code' },
    { text: this.$t('description').toString(), value: 'description' },
    { text: this.$t('actions').toString(), value: 'actions', width: '110px' },
  ]

  view(group: Group) {
    return this.$router.push({ name: 'Group', params: { id: group.id.toString() } });
  }
}
</script>

<i18n>
{
  "en": {
    "no-data": "No data",
    "title": "Groups",
    "method": "Method",
    "code": "Code",
    "id": "Id",
    "description": "Description",
    "actions": "Actions"
  },
  "ru": {
    "no-data": " Нет данных",
    "title": "Группы",
    "method": "Метод",
    "code": "Код",
    "id": "Id",
    "description": "Описание",
    "actions": "Дейтсвия"
  }
}
</i18n>
