import { defineStore } from 'pinia'
import { useGetUserInfo } from '~/api/server';
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
        isLogin: false
    }),
    getters: {
        getUserInfo: (state) => state.userInfo,
        getIsLogin: (state) => state.isLogin,
    },
    actions: {
        async fetchUserInfo() {
            try {
                const { data } = await useGetUserInfo();
                this.setIslogin(true)
                this.setUserInfo(data);
            } catch (error) {
                this.setIslogin(false)
                console.error("An error occurred while fetching the userInfo:", error);
            }
        },
        setUserInfo(userInfo) {
            this.userInfo = userInfo
        },
        resetUserInfo() {
            this.userInfo = {}
        },
        setIslogin(bl) {
            this.isLogin = bl
        },

    },
})