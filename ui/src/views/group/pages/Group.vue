<template>
  <div>
    <h2>Группа {{groupC.code}}</h2>
    <GroupRoutesSelectableTable
      v-model="selected"
      :items="routes"
    >
      <template v-slot:top>
        <v-toolbar
          flat
        >
          <v-toolbar-title>{{ dict.title }}</v-toolbar-title>
          <v-divider
            class="mx-4"
            inset
            vertical
          />
          <v-spacer />
          <div>
            <v-btn
              v-if="showDeleteBtn"
              color="primary"
              class="mb-2"
              @click="deleteSelectedRoutes"
            >
              Удалить выбранные маршруты
            </v-btn>
          </div>
          <AddRoutesButton :group-id="groupIdC" />
        </v-toolbar>
      </template>
    </GroupRoutesSelectableTable>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import { Group } from '@/views/group/services/group';
import GroupRoutesSelectableTable from '../components/GroupRoutesSelectableTable.vue';
import AddRoutesButton from '../components/AddRoutesButton.vue';

@Component({
  components: {
    GroupRoutesSelectableTable,
    AddRoutesButton,
  },
})
export default class RoutesTab extends Vue {
  dict = {
    title: 'Маршруты',
  }

  group: Group = {
    id: -1,
    description: 'не известный',
    code: 'Не определена',
  }

  get groupIdC(): number {
    return this.groupId;
  }

  get groupC(): Group {
    return this.group;
  }

  groupId = Number(this.$route.params.id);

  mounted() {
    this.$store.direct.dispatch.groupRoute.GetListBelongToUser(this.groupId);
    this.$store.direct.dispatch.group.GetById(this.groupId).then((group) => {
      this.group = group;
    });
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.groupRoute.getRoutesBelongToGroup;
  }

  get showDeleteBtn(): boolean {
    return this.selected.length > 0;
  }

  selected: Route[] = []

  entities: Route[] = []

  valid = true

  editedIndex = -1

  async deleteSelectedRoutes() {
    const routes = this.selected;
    const groupId = Number(this.$route.params.id);
    const params = routes.map((route) => ({
        groupId,
        routeId: route.id,
      }));

    await this.$store.direct.dispatch.groupRoute.Delete(params);
    this.selected = [];
  }

  readonly headers = [
    {
      text: 'ID',
      value: 'id',
    },
    {
      text: 'Маршрут',
      value: 'route',
    },
    {
      text: 'Метод',
      value: 'method',
    },
    {
      text: 'Описание',
      value: 'description',
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
