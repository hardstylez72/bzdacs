<template>
  <div>
    <v-btn
      color="primary"
      class="mb-2"
      @click="addSelectedRoutes"
    >
      {{$t('add-selected')}}
    </v-btn>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Group } from '@/group/entity';
import { Route } from '@/route/entity';

@Component

export default class AddRoutesButton extends Vue {
  @Prop() items!: Route[]

  userId = Number(this.$route.params.id);

  async addSelectedRoutes() {
    const routes = this.items;
    const params = routes.map((route) => ({
      userId: this.userId,
      routeId: route.id,
    }));

    await this.$store.direct.dispatch.userRoute.Create(params);
    this.$emit('routesAdded');
  }
}
</script>

<style scoped lang="scss">

</style>
<i18n>
{
  "en": {
    "add-selected": "Add selected routes"
  },
  "ru": {
    "add-selected": "Добавить выбранные маршруты"
  }
}
</i18n>
