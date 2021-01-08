<template>
  <v-tooltip bottom v-if="item.isIndependent">
    <template v-slot:activator="{ on, attrs }">
      <div @click="deleteRoute" v-on="on" style="cursor: pointer">
        <v-icon>mdi-close</v-icon>
      </div>
    </template>
    <span v-if="item.isIndependent && item.isOverwritten">Отменить перезатирание маршрута</span>
    <span v-else>Удалить маршрут</span>
  </v-tooltip>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import {
 CreateUserRoute, RouteWithGroups, UpdateUserRoute, UserRoute,
} from '@/views/user/services/userroute';
import { User } from '@/views/user/services/user';

@Component
export default class OverwriteRouteButton extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups

  @Prop({ default: {} }) user!: User

  async deleteRoute() {
    const params: UserRoute = {
      userId: this.user.id,
      routeId: this.item.id,
    };

    await this.$store.direct.dispatch.userRoute.DeleteOne(params);
  }
}

</script>

<style scoped lang="scss">

</style>
