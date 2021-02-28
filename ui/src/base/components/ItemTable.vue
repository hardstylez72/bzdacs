<template>
  <div>
    <v-data-table
      v-model="selected"
      :items="items"
      :headers="headers"
      hide-default-footer
      :items-per-page="-1"
      class="elevation-1"
      :loading="loading"
      :show-select="showSelect"
    >
      <template v-slot:no-data>
        {{$t('no-data')}}
      </template>

      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{$t('title')}}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
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
</template>

<script lang="ts">
import {
  Component, Vue, Watch,
} from 'vue-property-decorator';
import { DataTableHeader } from 'vuetify';
import { Entity } from '@/base/entity';
import TablePagination from '@/common/components/TablePagination.vue';
import { QueryParams } from '@/tree-menu/helper';
import { ListResponse } from '@/common/helpers/types';

@Component({
  components: {
    TablePagination,
  },
})
export default class ItemTable<T extends Entity> extends Vue {
  protected items: T[] = []

  protected selected: T[] = []

  protected loading = false

  protected showSelect = false

  protected headers: DataTableHeader[] = [];

  protected showDeleteDialog = false

  protected showEditDialog = false

  protected activeItemId = -1

  protected pageSize =10

  protected page = 1

  protected total = 0

  @Watch('items', { deep: true })
  itemsChanged(items: T[]) {
    this.items = items;
  }

  @Watch('pageSize')
  protected paginationChanged(pageSize: number) {
    this.page = 1;
    this.pageSize = pageSize;
    this.loadItemsWrapper();
  }

  @Watch('page')
  protected pageChanged(pageSize: number) {
    this.page = pageSize;
    this.loadItemsWrapper();
  }

  queryParams = this.getViewTreeQueryParams()

  @Watch('$route.query')
  protected routerChanged() {
    this.queryParams = this.getViewTreeQueryParams();
    this.loadItemsWrapper();
  }

  protected getViewTreeQueryParams(): QueryParams {
    return {
      systemName: this.$route.query.systemName,
      systemId: Number(this.$route.query.systemId),
      namespaceName: this.$route.query.namespaceName,
      namespaceId: Number(this.$route.query.namespaceId),
    };
  }

  async loadItems(): Promise<ListResponse<T>> {
    console.error('must be overwtitten');
    return {
      items: [],
      total: 1,
    };
  }

  async loadItemsWrapper(): Promise<T[]> {
    this.loading = true;
    const res = await this.loadItems().finally(() => {
      this.loading = false;
    });
    this.items = res.items;
    this.total = res.total;
    return this.items;
  }

  created() {
    this.loadItemsWrapper();
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
