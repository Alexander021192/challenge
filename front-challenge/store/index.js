export const state = () => ({
    sessionId: ""
})

export const mutations = {
    setSession(state,  sessionId ) {
      state.sessionId = sessionId
    }
}
