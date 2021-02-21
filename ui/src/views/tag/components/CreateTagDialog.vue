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

        <TagForm v-model="tag">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableCreateButton" @click="saveTag(ref)">{{$t('save')}}</v-btn>
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

import Dialog from '@/views/common/components/Dialog.vue';
import { Tag } from '@/views/tag/service';
import TagForm from './TagForm.vue';

@Component({
  components: {
    Dialog,
    TagForm,
  },
})

export default class CreateRouteDialog extends Vue {
  show = false

  valid = true

  title = 'Создание тега'

  tag: Tag = {
    name: '',
    id: -1,
  }

  disableCreateButton = true

  tagInitialState: Tag = {
    name: '',
    id: -1,
  }

  @Watch('tag', { deep: true })
  protected onChangeRoute(tag: Tag): void {
    this.disableCreateButton = this.tagsSame(this.tagInitialState, tag);
  }

  tagsSame(a: Tag, b: Tag): boolean {
    return a.name === b.name;
  }

  async saveTag(ref: any) {
    if (ref) {
      ref.validate();
    }
    if (!this.tag.name) {
      return;
    }

    await this.$store.direct.dispatch.tag.Create(this.tag);
    this.clearForm();
    this.show = false;
  }

  clearForm() {
    this.tag = {
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
    "add-btn": "Add tag",
    "title": "Tag creation",
    "cancel": "Cancel",
    "save": "Save"
  },
  "ru": {
    "add-btn": "Добавить тег",
    "title": "Создание тега",
    "cancel": "Отмена",
    "save": "Создать"
  }
}
</i18n>
