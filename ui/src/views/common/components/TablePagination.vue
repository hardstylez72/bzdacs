<template>
  <v-row align="center">
    <v-col cols="auto">
      <div class="m-text-body-l grey--text">{{$t('total')}}{{ total }}</div>
    </v-col>
    <v-spacer />
    <v-col>
      <v-pagination
        color="none"
        total-visible="9"
        v-bind:length="Math.trunc((total - 1) / limit + 1)"
        v-if="total > limit"
        v-model="page" />
    </v-col>
    <v-spacer />
    <v-col cols="auto">
      <v-menu offset-y>
        <template v-slot:activator="{ on }">
          <div class="m-text-body-l grey--text" v-on="on">
            {{$t('pageSize')}}
            <span class="primary--text">{{ limit }}<v-icon color="primary" small>mdi-chevron-down</v-icon></span>
          </div>
        </template>
        <v-list dense>
          <v-list-item v-bind:key="pageSize" v-for="pageSize in pageSizes" v-on:click="$emit('update:limit', pageSize)">
            <v-list-item-title>{{ pageSize }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-col>
  </v-row>
</template>

<script lang="ts">
  import { Component, Prop, Vue } from 'vue-property-decorator';

  @Component
  export default class ItemsTablePagination extends Vue {
    protected readonly pageSizes = [10, 25, 50, 100];

    protected get page() {
      return this.value;
    }

    protected set page(value) {
      this.$emit('input', value);
    }

    @Prop({ default: 10, type: Number })
    protected readonly limit!: number;

    @Prop({ default: 0, type: Number })
    protected readonly total!: number;

    @Prop({ default: 1, type: Number })
    protected readonly value!: number;
  }
</script>

<i18n>
{
  "en": {
    "pageSize": "Items per page: ",
    "total": "Total: "
  },
  "ru": {
    "pageSize": "Показывать по: ",
    "total": "Всего строк: "
  }
}
</i18n>
