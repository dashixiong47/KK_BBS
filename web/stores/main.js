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
        isLogin:false
    }),
    getters: {
        getUserInfo: (state) => state.userInfo,
        getIsLogin: (state) => state.isLogin,
    },
    actions: {
        async fetchUserInfo(id) {
            if (!Object.keys(this.userInfo).length) {
             
                try {
                    const data = await getUserInfo(id);
                    this.setUserInfo(data);
                    this.setIslogin(true)
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
        },
        setIslogin(bl){
            this.isLogin=bl
        },

    },
})