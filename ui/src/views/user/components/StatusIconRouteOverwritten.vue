<template>
  <v-tooltip v-if="item.isOverwritten && item.isIndependent" bottom>
    <template v-slot:activator="{ on, attrs }">
      <v-icon v-on="on" color="red">mdi-orbit-variant</v-icon>
    </template>
    <span>Маршрут перезатерт</span>
  </v-tooltip>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { RouteWithGroups, UpdateUserRoute } from '@/views/user/services/userroute';
import { User } from '@/views/user/services/user';

@Component
export default class StatusIconRouteOverwritten extends Vue {
  @Prop({ default: {} }) item!: RouteWithGroups
}

</script>

<style scoped lang="scss">

</style>
