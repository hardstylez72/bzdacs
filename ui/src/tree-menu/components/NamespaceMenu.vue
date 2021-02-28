<template>
  <div>
    <v-menu  right>
      <template v-slot:activator="{ on, attrs }">
        <v-btn dark icon v-bind="attrs" v-on="on">
          <v-icon color="grey">mdi-dots-vertical</v-icon>
        </v-btn>
      </template>

      <v-list>
        <v-list-item data-test="namespace-edit-option" @click="clickedUpdateButton">
          <v-list-item-title>{{this.$t('edit')}}</v-list-item-title>
          <v-list-item-icon>
            <v-icon>mdi-pencil</v-icon>
          </v-list-item-icon>
        </v-list-item>
        <v-list-item data-test="namespace-delete-option" @click="clickedDeleteButton">
          <v-list-item-title>{{this.$t('delete')}}</v-list-item-title>
          <v-list-item-icon>
            <v-icon>mdi-delete</v-icon>
          </v-list-item-icon>
        </v-list-item>
      </v-list>
    </v-menu>
    <DeleteDialog v-model="showDeleteDialog" :system-id="system.id" :id="namespace.id" @deleted="namespaceDeleted"/>
    <UpdateDialog v-model="showUpdateDialog" :id="namespace.id" @updated="namespaceUpdated"/>
  </div>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { Namespace } from '@/namespace/entity';
import { System } from '@/system/entity';
import DeleteDialog from '../../namespace/components/DeleteDialog.vue';
import UpdateDialog from '../../namespace/components/UpdateDialog.vue';

@Component({
  components: {
    DeleteDialog,
    UpdateDialog,
  },
})

export default class NamespaceMenu extends Vue {
  @Prop() namespace!: Namespace

  @Prop() system!: System

  showUpdateDialog = false

  showDeleteDialog = false

  clickedDeleteButton() {
    this.showDeleteDialog = true;
  }

  clickedUpdateButton() {
    this.showUpdateDialog = true;
  }

  namespaceUpdated(namespace: Namespace) {
    this.$emit('updated', namespace, this.system);
  }

  namespaceDeleted(namespaceId: number, systemId: number) {
    this.$emit('deleted', namespaceId, this.system.id);
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
