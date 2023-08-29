import { defineStore } from 'pinia'
export const useLoginStore = defineStore('login', {
    state: () => ({
        loginStatus: false,

    }),
    getters: {
        getLoginStatus: (state) => state.loginStatus,

    },
    actions: {
        setLoginStatus() {
            this.loginStatus = !this.loginStatus
        },
    },
})
export const useUserStore = defineStore('userInfo', {
    state: () => ({
        userInfo: {},
    }),
    getters: {
        getUserInfo: (state) => state.userInfo,
    },
    actions: {
        setUserInfo(userInfo) {
            this.userInfo = { ...this.userInfo, ...userInfo }
        }
    },
})