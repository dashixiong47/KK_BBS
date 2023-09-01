import { defineStore } from 'pinia'
import { getUserInfo } from '~/api';
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
        async fetchUserInfo(id) {
            if (!Object.keys(this.userInfo).length) {
                try {
                    const data = await getUserInfo(id);
                    console.log(data);
                    this.setUserInfo(data);
                } catch (error) {
                    console.error("An error occurred while fetching the userInfo:", error);
                }
            }
        },
        setUserInfo(userInfo) {
            this.userInfo = { ...this.userInfo, ...userInfo }
        },
        resetUserInfo() {
            this.userInfo = {}
        }
    },
})