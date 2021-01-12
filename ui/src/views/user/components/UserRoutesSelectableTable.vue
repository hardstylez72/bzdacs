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
        <slot name="item.actions" v-bind="{ item }"/>
      </template>

      <template v-slot:item.statuses="{ item }">
        <slot name="item.statuses" v-bind="{ item }"/>
      </template>

      <template v-slot:item.method="{ item }">
        <HttpMethodBox :method="item.method"></HttpMethodBox>
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
      { text: this.$t('route').toString(), value: 'route' },
      { text: this.$t('description').toString(), value: 'description' },
      { text: this.$t('groupCodes').toString(), value: 'groupCodes' },
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
    "id": "Id",
    "method": "Method",
    "route": "Route",
    "description": "Description",
    "groupCodes": "Groups"
  },
  "ru": {
    "no-data": " Нет данных",
    "id": "Id",
    "method": "Метод",
    "route": "Маршрут",
    "description": "Описание",
    "groupCodes": "Группы"
  }
}
</i18n>
