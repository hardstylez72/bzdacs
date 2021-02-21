<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('editing')}}
      </v-card-title>
      <v-card-text>
        <Form v-model="local">
          <template v-slot:actions="{validate}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" data-test="update-namespace-btn" text :disabled="disableUpdateButton" @click="update(validate)">{{$t('update')}}</v-btn>
            </v-card-actions>
          </template>
        </Form>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';

import Dialog from '@/views/common/components/Dialog.vue';
import SimpleEntityUpdateDialog from '@/views/base/components/SimpleEntityUpdateDialog.vue';
import { Namespace } from '@/views/namespace/service';
import Form from './Form.vue';

@Component({
  components: {
    Dialog, Form,
  },
})

export default class UpdateDialog extends SimpleEntityUpdateDialog<Namespace> {
  async updateValid(namespace: Namespace) {
    const ns = await this.$store.direct.dispatch.namespace.Update(namespace);
    this.$emit('updated', ns);
  }

  async loadItemWithId(id: number): Promise<Namespace> {
    return this.$store.direct.dispatch.namespace.GetById(id);
  }
}
</script>

<i18n>
{
  "en": {
    "editing": "Namespace name editing",
    "cancel": "Cancel",
    "update": "Update"
  },
  "ru": {
    "editing": "Редактирование названия сервиса",
    "cancel": "Отмена",
    "update": "Обновить"
  }
}
</i18n>
