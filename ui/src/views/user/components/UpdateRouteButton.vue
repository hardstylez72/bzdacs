<template>
  <v-tooltip bottom v-if="item.isIndependent && !item.isOverwritten">
    <template v-slot:activator="{ on, attrs }">
      <div @click="updateRoute" v-on="on" style="cursor: pointer">
        <v-icon v-if="!item.isExcluded" >mdi-account-minus</v-icon>
        <v-icon v-else >mdi-account-plus</v-icon>
      </div>
    </template>
      <span v-if="!item.isExcluded">{{$t('forbid')}}</span>
      <span v-else>{{$t('unlock')}}</span>
  </v-tooltip>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { RouteWithGroups, UpdateUserRoute } from '@/views/user/services/userroute';
import { User } from '@/views/user/entity';

@Component
export default class UpdateRouteButton extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups

  @Prop({ default: {} }) user!: User

  loading = false

  async updateRoute() {
    const param: UpdateUserRoute = {
      isExcluded: !this.item.isExcluded,
      routeId: this.item.id,
      userId: this.user.id,
    };

    try {
      this.loading = true;
      await this.$store.direct.dispatch.userRoute.UpdateRoute(param);
    } finally {
      this.loading = false;
    }
  }
}

</script>

<i18n>
{
  "en": {
    "unlock": "Unlock route",
    "forbid": "Forbid"
  },
  "ru": {
    "unlock": "Разблокировать маршрут",
    "forbid": "Заблокировать маршрут"
  }
}
</i18n>
