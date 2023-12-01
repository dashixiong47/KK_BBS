export const useSocketStore = defineStore('socket', {
    state: () => ({
        message: null
    }),
    getters: {
        getSocketMsg(state) {
            return state.message
        }
    },
    actions: {
        setSocketMsg(msg) {
            this.message = msg
        }
    }
})