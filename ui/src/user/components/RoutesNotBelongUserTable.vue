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
          <AddRoutesButton v-if="showAddButton" :items="selected" @routesAdded="routesAdded"/>
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

      <template v-slot:item.groups="{ item }">
        {{item.groups.map(g => g.code).join(', ')}}
      </template>

    </v-data-table>
    <TablePagination v-bind:limit.sync="pageSize" :total="total" v-model="page"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Prop,
} from 'vue-property-decorator';
import ItemTable from '@/base/components/ItemTable.vue';
import { DataTableHeader } from 'vuetify';
import { ListResponse } from '@/common/helpers/types';
import { Route } from '@/route/entity';
import CreateRouteDialog from './CreateDialog.vue';
import DeleteRouteDialog from './DeleteDialog.vue';
import UpdateRouteDialog from './UpdateDialog.vue';
import HttpMethodBox from '../../common/components/HttpMethodBox.vue';
import Breadcrumbs from '../../common/components/Breadcrumbs.vue';
import TablePagination from '../../common/components/TablePagination.vue';
import AddRoutesButton from './AddRoutesButton.vue';

@Component({
  components: {
    CreateRouteDialog,
    DeleteRouteDialog,
    UpdateRouteDialog,
    HttpMethodBox,
    Breadcrumbs,
    TablePagination,
    AddRoutesButton,
  },
})

export default class RoutesNotBelongGroupTable extends ItemTable<Route> {
  @Prop() userId!: number

  showSelect = true

  async loadItems(): Promise<ListResponse<Route>> {
    return this.$store.direct.dispatch.userRoute.GetList({
      filter: {
        namespaceId: this.queryParams.namespaceId,
        page: this.page,
        pageSize: this.pageSize,
        belongToUser: false,
        userId: this.userId,
      },
    });
  }

  routesAdded() {
    this.loadItemsWrapper();
    this.selected = [];
    this.$emit('routesAdded');
  }

  get showAddButton(): boolean {
    return this.selected.length > 0;
  }

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('method').toString(), value: 'method', width: '100px' },
    { text: this.$t('route').toString(), value: 'route' },
    { text: this.$t('description').toString(), value: 'description' },
    { text: this.$t('groups').toString(), value: 'groups' },
  ]
}
</script>

<style scoped lang="scss">

</style>

<i18n>
{
  "en": {
    "title": "Routes",
    "no-data": "No data",
    "id": "Id",
    "method": "Method",
    "route": "Route",
    "description": "Description",
    "groups": "Groups",
    "actions": "Actions",
    "statuses": "Statuses"

  },
  "ru": {
    "title": "Маршруты",
    "no-data": "Нет данных",
    "id": "Id",
    "method": "Метод",
    "route": "Маршрут",
    "description": "Описание",
    "groups": "Группы",
    "actions": "Действия",
    "statuses": "Статусы"
  }
}
</i18n>
