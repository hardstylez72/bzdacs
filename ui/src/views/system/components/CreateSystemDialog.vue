<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('title')}}
      </v-card-title>
      <v-card-text>

        <SystemForm v-model="local">
          <template v-slot:actions="{validate}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableCreateButton" @click="save(validate)">{{$t('save')}}</v-btn>
            </v-card-actions>
          </template>
        </SystemForm>

      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';

import Dialog from '@/views/base/components/Dialog.vue';
import SimpleEntityCreateDialog from '@/views/base/components/SimpleEntityCreateDialog.vue';
import { System } from '@/views/tree-menu/entity';
import SystemForm from './SystemForm.vue';

@Component({
  components: {
    Dialog, SystemForm,
  },
})

export default class CreateSystemDialog extends SimpleEntityCreateDialog<System> {
  async saveValid(s: System) {
    const system = await this.$store.direct.dispatch.system.Create(s);
    this.$emit('itemCreated', system);
  }
}
</script>

<i18n>
{
  "en": {
    "add-btn": "Add system",
    "title": "System creation",
    "cancel": "Cancel",
    "save": "Save"
  },
  "ru": {
    "add-btn": "Добавить систему",
    "title": "Создание системы",
    "cancel": "Отмена",
    "save": "Создать"
  }
}
</i18n>
