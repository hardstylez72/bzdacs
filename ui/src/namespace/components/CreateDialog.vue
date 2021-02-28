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
              <v-btn color="blue darken-1" text @click="close" >{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" data-test="save-namespace-btn" text :disabled="disableCreateButton" @click="save(validate)">{{$t('save')}}</v-btn>
            </v-card-actions>
          </template>
        </SystemForm>

      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Prop,
} from 'vue-property-decorator';

import Dialog from '@/common/components/Dialog.vue';
import SimpleEntityCreateDialog from '@/base/components/SimpleEntityCreateDialog.vue';
import { Namespace } from '@/namespace/entity';
import SystemForm from './Form.vue';

@Component({
  components: {
    Dialog, SystemForm,
  },
})

export default class CreateNamespaceDialog extends SimpleEntityCreateDialog<Namespace> {
  @Prop() systemId!: number

  async saveValid(namespace: Namespace) {
    const n = await this.$store.direct.dispatch.namespace.Create({ systemId: this.systemId, name: namespace.name });
    this.$emit('itemCreated', n, this.systemId);
  }
}
</script>

<i18n>
{
  "en": {
    "add-btn": "Add service",
    "title": "Service creation",
    "cancel": "Cancel",
    "save": "Save"
  },
  "ru": {
    "add-btn": "Добавить сервис",
    "title": "Создание сервиса",
    "cancel": "Отмена",
    "save": "Создать"
  }
}
</i18n>
