<template>
  <v-tooltip bottom>
    <template v-slot:activator="{ on, attrs }">
      <div v-on="on">
        <v-icon v-if="item.isExcluded" color="red">mdi-account-off</v-icon>
        <v-icon v-else color="green">mdi-account</v-icon>
      </div>
    </template>
    <span v-if="item.isExcluded">{{$t('forbidden')}}</span>
    <span v-else>{{$t('allowed')}}</span>
  </v-tooltip>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { RouteWithGroups } from '@/user/services/userroute';

@Component
export default class StatusIconRouteAccess extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups
}

</script>

<i18n>
{
  "en": {
    "forbidden": "Route is forbidden",
    "allowed": "Route is allowed"
  },
  "ru": {
    "forbidden": "Маршрут запрещен",
    "allowed": "Маршрут разрешен"
  }
}
</i18n>
