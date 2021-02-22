<template>
  <div>
    <v-menu  right>
      <template v-slot:activator="{ on, attrs }">
        <v-btn dark icon v-bind="attrs" v-on="on">
          <v-icon color="grey">mdi-dots-vertical</v-icon>
        </v-btn>
      </template>

      <v-list>
        <v-list-item data-test="system-edit-option" @click="clickedUpdateSystemButton">
          <v-list-item-title>{{this.$t('edit')}}</v-list-item-title>
          <v-list-item-icon>
            <v-icon>mdi-pencil</v-icon>
          </v-list-item-icon>
        </v-list-item>
        <v-list-item data-test="system-delete-option" @click="clickedDeleteSystemButton">
          <v-list-item-title>{{this.$t('delete')}}</v-list-item-title>
          <v-list-item-icon>
            <v-icon>mdi-delete</v-icon>
          </v-list-item-icon>
        </v-list-item>
      </v-list>
    </v-menu>
    <DeleteSystemDialog v-model="showDeleteSystemDialog" :id="system.id" @itemDeleted="systemDeleted"/>
    <UpdateSystemDialog v-model="showUpdateSystemDialog" :id="system.id" @itemUpdated="systemUpdated"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { System } from '@/views/system/entity';
import DeleteSystemDialog from '../../system/components/DeleteDialog.vue';
import UpdateSystemDialog from '../../system/components/UpdateDialog.vue';

@Component({
  components: {
    DeleteSystemDialog,
    UpdateSystemDialog,
  },
})

export default class SystemMenu extends Vue {
  @Prop() system!: System

  showUpdateSystemDialog = false

  showDeleteSystemDialog = false

  clickedDeleteSystemButton() {
    this.showDeleteSystemDialog = true;
  }

  clickedUpdateSystemButton() {
    this.showUpdateSystemDialog = true;
  }

  systemUpdated(system: System) {
    this.$emit('systemUpdated');
  }

  systemDeleted(id: number) {
    this.$emit('systemDeleted');
  }
}
</script>

<style scoped lang="scss">

</style>

<i18n>
{
  "en": {
    "edit": "Edit",
    "delete": "Delete"
  },
  "ru": {
    "edit": "Редактировать",
    "delete": "Удалить"
  }
}
</i18n>
