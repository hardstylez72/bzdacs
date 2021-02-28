<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('editing')}}
      </v-card-title>
      <v-card-text>
        <SystemForm v-model="local">
          <template v-slot:actions="{validate}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" data-test="update-system-btn" text :disabled="disableUpdateButton" @click="update(validate)">{{$t('update')}}</v-btn>
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

import Dialog from '@/common/components/Dialog.vue';
import SimpleEntityUpdateDialog from '@/base/components/SimpleEntityUpdateDialog.vue';
import { System } from '@/system/entity';
import SystemForm from './Form.vue';

@Component({
  components: {
    Dialog, SystemForm,
  },
})

export default class UpdateSystemDialog extends SimpleEntityUpdateDialog<System> {
  async updateValid(s: System) {
    const system = await this.$store.direct.dispatch.system.Update(s);
    this.$emit('itemUpdated', system);
  }

  async loadItemWithId(id: number): Promise<System> {
    return this.$store.direct.dispatch.system.GetById(id);
  }
}
</script>

<i18n>
{
  "en": {
    "editing": "System name editing",
    "cancel": "Cancel",
    "update": "Update"
  },
  "ru": {
    "editing": "Редактирование названия системы",
    "cancel": "Отмена",
    "update": "Обновить"
  }
}
</i18n>
