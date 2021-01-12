<template>
  <Dialog v-model="show" max-width="450px">
    <v-card>
      <v-card-title>
        {{$t('sure')}}
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
        <v-btn color="blue darken-1" text @click="deleteTag">{{$t('ok')}}</v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import Dialog from '@/views/base/components/Dialog.vue';

@Component({
  components: {
    Dialog,
  },
})
export default class DeleteTagDialog extends Vue {
  show = false

  @Prop() tagId!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }

  async deleteTag() {
    await this.$store.direct.dispatch.tag.Delete(this.tagId);
    this.$emit('change', false);
  }
}
</script>

<i18n>
{
  "en": {
    "sure": "Are you sure want to delete that tag?",
    "cancel": "Cancel",
    "ok": "OK"
  },
  "ru": {
    "sure": "Вы уверены что хотите удлить тег?",
    "cancel": "Отмена",
    "ok": "ОК"
  }
}
</i18n>
