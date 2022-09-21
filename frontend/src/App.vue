<template>
  <v-app>
    <side-menu
      :open="drawerOpen"
      :mobile="mobile"
      @close="drawerOpen = false"
    />
    <v-app-bar app v-if="mobile" dense>
      <v-app-bar-nav-icon
        @click="drawerOpen = !drawerOpen"
      ></v-app-bar-nav-icon>
    </v-app-bar>
    <v-main>
      <v-container fluid v-if="error">
        <v-alert type="error">{{ error }}</v-alert>
      </v-container>
      <v-container fluid>
        <router-view class="main-content" />
      </v-container>
    </v-main>
    <v-overlay :value="isLoading">
      <v-progress-circular indeterminate></v-progress-circular>
    </v-overlay>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import SideMenu from "./components/SideMenu.vue";

export default Vue.extend({
  components: {
    SideMenu,
  },

  computed: {
    mobile() {
      return this.$vuetify.breakpoint.smAndDown;
    },
    isLoading() {
      return this.$store.getters.isLoading;
    },
    error() {
      return this.$store.getters.error;
    },
  },
  data: () => ({
    drawerOpen: false,
  }),
});
</script>
