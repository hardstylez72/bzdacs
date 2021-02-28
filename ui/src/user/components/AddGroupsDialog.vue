<template>
  <Dialog v-model="show" max-width="1700px">
    <template v-slot:activator="props">
      <v-btn
        color="primary"
        class="mb-2"
        v-bind="props"
        v-on="props.on"
      >
        {{$t('add-btn')}}
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('title')}}
      </v-card-title>

      <v-card-text>
        <GroupsNotBelongUserTable :user-id="userId" @groupsAdded="groupsAdded"/>
      </v-card-text>

      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';

import Dialog from '@/common/components/Dialog.vue';
import GroupsNotBelongUserTable from './GroupsNotBelongUserTable.vue';

@Component({
  components: {
    Dialog,
    GroupsNotBelongUserTable,
  },
})
export default class AddGroupsDialog extends Vue {
  show = false

  @Prop({ type: Number, default: -1 })
  private readonly userId!: number

  groupsAdded() {
    this.close();
    this.$emit('groupsAdded');
  }

  close() {
    this.show = false;
  }
}
</script>

<i18n>
{
  "en": {
    "add-btn": "Add groups",
    "title": "Adding group to the user",
    "add-selected": "Add selected groups",
    "cancel": "Cancel"
  },
  "ru": {
    "add-btn": "Добавить группы",
    "title": "Добавление групп к пользователю",
    "add-selected": "Добавить выбранные группы",
    "cancel": "Отмена"
  }
}
</i18n>
