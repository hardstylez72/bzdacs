<template>
  <div>
    <v-btn
      color="primary"
      class="mb-2"
      @click="deleteSelected"
    >
      {{$t('delete-selected-groups')}}
    </v-btn>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Group } from '@/views/group/entity';
import { Route } from '@/views/route/entity';

@Component

export default class DeleteRoutesButton extends Vue {
  @Prop() items!: Route[]

  userId = Number(this.$route.params.id);

  async deleteSelected() {
    const routes = this.items;
    const params = routes.map((route) => ({
      routeId: route.id,
      userId: this.userId,
    }));

    await this.$store.direct.dispatch.userRoute.Delete(
      params,
    );
    this.$emit('routesDeleted');
  }
}
</script>

<style scoped lang="scss">

</style>

<i18n>
{
  "en": {
    "delete-selected-groups": "Delete selected groups"
  },
  "ru": {
    "delete-selected-groups": "Удалить выбранные группы"
  }
}
</i18n>
