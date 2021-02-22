<template>
  <div>
    <v-data-table
      :items="getTags"
      :headers="headers"
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
          <v-toolbar-title>{{ $t('title') }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
          <CreateTagDialog @tagCreated="loadTags"/>
        </v-toolbar>
      </template>

      <template v-slot:item.method="{ item }">
        <HttpMethodBox :method="item.method"></HttpMethodBox>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon small @click="remove(item)">mdi-delete</v-icon>
      </template>
    </v-data-table>
    <TablePagination v-bind:limit.sync="pageSize" :total="total" v-model="page"/>
    <DeleteTagDialog :tag-id="activeItemId" v-model="showDeleteDialog" @tagDeleted="loadTags"/>
    <UpdateTagDialog :id="activeItemId" v-model="showEditDialog" @tagUpdated="loadTags"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Watch,
} from 'vue-property-decorator';
import { Route } from '@/views/route/entity';
import DictTable from '@/views/base/components/DictTable.vue';
import { DataTableHeader } from 'vuetify';
import { Tag } from '@/views/tag/entity';
import { QueryParams } from '@/views/tree-menu/entity';
import CreateTagDialog from '../components/CreateTagDialog.vue';
import DeleteTagDialog from '../components/DeleteTagDialog.vue';
import UpdateTagDialog from '../components/UpdateTagDialog.vue';
import HttpMethodBox from '../../common/components/HttpMethodBox.vue';
import TablePagination from '../../common/components/TablePagination.vue';

@Component({
  components: {
    CreateTagDialog,
    DeleteTagDialog,
    UpdateTagDialog,
    HttpMethodBox,
    TablePagination,
  },
})
export default class TabRouteTable extends DictTable<Route> {
  pageSize =10

  page = 1

  total = 0

  tags = []

  @Watch('pageSize')
  paginationChanged(pageSize: number) {
    this.page = 1;
    this.pageSize = pageSize;
    this.loadTags();
  }

  @Watch('page')
  pageChanged(pageSize: number) {
    this.page = pageSize;
    this.loadTags();
  }

  queryParams = this.getViewTreeQueryParams()

  @Watch('$route.query')
  routerChanged() {
    this.queryParams = this.getViewTreeQueryParams();
    this.loadTags();
  }

  getViewTreeQueryParams(): QueryParams {
    return {
      systemName: this.$route.query.systemName,
      systemId: Number(this.$route.query.systemId),
      namespaceName: this.$route.query.namespaceName,
      namespaceId: Number(this.$route.query.namespaceId),
    };
  }

  async loadTags() {
    const list = await this.$store.direct.dispatch.tag.GetList({
      filter: {
        namespaceId: this.queryParams.namespaceId,
        page: this.page,
        pageSize: this.pageSize,
      },
    });
    this.total = list.total;
    this.tags = list.items;
  }

  mounted() {
    this.loadTags();
  }

  get getTags(): readonly Tag[] {
    return this.tags;
  }

  get activeRoute(): Tag {
    return this.tags.filter(((tag) => tag.id === this.activeItemId))[0];
  }

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('tag').toString(), value: 'name' },
    { text: this.$t('actions').toString(), value: 'actions', width: '80px' },
  ]
}
</script>

<style scoped lang="scss">

</style>

<i18n>
{
  "en": {
    "no-data": "No data",
    "title": "Tags",
    "id": "Id",
    "tag": "Tag",
    "actions": "Actions"
  },
  "ru": {
    "no-data": "Нет данных",
    "title": "Теги",
    "id": "Id",
    "tag": "Тег",
    "actions": "Дейтсвия"
  }
}
</i18n>
