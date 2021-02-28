<template>
  <Dialog v-model="show" max-width="450px">
    <v-card>
      <v-card-title>
        {{$t('sure')}}
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
        <v-btn color="blue darken-1" text @click="deleteGroup">{{$t('ok')}}</v-btn>
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

@Component({
  components: {
    Dialog,
    'c-dialog': () => import('../../common/components/Dialog.vue'),
  },
})
export default class DeleteRouteDialog extends Vue {
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

  async deleteGroup() {
    const namespaceId = Number(this.$route.query.namespaceId);
    await this.$store.direct.dispatch.group.Delete({ namespaceId, id: this.id });
    this.$emit('change', false);
    this.$emit('groupDeleted');
  }
}
</script>

<i18n>
{
  "en": {
    "sure": "Are you sure want to delete that group?",
    "cancel": "Cancel",
    "ok": "OK"
  },
  "ru": {
    "sure": "Вы уверены что хотите удлить группу?",
    "cancel": "Отмена",
    "ok": "ОК"
  }
}
</i18n>
