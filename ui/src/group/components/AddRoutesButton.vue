<template>
  <div>
    <v-btn
      color="primary"
      class="mb-2"
      @click="addSelectedRoutes"
    >
      {{$t('add-selected-routes')}}
    </v-btn>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Route } from '@/route/entity';

@Component

export default class AddRoutesButton extends Vue {
  @Prop() items!: Route[]

  async addSelectedRoutes() {
    const routes = this.items;
    const groupId = Number(this.$route.params.id);
    const params = routes.map((route) => ({
      groupId,
      routeId: route.id,
    }));

    await this.$store.direct.dispatch.groupRoute.Create(params);
    this.$emit('routesAdded');
  }
}
</script>

<style scoped lang="scss">

</style>
<i18n>
{
  "en": {
    "add-selected-routes": "Add selected routes"
  },
  "ru": {
    "add-selected-routes": "Добавить выбранные маршруты"
  }
}
</i18n>
