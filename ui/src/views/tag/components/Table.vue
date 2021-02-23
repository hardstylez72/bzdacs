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
          <v-toolbar-title>{{ $t('title') }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
          <CreateTagDialog @tagCreated="loadItemsWrapper"/>
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
    <DeleteTagDialog :tag-id="activeItemId" v-model="showDeleteDialog" @tagDeleted="loadItemsWrapper"/>
    <UpdateTagDialog :id="activeItemId" v-model="showEditDialog" @tagUpdated="loadItemsWrapper"/>
  </div>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';
import ItemTable from '@/views/base/components/ItemTable.vue';
import { DataTableHeader } from 'vuetify';
import { Tag } from '@/views/tag/entity';
import CreateTagDialog from './CreateTagDialog.vue';
import DeleteTagDialog from './DeleteTagDialog.vue';
import UpdateTagDialog from './UpdateTagDialog.vue';
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
export default class TabRouteTable extends ItemTable<Tag> {
  async loadItems() {
    return this.$store.direct.dispatch.tag.GetList({
      filter: {
        namespaceId: this.queryParams.namespaceId,
        page: this.page,
        pageSize: this.pageSize,
      },
    });
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
