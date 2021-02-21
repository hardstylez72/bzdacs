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

        <SimpleEntityForm v-model="local">
          <template v-slot:actions="{validate}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableCreateButton" @click="save(validate)">{{$t('save')}}</v-btn>
            </v-card-actions>
          </template>
        </SimpleEntityForm>

      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Vue, Watch,
} from 'vue-property-decorator';

import Dialog from '@/views/common/components/Dialog.vue';
import { SimpleEntity } from '@/views/base/services/entity';
import SimpleEntityForm from './SimpleEntityForm.vue';

@Component({
  components: {
    Dialog, SimpleEntityForm,
  },
})

export default class CreateNamespaceDialog<T extends SimpleEntity> extends Vue {
  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  @Watch('show')
  protected onChangeIsShowDialog(show: boolean): void {
    if (show !== this.value) {
      this.$emit('change', show);
    }
  }

  show = false

  valid = true

  local: T = {
    name: '',
    id: -1,
  }

  disableCreateButton = true

  propInitialState: T = {
    name: '',
    id: -1,
  }

  @Watch('local', { deep: true })
  protected localChanged(local: T): void {
    this.disableCreateButton = this.areSame(this.propInitialState, local);
  }

  areSame(a: T, b: T): boolean {
    return a.name === b.name;
  }

  async save(validate: any) {
    if (validate) {
      validate();
    }
    if (!this.local.name) {
      return;
    }

    await this.saveValid(this.local);

    this.clearForm();
    this.show = false;
  }

  async saveValid(entity: T) {
    // must be overwritten
  }

  clearForm() {
    this.local = {
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

<!--<i18n>-->
<!--{-->
<!--  "en": {-->
<!--    "add-btn": "Add namespace",-->
<!--    "title": "Namespace creation",-->
<!--    "cancel": "Cancel",-->
<!--    "save": "Save"-->
<!--  },-->
<!--  "ru": {-->
<!--    "add-btn": "Добавить пронстранство имен",-->
<!--    "title": "Создание пронстранства имен",-->
<!--    "cancel": "Отмена",-->
<!--    "save": "Создать"-->
<!--  }-->
<!--}-->
<!--</i18n>-->
