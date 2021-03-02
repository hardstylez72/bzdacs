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
          <DeleteGroupsButton v-if="showDeleteButton" :items="selected" @groupsDeleted="groupsDeleted"/>
          <AddGroupsDialog :user-id="userId" @groupsAdded="groupsAdded"/>
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
        <v-icon small class="mr-2" @click="view(item)">
          mdi-eye
        </v-icon>
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
import { Group } from '@/group/entity';
import CreateRouteDialog from './CreateDialog.vue';
import DeleteRouteDialog from './DeleteDialog.vue';
import UpdateRouteDialog from './UpdateDialog.vue';
import HttpMethodBox from '../../common/components/HttpMethodBox.vue';
import Breadcrumbs from '../../common/components/Breadcrumbs.vue';
import TablePagination from '../../common/components/TablePagination.vue';
import DeleteGroupsButton from './DeleteGroupsButton.vue';
import AddRoutesButton from './AddRoutesButton.vue';
import AddGroupsDialog from './AddGroupsDialog.vue';

@Component({
  components: {
    CreateRouteDialog,
    DeleteRouteDialog,
    UpdateRouteDialog,
    HttpMethodBox,
    Breadcrumbs,
    TablePagination,
    AddRoutesButton,
    DeleteGroupsButton,
    AddGroupsDialog,
  },
})

export default class GroupsBelongUserTable extends ItemTable<Group> {
  @Prop() userId!: number

  showSelect = true

  namespaceId = Number(this.$route.query.namespaceId)

  async loadItems(): Promise<ListResponse<Group>> {
    return this.$store.direct.dispatch.userGroup.GetList({
      filter: {
        namespaceId: this.namespaceId,
        page: this.page,
        pageSize: this.pageSize,
        belongToUser: true,
        userId: this.userId,
      },
    });
  }

  groupsDeleted() {
    this.loadItemsWrapper();
    this.selected = [];
  }

  groupsAdded() {
    this.loadItemsWrapper();
  }

  get showDeleteButton(): boolean {
    return this.selected.length > 0;
  }

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('code').toString(), value: 'code' },
    { text: this.$t('description').toString(), value: 'description' },
    { text: this.$t('actions').toString(), value: 'actions', width: '110px' },
  ]

  view(group: Group) {
    return this.$router.push({
      name: 'Group',
      params: { id: group.id.toString() },
      query: this.$route.query,
    });
  }
}
</script>

<style scoped lang="scss">

</style>

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
