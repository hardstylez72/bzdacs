<template>
  <v-tooltip bottom v-if="!item.isIndependent">
    <template v-slot:activator="{ on, attrs }">
      <div @click="createRoute" v-on="on" style="cursor: pointer">
        <v-icon>mdi-plus</v-icon>
      </div>
    </template>
      <span>Перезатереть маршрут</span>
  </v-tooltip>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { CreateUserRoute, RouteWithGroups, UpdateUserRoute } from '@/views/user/services/userroute';
import { User } from '@/views/user/services/user';

@Component
export default class OverwriteRouteButton extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups

  @Prop({ default: {} }) user!: User

  async createRoute() {
    const params: CreateUserRoute = {
      userId: this.user.id,
      routeId: this.item.id,
      isExcluded: true,
    };

    await this.$store.direct.dispatch.userRoute.RewriteOne(params);
}
}

</script>

<style scoped lang="scss">

</style>
