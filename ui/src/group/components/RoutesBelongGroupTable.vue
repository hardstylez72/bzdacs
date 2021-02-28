<template>
  <div>
    <v-data-table
      v-model="selected"
      :items="items"
      :headers="headers"
      :loading="loading"
      :show-select="showSelect"
      sort-by="calories"
      class="elevation-1"
      hide-default-footer
      :items-per-page="-1"
    >
      <template v-slot:no-data>
        {{$t('no-data')}}
      </template>

      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title> {{$t('title')}}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
          <DeleteRoutesButton v-if="showDeleteButton" :items="selected" @routesDeleted="routesDeleted"/>
          <AddRoutesDialog :group-id="groupId" @routesAdded="routesAdded"/>
        </v-toolbar>
      </template>

      <template v-slot:item.method="{ item }">
        <HttpMethodBox :method="item.method"></HttpMethodBox>
      </template>

      <template v-slot:item.tags="{ item }">
        <div class="d-inline-block" v-for="tag in item.tags">
          <v-chip>{{tag}}</v-chip>
        </div>
      </template>

    </v-data-table>
    <TablePagination v-bind:limit.sync="pageSize" :total="total" v-model="page"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Prop,
} from 'vue-property-decorator';
import { Route } from '@/route/entity';
import ItemTable from '@/base/components/ItemTable.vue';
import { DataTableHeader } from 'vuetify';
import { ListResponse } from '@/common/helpers/types';
import CreateRouteDialog from './CreateDialog.vue';
import DeleteRouteDialog from './DeleteDialog.vue';
import UpdateRouteDialog from './UpdateDialog.vue';
import HttpMethodBox from '../../common/components/HttpMethodBox.vue';
import Breadcrumbs from '../../common/components/Breadcrumbs.vue';
import TablePagination from '../../common/components/TablePagination.vue';
import DeleteRoutesButton from './DeleteRoutesButton.vue';
import AddRoutesButton from './AddRoutesButton.vue';
import AddRoutesDialog from './AddRoutesDialog.vue';

@Component({
  components: {
    CreateRouteDialog,
    DeleteRouteDialog,
    UpdateRouteDialog,
    HttpMethodBox,
    Breadcrumbs,
    TablePagination,
    AddRoutesButton,
    DeleteRoutesButton,
    AddRoutesDialog,
  },
})

export default class RoutesBelongGroupTable extends ItemTable<Route> {
  @Prop() groupId!: number

  showSelect = true

  async loadItems(): Promise<ListResponse<Route>> {
    return this.$store.direct.dispatch.groupRoute.GetList({
      filter: {
        namespaceId: this.queryParams.namespaceId,
        page: this.page,
        pageSize: this.pageSize,
        belongToGroup: true,
        groupId: this.groupId,
      },
    });
  }

  routesDeleted() {
    this.loadItemsWrapper();
    this.selected = [];
  }

  routesAdded() {
    this.loadItemsWrapper();
  }

  get showDeleteButton(): boolean {
    return this.selected.length > 0;
  }

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('method').toString(), value: 'method', width: '100px' },
    { text: this.$t('route').toString(), value: 'route' },
    { text: this.$t('description').toString(), value: 'description' },
    { text: this.$t('tags').toString(), value: 'tags', width: '30%' },
  ]
}
</script>

<style scoped lang="scss">

</style>

<i18n>
{
  "en": {
    "no-data": "No data",
    "title": "Routes",
    "id": "Id",
    "method": "Method",
    "route": "Route",
    "description": "Description",
    "tags": "Tags",
    "actions": "Actions"
  },
  "ru": {
    "no-data": " Нет данных",
    "title": "Муршруты",
    "id": "Id",
    "method": "Метод",
    "route": "Маршрут",
    "description": "Описание",
    "tags": "Теги",
    "actions": "Действия"
  }
}
</i18n>
