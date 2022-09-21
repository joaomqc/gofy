import Vue from "vue";
import Vuex from "vuex";
// import createPersistedState from "vuex-persistedstate";
import shared from "./modules/shared";
import search from "./modules/search";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    shared,
    search,
  },
  //   plugins: [createPersistedState({ paths: [""] })],
});
