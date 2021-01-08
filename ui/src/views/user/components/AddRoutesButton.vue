<template>
  <Dialog
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
        Добавить маршруты
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        Добавление маршрутов к пользователю
      </v-card-title>

      <UserRoutesSelectableTable
        v-model="selected"
        :items="routes"
      >
        <template v-slot:top>
          <v-toolbar
            flat
          >
            <v-spacer />
            <div>
              <v-btn
                v-if="showAddBtn"
                color="primary"
                class="mb-2"
                @click="addSelectedRoutes"
              >
                Добавить выбранные маршруты
              </v-btn>
            </div>
          </v-toolbar>
        </template>
      </UserRoutesSelectableTable>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="blue darken-1"
          text
          @click="close"
        >
          Cancel
        </v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Prop, Watch,
} from 'vue-property-decorator';
import { Group } from '@/views/group/services/group';
import { Route } from '@/views/route/service';
import Dialog from '@/views/base/components/Dialog.vue';
import UserRoutesSelectableTable from './UserRoutesSelectableTable.vue';

@Component({
  components: {
    Dialog,
    UserRoutesSelectableTable,
  },
})
export default class RoutesTableSelectAddDialog extends Vue {
  show = false

  @Prop({ type: Number, default: -1 })
  private readonly userId!: number

  @Watch('show')
  change(value: boolean): void {
    if (value) {
      this.$store.direct.dispatch.userRoute.GetListNotBelongToUser(this.userId);
    }
  }

  entities: Group[] =[]

  selected: Group[] =[]

  get routes(): readonly Route[] {
    return this.$store.direct.getters.userRoute.getRoutesNotBelongToGroup;
  }

  get showAddBtn(): boolean {
    return this.selected.length > 0;
  }

  async addSelectedRoutes() {
    const routes = this.selected;
    const params = routes.map((route) => ({
      userId: this.userId,
      routeId: route.id,
      isExcluded: false,
    }));

    await this.$store.direct.dispatch.userRoute.Create(params);
    await this.$store.direct.dispatch.userRoute.GetListBelongToUser(this.userId);
    this.selected = [];
    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
