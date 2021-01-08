<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">Новый маршрут</v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{title}}
      </v-card-title>
      <v-card-text>

        <RouteForm v-model="route">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">Отмена</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableCreateButton" @click="createRoute(ref)">Создать</v-btn>
            </v-card-actions>
          </template>
        </RouteForm>

      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import Dialog from '@/views/base/components/Dialog.vue';
import _ from 'lodash';
import RouteForm from './RouteForm.vue';

@Component({
  components: {
    Dialog,
    RouteForm,
  },
})

export default class CreateRouteDialog extends Vue {
  show = false

  valid = true

  title = 'Создание маршрута'

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '/',
    tags: [],
  }

  disableCreateButton = true

  routeInitialState: Route = {
    description: '',
    id: -1,
    method: '',
    route: '/',
    tags: [],
  }

  @Watch('route', { deep: true })
  protected onChangeRoute(route: Route): void {
    this.disableCreateButton = this.routesSame(this.routeInitialState, route);
  }

  routesSame(a: Route, b: Route): boolean {
    return (a.description === b.description
      && a.method === b.method
      && a.route === b.route
     && _.isEqual(a, b));
  }

  async createRoute(ref: any) {
    if (ref) {
      ref.validate();
    }
    if (!this.route.description || !this.route.method || !this.route.route) {
      return;
    }

    await this.$store.direct.dispatch.route.Create(this.$data.route);
    this.clearForm();
    this.show = false;
  }

  clearForm() {
    this.route = {
      route: '',
      method: '',
      description: '',
      id: -1,
      tags: [],
    };
  }

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
