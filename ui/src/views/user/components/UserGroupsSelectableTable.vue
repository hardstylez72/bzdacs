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
    { text: 'ID', value: 'id', width: '50px' },
    { text: 'Метод', value: 'method', width: '10%' },
    { text: 'Код', value: 'code' },
    { text: 'Описание', value: 'description' },
    { text: 'Actions', value: 'actions' },
  ]

  view(group: Group) {
    return this.$router.push({ name: 'Group', params: { id: group.id.toString() } });
  }
}
</script>

<style scoped lang="scss">
</style>
