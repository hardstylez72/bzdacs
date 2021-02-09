<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">{{$t('add-btn')}}</v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('title')}}
      </v-card-title>
      <v-card-text>

        <TagForm v-model="namespace">
          <template v-slot:actions="{validate}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableCreateButton" @click="saveNamespace(validate)">{{$t('save')}}</v-btn>
            </v-card-actions>
          </template>
        </TagForm>

      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Watch,
} from 'vue-property-decorator';

import Dialog from '@/views/base/components/Dialog.vue';
import { Tag } from '@/views/tag/service';
import { Namespace } from '@/views/namespace/service';
import NamespaceFormForm from './NamespaceForm.vue';

@Component({
  components: {
    Dialog,
    TagForm: NamespaceFormForm,
  },
})

export default class CreateNamespaceDialog extends Vue {
  show = false

  valid = true

  title = 'Создание тега'

  namespace: Namespace = {
    name: '',
    id: -1,
  }

  disableCreateButton = true

  namespaceInitialState: Namespace = {
    name: '',
    id: -1,
  }

  @Watch('namespace', { deep: true })
  protected onChangeRoute(namespace: Namespace): void {
    this.disableCreateButton = this.namespacesSame(this.namespaceInitialState, namespace);
  }

  namespacesSame(a: Namespace, b: Namespace): boolean {
    return a.name === b.name;
  }

  async saveNamespace(validate: any) {
    if (validate) {
      validate();
    }
    if (!this.namespace.name) {
      return;
    }

    await this.$store.direct.dispatch.namespace.Create(this.namespace);
    this.clearForm();
    this.show = false;
  }

  clearForm() {
    this.namespace = {
      name: '',
      id: -1,
    };
  }

  close() {
    this.show = false;
    this.clearForm();
  }
}
</script>

<i18n>
{
  "en": {
    "add-btn": "Add namespace",
    "title": "Namespace creation",
    "cancel": "Cancel",
    "save": "Save"
  },
  "ru": {
    "add-btn": "Добавить пронстранство имен",
    "title": "Создание пронстранства имен",
    "cancel": "Отмена",
    "save": "Создать"
  }
}
</i18n>
