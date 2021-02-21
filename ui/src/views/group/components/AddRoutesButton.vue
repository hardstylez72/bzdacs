<template>
  <c-dialog
    v-model="show"
    max-width="2000px"
  >
    <template v-slot:activator="props">
      <v-btn
        color="primary"

        class="mb-2"
        v-bind="props"
        v-on="props.on"
      >
        {{$t('add-routes-btn')}}
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('add-routes-title')}}
      </v-card-title>

      <GroupRoutesSelectableTable v-model="selected" :items="routes">
        <template v-slot:top>
          <v-toolbar flat>
            <v-spacer />
            <div>
              <v-btn
                v-if="showAddBtn"
                color="primary"
                class="mb-2"
                @click="addSelectedRoutes"
              >
                {{$t('add-selected-routes')}}
              </v-btn>
            </div>
          </v-toolbar>
        </template>
      </GroupRoutesSelectableTable>
      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel-btn')}}</v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </c-dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import GroupRoutesSelectableTable from './GroupRoutesSelectableTable.vue';

@Component({
  components: {
    'c-dialog': () => import('../../common/components/Dialog.vue'),
    GroupRoutesSelectableTable,
  },
})
export default class RoutesTableSelectAddDialog extends Vue {
  show = false

  @Prop({ type: Number, default: -1 })
  private readonly groupId!: number

  entities: Route[] =[]

  selected: Route[] =[]

  mounted() {
    this.$store.direct.dispatch.groupRoute.GetListNotBelongToUser(this.groupId);
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.groupRoute.getRoutesNotBelongToGroup;
  }

  get showAddBtn(): boolean {
    return this.selected.length > 0;
  }

  async addSelectedRoutes() {
    const routes = this.selected;
    const params = routes.map((route) => ({
      groupId: this.groupId,
      routeId: route.id,
    }));

    await this.$store.direct.dispatch.groupRoute.Create(params);
    this.selected = [];
    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<i18n>
{
  "en": {
    "add-routes-btn": "Add routes",
    "add-routes-title": "Adding routes to the group",
    "add-selected-routes": "Add selected routes",
    "cancel-btn": "Cancel"
  },
  "ru": {
    "add-routes-btn": "Добавить маршруты",
    "add-routes-title": "Добавление маршрутов к группе",
    "add-selected-routes": "Добавить выбранные маршруты",
    "cancel-btn": "Отмена"
  }
}
</i18n>
