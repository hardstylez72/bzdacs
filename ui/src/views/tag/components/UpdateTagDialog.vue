<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('editing')}}
      </v-card-title>
      <v-card-text>
        <TagForm v-model="tag">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableUpdateButton" @click="updateTag(ref)">{{$t('update')}}</v-btn>
            </v-card-actions>
          </template>
        </TagForm>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import Dialog from '@/views/base/components/Dialog.vue';
import { Tag } from '@/views/tag/service';
import TagForm from './TagForm.vue';

@Component({
  components: {
    Dialog,
    TagForm,
  },
})
export default class UpdateTagDialog extends Vue {
  show = false

  disableUpdateButton = true

  valid = true

  @Prop({ required: true }) id!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  tag: Tag = {
    name: '',
    id: -1,
  }

   tagInitialState: Tag | undefined

  @Watch('id')
  protected onChangeItem(tagId: number): void {
    this.$store.direct.dispatch.tag.GetById(tagId).then((tag) => {
      this.tag = {
        name: tag.name,
        id: tag.id,
      };
      this.tagInitialState = {
        name: tag.name,
        id: tag.id,
      };
    });
  }

  @Watch('tag', { deep: true })
  protected onChangeRoute(tag: Tag): void {
    if (!tag.name) {
      this.disableUpdateButton = true;
      return;
    }
    if (this.tagInitialState) {
      if (this.routesSame(this.tagInitialState, tag)) {
        this.disableUpdateButton = true;
      } else {
        this.disableUpdateButton = false;
      }
    }
  }

  async updateTag(ref: any) {
    if (ref) {
      ref.validate(); // vee-validate specifics
    }

    if (!this.tag.id || !this.tag.name) {
      return;
    }

    await this.$store.direct.dispatch.tag.Update(this.tag);
    this.$emit('change', false);
    this.disableUpdateButton = true;
    this.$store.direct.dispatch.tag.GetById(this.id).then((tag) => {
      this.tag = {
        name: tag.name,
        id: tag.id,
      };
      this.tagInitialState = {
        name: tag.name,
        id: tag.id,
      };
    });

    this.show = false;
  }

  routesSame(a: Tag, b: Tag): boolean {
    return (a.name === b.name);
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }
}
</script>
<i18n>
{
  "en": {
    "editing": "Tag editing",
    "cancel": "Cancel",
    "update": "Update"
  },
  "ru": {
    "editing": "Редактирование тега",
    "cancel": "Отмена",
    "update": "Обновить"
  }
}
</i18n>
