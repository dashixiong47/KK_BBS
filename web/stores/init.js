import { useGetGroupDetail } from '~/api/server';
import { defineStore } from 'pinia';

export const useGroupStore = defineStore('group', {
    state: () => ({
        group: [],
        error: null,  // 用于存储错误信息
    }),
    getters: {
        // 同步 getter
        getGroup(state) {
            return state.group;
        },
        getError(state) {
            return state.error;
        }
    },
    actions: {
        // 异步 action
        async fetchGroup() {
            if (!Object.keys(this.group).length) {
                try {
                    const {data} = await useGetGroupDetail();
                    this.setGroup(data);
                    this.error = null;  // 清除任何旧的错误信息
                } catch (error) {
                    this.error = error;  // 存储错误信息
                    console.error("An error occurred while fetching the group:", error);
                }
            }
        },
        setGroup(group) {
            this.group = { ...this.group, ...group };
        }
    },
});
