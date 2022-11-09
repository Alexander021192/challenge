export const state = () => ({
    sessionId: "sessionId"
})

export const mutations = {
    setSession(state,  sessionId ) {
      state.sessionId = sessionId
    }
}