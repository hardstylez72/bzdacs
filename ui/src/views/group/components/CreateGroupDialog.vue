<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">Новая группа</v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">Создание группы</v-card-title>
      <v-card-text>
        <v-form ref="create-group-form" v-model="valid" lazy-validation>
          <v-row>
            <v-col cols="12" sm="4" md="4">
              <v-text-field
                v-model="group.code"
                required
                :rules="codeRules"
                label="Код"
              />
            </v-col>
            <v-col cols="12" sm="4" md="4">
              <v-select
                :items="groups"
                item-text="code"
                item-value="id"
                v-model="baseGroupId"
                label="Код базовой группы"
              />
            </v-col>
            <v-col cols="12" sm="10" md="10">
              <v-textarea
                v-model="group.description"
                outlined
                required
                :rules="descriptionRules"
                label="Описание"
              />
            </v-col>
          </v-row>
        </v-form>

        <v-card-actions>
          <v-spacer />
          <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
          <v-btn color="blue darken-1" text @click="createGroup">Save</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/views/group/services/group';

@Component({
  components: {
    Dialog: () => import('../../base/components/Dialog.vue'),
  },
  watch: {
    match: 'validate',
    max: 'validate',
    model: 'validate',
  },

})
export default class CreateRouteDialog extends Vue {
  show = false

  valid = false

  group: Group = {
    description: '',
    id: 0,
    code: '',
  }

  baseGroupId = 0

  get groups(): readonly Group[] {
    return this.$store.direct.getters.group.getEntities;
  }

    validate() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs['create-group-form'].validate();
  }

  rules = [
    (v: string) => !!v || 'Обязательное поле',
]

  codeRules = this.rules

  descriptionRules = this.rules

  async createGroup() {
    this.validate();
    if (!this.valid) {
      return;
    }

    if (!this.group.code) {
      return;
    }
    if (!this.group.description) {
      return;
    }

    if (this.baseGroupId === 0) {
      await this.$store.direct.dispatch.group.Create(this.group);
    } else {
      await this.$store.direct.dispatch.group.CreateBasedOnAnother({
        group: this.group,
        baseGroupId: this.baseGroupId,
      });
    }

    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
