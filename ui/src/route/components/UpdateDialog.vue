<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('editing')}}
      </v-card-title>
      <v-card-text>
        <RouteForm v-model="route">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close"> {{$t('cancel')}}</v-btn>
              <v-btn data-test="update-route-btn" color="blue darken-1" text :disabled="disableUpdateButton" @click="updateRoute(ref)"> {{$t('update')}}</v-btn>
            </v-card-actions>
          </template>
        </RouteForm>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
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
export default class UpdateRouteDialog extends Vue {
  show = false

  disableUpdateButton = true

  valid = true

  @Prop() routeId!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '/',
    tags: [],
  }

   initialRouteState: Route = this.route

  @Watch('routeId')
  protected onChangeItem(id: number): void {
    const namespaceId = Number(this.$route.query.namespaceId);
    this.$store.direct.dispatch.route.GetById({ namespaceId, id }).then((item) => {
      this.route = {
        description: item.description,
        id: item.id,
        method: item.method,
        route: item.route,
        tags: item.tags,
      };

      this.initialRouteState = {
        description: item.description,
        id: item.id,
        method: item.method,
        route: item.route,
        tags: item.tags,
      };
    });
  }

  @Watch('route', { deep: true })
  protected onChangeRoute(route: Route): void {
    if (this.routesSame(this.initialRouteState, route)) {
      this.disableUpdateButton = true;
      return;
    }

    if (this.routesSame(this.route, route)) {
      this.disableUpdateButton = false;
    }
  }

  async updateRoute(ref: any) {
    if (ref) {
      ref.validate(); // vee-validate specifics
    }

    if (!this.route.id || !this.route.description || !this.route.method || !this.route.route) {
      return;
    }

    const namespaceId = Number(this.$route.query.namespaceId);
    await this.$store.direct.dispatch.route.Update({
      namespaceId,
      id: this.route.id,
      tags: this.route.tags,
      method: this.route.method,
      description: this.route.description,
      route: this.route.route,
    });
    this.$emit('change', false);
    this.$emit('routeUpdated');
    this.show = false;
  }

  routesSame(a: Route, b: Route): boolean {
    return (a.description === b.description
      && a.method === b.method
      && a.route === b.route
     && _.isEqual(a.tags, b.tags));
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }
}
</script>

<i18n>
{
  "en": {
    "editing": "Route editing",
    "cancel": "Cancel",
    "update": "Update"
  },
  "ru": {
    "editing": "Редактирование маршрута",
    "cancel": "Отмена",
    "update": "Обновить"
  }
}
</i18n>
