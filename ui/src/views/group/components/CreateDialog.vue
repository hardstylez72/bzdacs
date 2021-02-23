<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">{{$t('add-group-btn')}}</v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">{{$t('add-group-title')}}</v-card-title>
      <v-card-text>
        <v-form ref="create-group-form" v-model="valid" lazy-validation>
          <v-row>
            <v-col cols="12" sm="4" md="4">
              <v-text-field
                v-model="group.code"
                required
                :rules="codeRules"
                :label="$t('label.code')"
              />
            </v-col>
            <v-col cols="12" sm="4" md="4">
              <v-select
                :items="getGroups"
                item-text="code"
                item-value="id"
                v-model="baseGroupId"
                :label="$t('label.base-code-group')"
              />
            </v-col>
            <v-col cols="12" sm="10" md="10">
              <v-textarea
                v-model="group.description"
                outlined
                required
                :rules="descriptionRules"
                :label="$t('label.desc')"
              />
            </v-col>
          </v-row>
        </v-form>

        <v-card-actions>
          <v-spacer />
          <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
          <v-btn color="blue darken-1" text @click="createGroup">{{$t('save')}}</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/views/group/entity';
import Dialog from '@/views/common/components/Dialog.vue';
import { ListResponse } from '@/views/common/helpers/types';

@Component({
  components: {
    Dialog,
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

  validate() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs['create-group-form'].validate();
  }

  rules = [
    (v: string) => !!v || this.$t('required'),
]

  codeRules = this.rules

  descriptionRules = this.rules

  groups: Group[] = []

  get getGroups(): readonly Group[] {
    return this.groups;
  }

  async created() {
    const res = await this.loadItems();
    this.groups = res.items;
  }

  async loadItems(): Promise<ListResponse<Group>> {
    const namespaceId = Number(this.$route.query.namespaceId);
    return this.$store.direct.dispatch.group.GetList({
      filter: {
        namespaceId,
        pageSize: 100,
        page: 1,
      },
    });
  }

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
    const namespaceId = Number(this.$route.query.namespaceId);
    await this.$store.direct.dispatch.group.Create({
      namespaceId,
      description: this.group.description,
      baseGroupId: this.baseGroupId,
      code: this.group.code,
    });
    this.$emit('groupCreated');

    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<i18n>
{
  "en": {
    "add-group-btn": "Add group",
    "add-group-title": "Group creation",
    "label": {
      "code": "Code",
      "base-code-group": "Base group code",
      "desc": "Description"
    },
    "cancel": "Cancel",
    "save": "Save",
    "required": "required"
  },
  "ru": {
    "add-group-btn": "Добавить группу",
    "add-group-title": "Создание группы",
    "label": {
      "code": "Код",
      "base-code-group": "Код базовой группы",
      "desc": "Описание"
    },
    "cancel": "Отмена",
    "save": "Создать",
    "required": "Обязательное поле"
  }
}
</i18n>
