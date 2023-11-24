import { useGetGroupDetail, useGetHost } from '~/api/server';
import { defineStore } from 'pinia';

export const useGroupStore = defineStore('group', {
    state: () => ({
        group: [],
        actived: 0,
        error: null,  // 用于存储错误信息
    }),
    getters: {
        // 同步 getter
        getGroup(state) {
            return state.group;
        },
        getActived(state) {
            return state.actived
        },
        getError(state) {
            return state.error;
        },

    },
    actions: {
        // 异步 action
        async fetchGroup() {
            if (!Object.keys(this.group).length) {
                try {
                    const { data } = await useGetGroupDetail();
                    this.setGroup(data.sort((a, b) => a.order - b.order));

                    this.error = null;  // 清除任何旧的错误信息
                } catch (error) {
                    this.error = error;  // 存储错误信息
                    console.error("An error occurred while fetching the group:", error);
                }
            }
        },

        setGroup(group) {
            this.group = group
        },
        setActived(id) {
            this.actived = id
        }

    },
});

export const useAppConfigStore = defineStore("appConfig", {
    state: () => ({
        config: {
            appName: "kkbbs"
        },
        host: ''
    }),
    getters: {
        getConfig(state) {
            return state.config;
        },
        getHost(state) {
            return state.host;
        }
    },
    actions: {
        async fetchHost() {
            try {
                const { data } = await useGetHost();
                this.setHost(data);

            } catch (error) {
                console.error("An error occurred while fetching the host:", error);
            }
        },
        setHost(host) {
            this.host = host;
        }
    }

})
