
import { messageList, unreadMessage } from '@/api/index'
import { defineStore } from 'pinia'
export const useMessageStore = defineStore('message', {
    state: () => ({
        unreadMessage: 0,
        messageData: [],
        messageTotal: 0,
    }),
    getters: {
        getUnreadMessage: (state) => {
            return state.unreadMessage
        },
        getMessageData: (state) => {
            return state.messageData
        },
        getMessageTotal: (state) => {
            return state.messageTotal
        },

    },
    actions: {
        async fetchMessageList(pages={}) {
            try {
                const { data } = await messageList(pages);
                this.setMessageData(data.list);
                this.setMessageTotal(data.total);
            } catch (error) {
                console.error("An error occurred while fetching the messageList:", error);
            }
        },
        async fetchUnreadMessage() {
            try {
                const { data } = await unreadMessage();
                this.setUnreadMessage(data);
            } catch (error) {
                console.error("An error occurred while fetching the unreadMessage:", error);
            }
        },
        setUnreadMessage(num) {
            this.unreadMessage = num
        },
        setMessageData(data) {
            this.messageData = data
        },
        setMessageTotal(num) {
            this.messageTotal = num
        },
    },
})