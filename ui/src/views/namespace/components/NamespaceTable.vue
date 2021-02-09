<template>
  <div>
    <v-data-table
      :items="namespaces"
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
        <v-icon small class="mr-2" @click="view(item)">mdi-eye</v-icon>
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
import { Group } from '@/views/group/services/group';
import { Namespace } from '@/views/namespace/services/entity';
import CreateTagDialog from './CreateNamespaceDialog.vue';
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
export default class TabRouteTable extends DictTable<Namespace> {
  mounted() {
    this.$store.direct.dispatch.tag.GetList();
  }

  view(namespace: Namespace) {
    return this.$router.push({ name: 'Group', params: { id: namespace.id.toString() } });
  }

  get namespaces(): readonly Namespace[] {
    return this.$store.direct.getters.tag.getTags;
  }

  readonly headers: DataTableHeader[] = [
    { text: this.$t('id').toString(), value: 'id', width: '70px' },
    { text: this.$t('tag').toString(), value: 'name' },
    { text: this.$t('actions').toString(), value: 'actions', width: '100px' },
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
    "title": "Namespaces",
    "id": "Id",
    "tag": "Tag",
    "actions": "Actions"
  },
  "ru": {
    "no-data": "Нет данных",
    "title": "Пространства имен",
    "id": "Id",
    "tag": "Тег",
    "actions": "Дейтсвия"
  }
}
</i18n>
