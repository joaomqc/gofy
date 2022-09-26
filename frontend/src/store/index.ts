import Vue from "vue";
import Vuex from "vuex";
// import createPersistedState from "vuex-persistedstate";
import shared from "./modules/shared";
import search from "./modules/search";
import artists from "./modules/artists";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    shared,
    search,
    artists,
  },
  //   plugins: [createPersistedState({ paths: [""] })],
});
