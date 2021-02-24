<template>
  <div>
    <v-btn
      color="primary"
      class="mb-2"
      @click="addSelectedRoutes"
    >
      {{$t('add-selected-routes')}}
    </v-btn>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Group } from '@/views/group/entity';

@Component

export default class AddGroupsButton extends Vue {
  @Prop() items!: Group[]

  async addSelectedRoutes() {
    const routes = this.items;
    const userId = Number(this.$route.params.id);
    const params = routes.map((group) => ({
      userId,
      groupId: group.id,
    }));

    await this.$store.direct.dispatch.userGroup.Create(params);
    this.$emit('groupsAdded');
  }
}
</script>

<style scoped lang="scss">

</style>
<i18n>
{
  "en": {
    "add-selected-routes": "Add selected routes"
  },
  "ru": {
    "add-selected-routes": "Добавить выбранные маршруты"
  }
}
</i18n>
