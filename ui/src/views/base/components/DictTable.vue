<template>
  <div>
    <v-data-table
      v-bind="tableAttrs"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:no-data>
        Нет данных
      </template>

      <template v-slot:top>
        <v-toolbar
          flat
        >
          <v-toolbar-title>{{ title }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
        </v-toolbar>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon small @click="remove(item)">mdi-delete</v-icon>
        <v-icon small class="mr-2" @click="view(item)">mdi-eye</v-icon>
      </template>

      <template v-slot:item.tags="{ item }">
        <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon small @click="remove(item)">mdi-delete</v-icon>
        <v-icon small class="mr-2" @click="view(item)">mdi-eye</v-icon>
      </template>

    </v-data-table>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { DataTableHeader } from 'vuetify';
import { Entity } from '@/views/base/services/entity';

@Component
export default class DictTable<T extends Entity> extends Vue {
  protected items: T[] = []

  protected title = ''

  protected headers: DataTableHeader[] = [];

  protected showDeleteDialog = false

  protected showEditDialog = false

   activeItemId = -1

  protected get tableAttrs() {
    return {
      headers: this.headers,
      items: this.items,
    };
  }

  protected edit(item: T): void {
    this.showEditDialog = true;
    this.activeItemId = item.id;
  }

  protected view(item: T): void {
    this.activeItemId = item.id;
  }

  protected remove(item: T): void {
    this.showDeleteDialog = true;
   this.activeItemId = item.id;
  }
}
</script>

<style scoped lang="scss">

</style>
