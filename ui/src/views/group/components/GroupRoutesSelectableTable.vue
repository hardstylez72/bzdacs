<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headersC"
      :items="items"
      sort-by="calories"
      class="elevation-1"
      show-select
    >
      <template slot="items" slot-scope="props">
        <td>{{ props.item.id }}</td>
        <td class="text-xs-right">{{ props.item.method }}</td>
      </template>

      <template v-slot:no-data>
        {{$t('no-data')}}
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
import HttpMethodBox from '@/views/common/components/HttpMethodBox.vue';

@Component({
  components: {
    HttpMethodBox,
  },
})
export default class GroupRoutesSelectableTable extends SelectableTable<Route> {
  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '7%' },
    { text: this.$t('method').toString(), value: 'method', width: '100px' },
    { text: this.$t('route').toString(), value: 'route' },
    { text: this.$t('description').toString(), value: 'description' },
  ]

  get headersC() {
    return this.headers;
  }
}

</script>

<style scoped lang="scss">
</style>

<i18n>
{
  "en": {
    "no-data": "No data",
    "id": "Id",
    "method": "Method",
    "route": "Route",
    "description": "Description"
  },
  "ru": {
    "no-data": " Нет данных",
    "id": "Id",
    "method": "Метод",
    "route": "Маршрут",
    "description": "Описание"
  }
}
</i18n>
