<template>
  <v-container fluid>
    <v-text-field
      label="search"
      v-model="searchTerm"
      @keydown.enter="onEnterKey"
    />
    <v-tabs v-model="tab" @click="onSwitchTab">
      <v-tab @click="onSwitchTab(0)"> Artists </v-tab>
      <v-tab @click="onSwitchTab(1)"> Albums </v-tab>
      <v-tab @click="onSwitchTab(2)"> Tracks </v-tab>
      <v-tab @click="onSwitchTab(3)"> Playlists </v-tab>
    </v-tabs>
    <v-tabs-items class="mt-4" v-model="tab">
      <v-tab-item>
        <search-result-list type="artist" />
      </v-tab-item>
      <v-tab-item>
        <search-result-list type="album" />
      </v-tab-item>
      <v-tab-item>
        <search-result-list type="track" />
      </v-tab-item>
      <v-tab-item>
        <search-result-list type="playlist" />
      </v-tab-item>
    </v-tabs-items>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import SearchResultList from "./SearchResultList.vue";

export default Vue.extend({
  components: {
    SearchResultList,
  },
  data: () => ({
    tab: 0,
    searchTerm: null,
  }),
  computed: {
    isLoading() {
      return this.$store.getters.isLoading;
    },
    error() {
      return this.$store.getters.error;
    },
  },
  methods: {
    async search(tabId: number): Promise<void> {
      if (this.searchTerm) {
        await this.$store.dispatch("search", {
          searchTerm: this.searchTerm,
          type: this.getTabType(tabId),
        });
      }
    },
    async onSwitchTab(tabId: number): Promise<void> {
      await this.search(tabId);
    },
    async onEnterKey(): Promise<void> {
      await this.search(this.tab);
    },
    getTabType(tabId: number): string {
      switch (tabId) {
        case 0:
          return "artist";
        case 1:
          return "album";
        case 2:
          return "track";
        case 3:
          return "playlist";
      }
      return "";
    },
  },
});
</script>
