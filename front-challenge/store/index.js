export const state = () => ({
    sessionId: "sessionTest"
})

export const mutations = {
    setSession(state,  sessionId ) {
      state.sessionId = sessionId
    }
}
