<template>
  <v-tooltip bottom>
    <template v-slot:activator="{ on, attrs }">
      <div v-on="on">
        <v-icon v-if="item.isExcluded" color="red">mdi-account-off</v-icon>
        <v-icon v-else color="green">mdi-account</v-icon>
      </div>
    </template>
    <span v-if="item.isExcluded">Маршрут запрещен</span>
    <span v-else>Маршрут разрешен</span>
  </v-tooltip>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { RouteWithGroups, UpdateUserRoute } from '@/views/user/services/userroute';
import { User } from '@/views/user/services/user';

@Component
export default class StatusIconRouteAccess extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups
}

</script>

<style scoped lang="scss">

</style>
