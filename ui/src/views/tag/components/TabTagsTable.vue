<template>
  <div>
    <v-data-table
      :items="tags"
      :headers="headers"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:no-data>
       {{$t('no-data')}}
      </template>

      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ $t('title') }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
          <CreateTagDialog/>
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
    <DeleteTagDialog :tag-id="activeItemId" v-model="showDeleteDialog"/>
    <UpdateTagDialog :id="activeItemId" v-model="showEditDialog"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import DictTable from '@/views/base/components/DictTable.vue';
import { DataTableHeader } from 'vuetify';
import { Tag } from '@/views/tag/service';
import CreateTagDialog from './CreateTagDialog.vue';
import DeleteTagDialog from './DeleteTagDialog.vue';
import UpdateTagDialog from './UpdateTagDialog.vue';
import HttpMethodBox from '../../base/components/HttpMethodBox.vue';

@Component({
  components: {
    CreateTagDialog,
    DeleteTagDialog,
    UpdateTagDialog,
    HttpMethodBox,
  },
})
export default class TabRouteTable extends DictTable<Route> {
  readonly title = 'Теги'

  mounted() {
    this.$store.direct.dispatch.tag.GetList();
  }

  get tags(): readonly Tag[] {
    return this.$store.direct.getters.tag.getTags;
  }

  get activeRoute(): Tag {
    return this.$store.direct.getters.tag.getTags.filter(((route) => route.id === this.activeItemId))[0];
  }

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('tag').toString(), value: 'name' },
    { text: this.$t('actions').toString(), value: 'actions', width: '80px' },
  ]
}
</script>

<style scoped lang="scss">
.routes-tab-container {
  display: flex;
  height: 1000px;
  justify-content: space-between;
}
.create-route-btn {
  margin: 10px;

}
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
