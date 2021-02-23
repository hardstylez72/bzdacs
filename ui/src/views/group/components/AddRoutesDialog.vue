<template>
  <Dialog
    v-model="show"
    max-width="2000px"
  >
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">{{$t('add-routes-btn')}}</v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('add-routes-title')}}
      </v-card-title>

      <v-card-text>
        <RoutesNotBelongGroupTable :group-id="groupId" @routesAdded="routesAdded"/>
      </v-card-text>

      <v-card-actions>
        <v-spacer />
        <v-btn color="blue darken-1" text @click="close">{{$t('cancel-btn')}}</v-btn>
        <v-spacer />
      </v-card-actions>

    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import Dialog from '@/views/common/components/Dialog.vue';
import RoutesNotBelongGroupTable from './RoutesNotBelongGroupTable.vue';

@Component({
  components: {
    Dialog,
    RoutesNotBelongGroupTable,
  },
})
export default class AddRoutesDialog extends Vue {
  show = false

  @Prop({ type: Number, default: -1 })
  private readonly groupId!: number

  routesAdded() {
    this.$emit('routesAdded');
    this.close();
  }

  close() {
    this.show = false;
  }
}
</script>

<i18n>
{
  "en": {
    "add-routes-btn": "Add routes",
    "add-routes-title": "Adding routes to the group",
    "add-selected-routes": "Add selected routes",
    "cancel-btn": "Cancel"
  },
  "ru": {
    "add-routes-btn": "Добавить маршруты",
    "add-routes-title": "Добавление маршрутов к группе",
    "add-selected-routes": "Добавить выбранные маршруты",
    "cancel-btn": "Отмена"
  }
}
</i18n>
