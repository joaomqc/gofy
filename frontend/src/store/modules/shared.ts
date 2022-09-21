import { Module } from "vuex";
import State from "../state";

export interface SharedState {
  isLoading: boolean;
  error: string | null;
}

const sharedModule: Module<SharedState, State> = {
  state: {
    isLoading: false,
    error: null,
  },
  getters: {
    isLoading: (state) => state.isLoading,
    error: (state) => state.error,
  },
  actions: {
    async setLoading({ commit }, isLoading) {
      commit("setLoading", isLoading);
    },
    async setError({ commit }, error) {
      commit("setError", error);
    },
    async clearError({ commit }) {
      commit("clearError");
    },
  },
  mutations: {
    setLoading(state, isLoading) {
      state.isLoading = isLoading;
    },
    setError(state, error) {
      state.error = error;
      setTimeout(() => (state.error = null), 3000);
    },
    clearError(state) {
      state.error = null;
    },
  },
};

export default sharedModule;
