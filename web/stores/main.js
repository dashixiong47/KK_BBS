import { defineStore } from 'pinia'
export const useLoginStore = defineStore('login', {
    state: () => ({ loginStatus: false }),
    getters: {
        getLoginStatus: (state) => state.loginStatus,
    },
    actions: {
        setLoginStatus() {
            this.loginStatus = !this.loginStatus
        },
    },
})