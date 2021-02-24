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

@Component

export default class DeleteGroupsButton extends Vue {
  @Prop() items!: Group[]

  async deleteSelected() {
    const groups = this.items;
    const userId = Number(this.$route.params.id);
    const params = groups.map((group) => ({
      groupId: group.id,
      userId,
    }));

    await this.$store.direct.dispatch.userGroup.Delete(
      params,
    );
    this.$emit('groupsDeleted');
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
