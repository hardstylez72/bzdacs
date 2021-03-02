<template>
  <div>
    <v-data-table
      :items="items"
      :headers="headers"
      :loading="loading"
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
          <CreateRouteDialog @routeCreated="loadItemsWrapper"/>
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

      <template v-slot:item.actions="{ item }">
        <v-icon data-test="route-edit-action" small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon data-test="route-delete-action" small @click="remove(item)">mdi-delete</v-icon>
      </template>
    </v-data-table>
    <TablePagination v-bind:limit.sync="pageSize" :total="total" v-model="page"/>
    <DeleteRouteDialog :id="activeItemId" v-model="showDeleteDialog" @routeDeleted="loadItemsWrapper"/>
    <UpdateRouteDialog :route-id="activeItemId" v-model="showEditDialog" @routeUpdated="loadItemsWrapper"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Watch,
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

@Component({
  components: {
    CreateRouteDialog,
    DeleteRouteDialog,
    UpdateRouteDialog,
    HttpMethodBox,
    Breadcrumbs,
    TablePagination,
  },
})

export default class TabRouteTable extends ItemTable<Route> {
  namespaceId = Number(this.$route.query.namespaceId)

  async loadItems(): Promise<ListResponse<Route>> {
    return this.$store.direct.dispatch.route.GetList({
      filter: {
        namespaceId: this.namespaceId,
        page: this.page,
        pageSize: this.pageSize,
      },
    });
  }

  methods = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('method').toString(), value: 'method', width: '100px' },
    { text: this.$t('route').toString(), value: 'route' },
    { text: this.$t('description').toString(), value: 'description' },
    { text: this.$t('tags').toString(), value: 'tags', width: '30%' },
    {
 text: this.$t('actions').toString(), value: 'actions', sortable: false, width: '80px',
},
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
