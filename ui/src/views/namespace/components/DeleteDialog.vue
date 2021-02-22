<script lang="ts">
import {
  Component, Prop,
} from 'vue-property-decorator';
import { System } from '@/views/system/entity';
import SimpleEntityDeleteDialog from '@/views/base/components/SimpleEntityDeleteDialog.vue';

@Component
export default class DeleteDialog extends SimpleEntityDeleteDialog<System> {
  @Prop() readonly systemId!: number

  async deleteEntity(namespaceId: number) {
    // must be overwritten!
    await this.$store.direct.dispatch.namespace.Delete({ namespaceId });
    this.$emit('deleted', namespaceId, this.systemId);
  }
}
</script>

<i18n>
{
  "en": {
    "sure": "Are you sure want to delete that namespace?",
    "cancel": "Cancel",
    "ok": "OK"
  },
  "ru": {
    "sure": "Вы уверены что хотите удлить сервис?",
    "cancel": "Отмена",
    "ok": "ОК"
  }
}
</i18n>
