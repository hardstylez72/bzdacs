<template>
  <v-row class="justify-center ">
    <v-col cols="12" sm="10" md="10">
      <v-autocomplete
        label="Теги"
        placeholder="Введите название тега"
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
              Тег не найден
            </v-list-item-title>
          </v-list-item>
        </template>
        <template v-slot:selection="data">
          <v-chip
            v-bind="data.attrs"
            :input-value="data.selected"
            :close="isSelectedTagDeletable(data.item)"
            @click:close="removeTagFromSelected(data.item)"
          >
            {{ data.item }}
          </v-chip>
        </template>
      </v-autocomplete>
      <v-switch
        v-if="this.selectedTags.length"
        :label="filter.tags.exclude ? 'Показываются все теги кроме выбранных' : 'Показываются только выбранные теги'"
        v-model="filter.tags.exclude"
      ></v-switch>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import {
  Component, Model, Vue, Watch,
} from 'vue-property-decorator';
import { Tag } from '@/views/tag/service';
import { Filter, Route } from '@/views/route/service';

@Component

export default class TagFilter extends Vue {
  filter: Filter = {
    tags: {
      names: ['sys'],
      exclude: true,
    },
  }

  valid = true

  excludeAllSelected = true

  selectedTags: string[] = []

  suggestedTags: string[] = []

  isSuggestUpdating = false

  searchTags = ''

  created() {
    this.selectedTags = this.filter.tags.names;
    this.$store.direct.commit.route.setFilter(this.filter);
  }

  @Watch('filter', { deep: true })
  async onChangeFilter(filter: Filter) {
    this.$store.direct.commit.route.setFilter(filter);
  }

  @Watch('storeFilter', { deep: true })
  async onChangeStoreFilter(f: Filter) {
    this.selectedTags = f.tags.names;
    this.suggestedTags = f.tags.names;
    this.filter = f;
  }

  get storeFilter() {
    return this.$store.direct.getters.route.getFilter;
  }

  @Watch('selectedTags')
  async onChangeSelectedTags(selected: string[]) {
    this.searchTags = '';
    this.filter.tags.names = selected;
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
    // if (!pattern) return;
    await this.delay(400);
    if (this.isSuggestUpdating) {
      return;
    }
    this.isSuggestUpdating = true;
    // api suggest
    const tags = await this.$store.direct.dispatch.tag.GetByPattern(pattern)
      .finally(() => {
        this.isSuggestUpdating = false;
      });

    this.suggestedTags = tags.map((tag) => tag.name);

    // // after suggest request focus fades away
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs['autocomplete-input'].focus();

    // to keep selected tags in input
    this.suggestedTags.push(...this.selectedTags);
  }

  isSelectedTagDeletable(tag: string): boolean {
    const lastSelectedTag = this.selectedTags[this.selectedTags.length - 1];
    return tag === lastSelectedTag;
  }

  removeTagFromSelected(tag: string) {
    this.selectedTags = this.selectedTags.filter((tagName) => tagName !== tag);
  }
}
</script>

<style scoped lang="scss">

</style>
