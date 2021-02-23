<template>
  <div>
    <v-btn
      color="primary"
      class="mb-2"
      @click="deleteSelectedRoutes"
    >
      Удалить выбранные маршруты
    </v-btn>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Route } from '@/views/route/entity';

@Component

export default class DeleteRoutesButton extends Vue {
  @Prop() items!: Route[]

  async deleteSelectedRoutes() {
    const routes = this.items;
    const groupId = Number(this.$route.params.id);
    const params = routes.map((route) => ({
      groupId,
      routeId: route.id,
    }));

    await this.$store.direct.dispatch.groupRoute.Delete(
      params,
    );
    this.$emit('routesDeleted');
  }
}
</script>

<style scoped lang="scss">

</style>
