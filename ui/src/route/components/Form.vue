<template>
  <div>
    <v-form ref="refss" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="10" md="10">
          <v-text-field
            v-model="route.route"
            required
            :rules="routeRules"
            :label="$t('label.route')"
            data-test="form-route-input"
          />
        </v-col>
        <v-col cols="2" sm="2" md="2" data-test="form-method-input">
          <v-select
            v-model="route.method"
            required
            :rules="methodRules"
            :items="httpMethodList"
            :label="$t('label.method')"
          />
        </v-col>
        <v-col cols="12" sm="12" md="12">
          <v-textarea
            v-model="route.description"
            outlined
            required
            :rules="descriptionRules"
            :label="$t('label.desc')"
            data-test="form-description-input"
          />
        </v-col>
        <v-col cols="12" sm="12" md="12">
          <v-autocomplete
            :label="$t('label.tags')"
            :placeholder="$t('tag.tip')"
            ref="autocomplete-input"
            v-model="selectedTags"
            :items="suggestedTags"
            :search-input.sync="searchTags"
            item-text="name"
            item-value="name"
            hide-selected
            :disabled="isSuggestUpdating"
            multiple
            chips
          >
            <template v-slot:no-data>
              <v-list-item>
                <v-list-item-title>
                  {{$t('tag.no-data')}}
                </v-list-item-title>
              </v-list-item>
            </template>
            <template v-slot:selection="data">
              <v-chip
                v-bind="data.attrs"
                :input-value="data.selected"
                close
                @click:close="removeTagFromSelected(data.item)"
              >
                {{ data.item.name }}
              </v-chip>
            </template>
          </v-autocomplete>
        </v-col>
      </v-row>
    </v-form>
    <slot name="actions" v-bind="{ ref: this.$refs.refss }"/>
  </div>

</template>

<script lang="ts">
import {
  Component, Model, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/route/entity';
import { Tag } from '@/tag/entity';
import _ from 'lodash';

@Component
export default class RouteForm extends Vue {
  valid = true

  selectedTags: string[] = []

  suggestedTags: Tag[] = []

  isSuggestUpdating = false

  searchTags = ''

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '',
    tags: [],
  }

  httpMethodList = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']

  @Watch('selectedTags')
  async onChangeSelectedTags(selected: string[]) {
    this.searchTags = '';
    this.route.tags = selected;
  }

  async delay(ms: number) {
    return new Promise((res) => {
      setTimeout(() => {
        res();
      }, ms);
    });
  }

  @Watch('searchTags')
  async onChangeSearchTags(pattern: string) {
    if (!pattern) return;
    await this.delay(400);
    if (this.isSuggestUpdating) {
      return;
    }
    const namespaceId = Number(this.$route.query.namespaceId);
    this.isSuggestUpdating = true;
    // api suggest
    this.suggestedTags = await this.$store.direct.dispatch.tag.GetByPattern({
      filter: {
        pattern,
        namespaceId,
        page: 1,
        pageSize: 10,
      },
    })
      .finally(() => {
        this.isSuggestUpdating = false;
      });

    // after suggest request focus fades away
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
     this.$refs['autocomplete-input'].focus();
    // if not found => create new tag
    if (this.suggestedTags.length === 0) {
      this.suggestedTags.push({ name: pattern, id: 1 });
    }

    // to keep selected tags in input
    const selected: Tag[] = this.selectedTags.map((tagName) => ({
        name: tagName,
        id: -1,
      }));

    this.suggestedTags.push(...selected);
  }

  @Model('change', { default: {} })
  readonly value!: Route

  isTagInitialized = false

  @Watch('value', { immediate: true, deep: true })
  protected onChangeValue(route: Route): void {
    if (route.id === -1) {
      return;
    }
    if (route.tags && !this.isTagInitialized) {
      this.isTagInitialized = true;
      this.selectedTags = route.tags;
      // without it does not load props...
      this.onChangeSearchTags('test');
    }
    // this.$set(this.route, route);
    this.route = {
      tags: this.selectedTags,
      id: route.id,
      route: route.route,
      method: route.method,
      description: route.description,
    };
  }

  @Watch('route', { deep: true })
  protected onChangeRoute(route: Route): void {
    if (!_.isEqual(route, this.value)) {
      this.$emit('change', route);
    }
  }

  removeTagFromSelected(tag: Tag) {
    this.selectedTags = this.selectedTags.filter((tagName) => tagName !== tag.name);
  }

  routeRules = [
    (v: string) => !!v || 'Обязательное поле',
    (v: string) => {
      if (v) {
        if (v.length) {
          if (v[0] !== '/') {
            return this.$t('route-rule-1');
          }
        }
      }
      return true;
    },
    (v: string) => {
      if (v) {
        if (v.length) {
          if (v.includes('//')) {
            return this.$t('route-rule-2');
          }
        }
      }
      return true;
    },
]

  methodRules = [
    (v: string) => !!v || this.$t('required'),
  ]

  descriptionRules = this.methodRules
}
</script>

<i18n>
{
  "en": {
    "label": {
      "route": "Route",
      "method": "Method",
      "desc": "Description",
      "tags": "tags"
    },
    "tag": {
      "tip": "Enter tag",
      "no-data": "No data"
    },
    "required": "required",
    "route-rule-1": "Route must start with /",
    "route-rule-2": "Route must not contain repeating like //"
  },
  "ru": {
    "label": {
      "route": "Маршрут",
      "method": "Метод",
      "desc": "Описание",
      "tags": "Теги"
    },
    "tag": {
      "tip": "Введите название тега",
      "no-data": "Нет данных"
    },
    "required": "Обязательное поле",
    "route-rule-1":"Маршрут должен начинаться со знака `/`",
    "route-rule-2": "Маршрут не должен содержать повторяющиеся знаки типа //"
  }
}
</i18n>
