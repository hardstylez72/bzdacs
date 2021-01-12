<template>
  <c-dialog
    v-model="show"
    max-width="450px"
  >
    <v-card>
      <v-card-title>
        {{$t('sure')}}
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
        <v-btn color="blue darken-1" text @click="deleteRoute">{{$t('ok')}}</v-btn>
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
    'c-dialog': () => import('../../base/components/Dialog.vue'),
  },
})
export default class DeleteRouteDialog extends Vue {
  show = false

  @Prop() routeId!: number

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

  async deleteRoute() {
    await this.$store.direct.dispatch.route.Delete(this.routeId);
    this.$emit('change', false);
  }
}
</script>

<i18n>
{
  "en": {
    "sure": "Are you sure want to delete that route?",
    "cancel": "Cancel",
    "ok": "OK"
  },
  "ru": {
    "sure": "Вы уверены что хотите удлить маршрут?",
    "cancel": "Отмена",
    "ok": "ОК"
  }
}
</i18n>
