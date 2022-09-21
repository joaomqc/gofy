<template>
  <v-navigation-drawer
    app
    :permanent="!mobile"
    :temporary="mobile"
    v-model="show"
  >
    <v-list-item>
      <v-list-item-content>
        <v-list-item-title class="text-h6"> Gofy </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
    <v-divider></v-divider>
    <div v-for="item in items" :key="item.title">
      <v-list-item
        :to="item.to"
        v-model="item.selected"
        :class="{ focused: !mobile && item.selected }"
      >
        <v-list-item-icon>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </div>
  </v-navigation-drawer>
</template>

<script lang="ts">
import Vue from "vue";

interface Link {
  title: string;
  icon: string;
  to: string;
  selected?: boolean;
}

export default Vue.extend({
  props: {
    mobile: Boolean,
    open: Boolean,
  },

  data: () => ({
    items: [
      {
        title: "Monitor",
        icon: "mdi-radar",
        to: "/monitor",
      },
      {
        title: "Search",
        icon: "mdi-magnify",
        to: "/search",
      },
      {
        title: "Settings",
        icon: "mdi-cog",
        to: "/settings",
      },
    ] as Link[],
  }),

  computed: {
    show: {
      get(): boolean {
        return !this.mobile || this.open;
      },
      set(val: boolean) {
        if (!val) {
          this.$emit("close");
        }
      },
    },
  },
});
</script>
