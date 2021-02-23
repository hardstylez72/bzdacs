<script lang="ts">
import {
  Component, Prop,
} from 'vue-property-decorator';
import { Route } from '@/views/route/entity';
import SimpleEntityDeleteDialog from '@/views/base/components/SimpleEntityDeleteDialog.vue';

@Component
export default class DeleteDialog extends SimpleEntityDeleteDialog<Route> {
  @Prop() readonly systemId!: number

  async deleteEntity(routeId: number) {
    const namespaceId = Number(this.$route.query.namespaceId);
    await this.$store.direct.dispatch.route.Delete({ namespaceId, id: routeId });
    this.$emit('change', false);
    this.$emit('routeDeleted');
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
