<template>
  <v-tooltip bottom>
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
import { RouteWithGroups, UpdateUserRoute } from '@/user/services/userroute';
import { User } from '@/user/entity';
import { userRoutePair } from '@/user/generated';

@Component
export default class UpdateRouteButton extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups

  @Prop() userId!: number

  loading = false

  async updateRoute() {
    const param: userRoutePair = {
      isExcluded: !this.item.isExcluded,
      routeId: this.item.id,
      userId: this.userId,
    };

    try {
      this.loading = true;
      await this.$store.direct.dispatch.userRoute.UpdateRoute(param);
      this.$emit('routeUpdated');
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
