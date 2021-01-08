<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">Новый пользователь</v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">Создание пользователя</v-card-title>
      <v-card-text>
        <v-form ref="create-user-form" v-model="valid" lazy-validation>
          <v-row>
            <v-col cols="12" sm="4" md="4">
              <v-text-field
                v-model="user.externalId"
                required
                :rules="externalIdRules"
                label="Идентфикатор"
              />
            </v-col>
          </v-row>
        </v-form>

        <v-card-actions>
          <v-spacer />
          <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
          <v-btn color="blue darken-1" text @click="create">Save</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { User } from '@/views/user/services/user';

@Component({
  components: {
    Dialog: () => import('../../base/components/Dialog.vue'),
  },
})
export default class CreateRouteDialog extends Vue {
  show = false

  valid = false

  user: User = {
    description: '',
    id: -1,
    externalId: '',
    isSystem: false,
  }

  validate() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs['create-user-form'].validate();
  }

  rules = [
    (v: string) => !!v || 'Обязательное поле',
]

  externalIdRules = this.rules

  async create() {
    this.validate();
    if (!this.valid) {
      return;
    }

    if (!this.user.externalId) {
      return;
    }

    await this.$store.direct.dispatch.user.Create(this.user);
    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
