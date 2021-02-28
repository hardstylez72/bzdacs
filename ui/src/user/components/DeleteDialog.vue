<template>
  <c-dialog v-model="show" max-width="450px">
    <v-card>
      <v-card-title>
        {{$t('sure')}}
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
        <v-btn color="blue darken-1" text @click="remove">{{$t('ok')}}</v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </c-dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';

@Component({
  components: {
    'c-dialog': () => import('../../common/components/Dialog.vue'),
  },
})
export default class DeleteDialog extends Vue {
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
    const namespaceId = Number(this.$route.query.namespaceId);
    await this.$store.direct.dispatch.user.Delete({ namespaceId, id: this.id });
    this.$emit('change', false);
    this.$emit('userDeleted');
  }
}
</script>

<i18n>
{
  "en": {
    "sure": "Are you sure want to delete that user?",
    "cancel": "Cancel",
    "ok": "OK"
  },
  "ru": {
    "sure": "Вы уверены что хотите удлить пользователя?",
    "cancel": "Отмена",
    "ok": "ОК"
  }
}
</i18n>
