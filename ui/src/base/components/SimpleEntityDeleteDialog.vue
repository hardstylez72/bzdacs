<template>
  <Dialog v-model="show" max-width="450px">
    <v-card>
      <v-card-title>
        {{$t('sure')}}
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
        <v-btn data-test="delete-btn" color="blue darken-1" text @click="remove">{{$t('ok')}}</v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import Dialog from '@/common/components/Dialog.vue';
import { SimpleEntity } from '@/base/entity';

@Component({
  components: {
    Dialog,
  },
})
export default class SimpleEntityDeleteDialog<T extends SimpleEntity> extends Vue {
  show = false

  @Prop() id!: number

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

  async remove() {
    await this.deleteEntity(this.id);
    this.$emit('change', false);
  }

  async deleteEntity(id: number) {
    // must be overwritten!
  }
}
</script>

<!--<i18n>-->
<!--{-->
<!--  "en": {-->
<!--    "sure": "Are you sure want to delete that tag?",-->
<!--    "cancel": "Cancel",-->
<!--    "ok": "OK"-->
<!--  },-->
<!--  "ru": {-->
<!--    "sure": "Вы уверены что хотите удлить тег?",-->
<!--    "cancel": "Отмена",-->
<!--    "ok": "ОК"-->
<!--  }-->
<!--}-->
<!--</i18n>-->
