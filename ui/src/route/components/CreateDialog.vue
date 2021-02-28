<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn
        data-test="add-route-btn"
        color="primary"
        class="mb-2"
        v-bind="props"
        v-on="props.on"
      >
        {{$t('add-btn')}}
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('add-title')}}
      </v-card-title>
      <v-card-text>

        <RouteForm v-model="route">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                {{$t('cancel')}}
              </v-btn>
              <v-btn
                color="blue darken-1"
                text
                :disabled="disableCreateButton"
                @click="createRoute(ref)"
                data-test="save-route-btn"
              >
                {{$t('save')}}
              </v-btn>
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
import { Route } from '@/route/entity';
import Dialog from '@/common/components/Dialog.vue';
import _ from 'lodash';
import RouteForm from './Form.vue';

@Component({
  components: {
    Dialog,
    RouteForm,
  },
})

export default class CreateRouteDialog extends Vue {
  show = false

  valid = true

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '',
    tags: [],
  }

  disableCreateButton = true

  routeInitialState: Route = this.route

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
    const namespaceId = Number(this.$route.query.namespaceId);
    await this.$store.direct.dispatch.route.Create({
      namespaceId,
      tags: this.route.tags,
      route: this.route.route,
      description: this.route.description,
      method: this.route.method,
    });
    this.$emit('routeCreated');
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

<i18n>
{
  "en": {
    "add-btn": "Add route",
    "add-title": "Route creation",
    "cancel": "Cancel",
    "save": "Save"
  },
  "ru": {
    "add-btn": "Добавить маршрут",
    "add-title": "Создание маршрута",
    "cancel": "Отмена",
    "save": "Создать"
  }
}
</i18n>
