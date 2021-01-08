<template>
  <div>
    <TagFilter></TagFilter>
    <v-data-table
      :items="routes"
      :headers="headers"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:no-data>
        Нет данных
      </template>

      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ title }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer />
          <CreateRouteDialog/>
        </v-toolbar>
      </template>

      <template v-slot:item.method="{ item }">
        <HttpMethodBox :method="item.method"></HttpMethodBox>
      </template>

      <template v-slot:item.tags="{ item }">
        <div class="d-inline-block" v-for="tag in item.tags">
          <v-chip  @click="tagClick(tag)">{{tag}}</v-chip>
        </div>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon small class="mr-2" @click="edit(item)">mdi-pencil</v-icon>
        <v-icon small @click="remove(item)">mdi-delete</v-icon>
      </template>
    </v-data-table>
    <DeleteRouteDialog :route-id="activeItemId" v-model="showDeleteDialog"/>
    <UpdateRouteDialog :route-id="activeItemId" v-model="showEditDialog"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Watch,
} from 'vue-property-decorator';
import { Filter, Route } from '@/views/route/service';
import DictTable from '@/views/base/components/DictTable.vue';
import { DataTableHeader } from 'vuetify';
import CreateRouteDialog from './CreateRouteDialog.vue';
import DeleteRouteDialog from './DeleteRouteDialog.vue';
import UpdateRouteDialog from './UpdateRouteDialog.vue';
import TagFilter from './TagFilter.vue';
import HttpMethodBox from '../../base/components/HttpMethodBox.vue';

@Component({
  components: {
    CreateRouteDialog,
    DeleteRouteDialog,
    UpdateRouteDialog,
    HttpMethodBox,
    TagFilter,
  },
})
export default class TabRouteTable extends DictTable<Route> {
  readonly title = 'Маршруты'

  mounted() {
    const filter = this.$store.direct.getters.route.getFilter;
    this.$store.direct.dispatch.route.GetList(filter);
  }

  @Watch('filter', { deep: true })
  private onFilterChange(filter: Filter) {
    this.$store.direct.dispatch.route.GetList(filter);
  }

  get filter(): Filter {
    return this.$store.direct.getters.route.getFilter;
  }

  tagClick(tagName: string) {
    const filter = this.$store.direct.getters.route.getFilter;
    filter.tags.exclude = false;
    this.$store.direct.commit.route.setFilter(filter);
    this.$store.direct.commit.route.setFilterTagNames([tagName]);
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.route.getRoutes;
  }

  get activeRoute(): Route {
    return this.$store.direct.getters.route.getRoutes.filter(((route) => route.id === this.activeItemId))[0];
  }

  methods = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']

  readonly headers: DataTableHeader[] = [
    { text: 'ID', value: 'id', width: '50px' },
    { text: 'Метод', value: 'method', width: '80px' },
    { text: 'Маршруты', value: 'route' },
    { text: 'Описание', value: 'description' },
    { text: 'Теги', value: 'tags', width: '30%' },
    {
 text: 'Actions', value: 'actions', sortable: false, width: '80px',
},
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
