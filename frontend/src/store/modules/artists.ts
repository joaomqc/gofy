import { Module } from "vuex";
import State from "../state";
import { requestArtists } from "@/server/requests";

export interface ArtistsState {
  artists: any[];
}

const sharedModule: Module<ArtistsState, State> = {
  state: {
    artists: [],
  },
  getters: {
    artists: (state) => state.artists,
  },
  actions: {
    async getArtists(state) {
      state.dispatch("setLoading", true);
      await state.dispatch("clearError");

      const { error, data } = await requestArtists();

      if (error) {
        state.dispatch("setLoading", false);
        await state.dispatch("setError", error);
      } else {
        state.commit("setResult", { data });
        state.dispatch("setLoading", false);
      }
    },
  },
  mutations: {
    setResult(state, { data }) {
      state.artists = data;
    },
  },
};

export default sharedModule;
