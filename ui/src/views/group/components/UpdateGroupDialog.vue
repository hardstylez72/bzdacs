<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">Редактирование группы</v-card-title>
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
          <v-btn color="blue darken-1" text @click="updateGroup">Save</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { Group } from '@/views/group/services/group';
import { Route } from '@/views/route/service';

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
export default class UpdateGroupDialog extends Vue {
  @Prop({ required: true }) id!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  @Watch('id')
  protected onChangeItem(groupId: number): void {
    this.group.id = groupId;
  }

  @Watch('activeGroup')
  protected onChangeActiveGroup(g: Group): void {
    this.group.id = g.id;
    this.group.code = g.code;
    this.group.description = g.description;
  }

  show = false

  valid = false

  group: Group = {
    description: '',
    id: 0,
    code: '',
  }

  get activeGroup(): Group {
    return this.$store.direct.getters.group.getEntities.filter(((route) => route.id === this.id))[0];
  }

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

  async updateGroup() {
    this.validate();
    if (!this.valid) {
      return;
    }

    if (!this.group.code || !this.group.description || !this.group.id) {
      return;
    }

    await this.$store.direct.dispatch.group.Update(this.group);
    this.$emit('change', false);
    this.show = false;
  }

  close() {
    this.show = false;
    this.$emit('change', false);
  }
}
</script>

<style scoped lang="scss">

</style>
