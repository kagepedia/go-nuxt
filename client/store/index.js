export const state = () => ({
  loggedIn: false
})

export const mutations = {
  setLogin(state, payload) {
    state.loggedIn = payload
  }
}

export const actions = {
  userLogin({ commit }, payload) {
    commit('setLogin', payload)
  }
}
