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
        Нет данных
      </template>

      <template v-slot:top>
        <slot name="top" />
      </template>

      <template v-slot:item.method="{ item }">
        <HttpMethodBox :method="item.method"></HttpMethodBox>
      </template>

    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Component } from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import SelectableTable from '@/views/base/components/SelectableTable.vue';
import { DataTableHeader } from 'vuetify';
import HttpMethodBox from '@/views/base/components/HttpMethodBox.vue';

@Component({
  components: {
    HttpMethodBox,
  },
})
export default class GroupRoutesSelectableTable extends SelectableTable<Route> {
  readonly headers: DataTableHeader[] = [
    { text: 'ID', value: 'id', width: '5%' },
    { text: 'Метод', value: 'method', width: '80px' },
    { text: 'Маршрут', value: 'route' },
    { text: 'Описание', value: 'description' },
  ]
}

</script>

<style scoped lang="scss">
</style>
