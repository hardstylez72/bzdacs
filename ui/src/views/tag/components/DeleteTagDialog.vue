<template>
  <Dialog v-model="show" max-width="450px">
    <v-card>
      <v-card-title>
        Вы уверены что хотите удлить тег?
      </v-card-title>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="blue darken-1"
          text
          @click="close"
        >
          Cancel
        </v-btn>
        <v-btn
          color="blue darken-1"
          text
          @click="deleteTag"
        >
          OK
        </v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import Dialog from '@/views/base/components/Dialog.vue';

@Component({
  components: {
    Dialog,
  },
})
export default class DeleteTagDialog extends Vue {
  show = false

  @Prop() tagId!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }

  async deleteTag() {
    await this.$store.direct.dispatch.tag.Delete(this.tagId);
    this.$emit('change', false);
  }
}
</script>

<style scoped lang="scss">

</style>
